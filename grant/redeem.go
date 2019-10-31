package grant

import (
	"context"
	"fmt"
	"time"

	"github.com/brave-intl/bat-go/utils/altcurrency"
	"github.com/brave-intl/bat-go/utils/closers"
	"github.com/brave-intl/bat-go/wallet"
	"github.com/brave-intl/bat-go/wallet/provider"
	raven "github.com/getsentry/raven-go"
	"github.com/pkg/errors"
	"github.com/pressly/lg"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// RedeemGrantsRequest a request to redeem the included grants for the wallet whose information
// is included in order to fulfill the included transaction
type RedeemGrantsRequest struct {
	WalletInfo  wallet.Info `json:"wallet" valid:"required"`
	Transaction string      `json:"transaction" valid:"base64"`
}

// RedemptionDisabled due to fail safe condition
func RedemptionDisabled() bool {
	return safeMode || breakerTripped
}

// Consume one or more grants to fulfill the included transaction for wallet
// Note that this is destructive, on success consumes grants.
// Further calls to Verify with the same request will fail as the grants are consumed.
//
// 1. Sort grants, closest expiration to furthest, short circuit if no grants
//
// 2. Enforce transaction checks and verify transaction signature
//
// 3. Sum from largest to smallest until value is gt transaction amount
//
// 4. Iterate through grants and check that:
//
// a) this wallet has not yet redeemed a grant for the given promotionId
//
// b) this grant has not yet been redeemed by any wallet
//
// Returns transaction info for grant fufillment
func (service *Service) Consume(ctx context.Context, walletInfo wallet.Info, transaction string) (*wallet.TransactionInfo, error) {
	var txProbi *decimal.Decimal
	var redeemTxInfo wallet.TransactionInfo
	{
		tmp := altcurrency.BAT
		redeemTxInfo.AltCurrency = &tmp
	}

	// 1. Sort grants, closest expiration to furthest, short circuit if no grants
	unredeemedGrants, err := service.datastore.GetGrantsOrderedByExpiry(walletInfo)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch grants ordered by expiration date")
	}

	if len(unredeemedGrants) == 0 {
		return nil, nil
	}

	// 2. Enforce transaction checks and verify transaction signature
	userWallet, err := provider.GetWallet(walletInfo)
	if err != nil {
		return nil, err
	}
	// this ensures we have a valid wallet if refreshBalance == true
	balance, err := userWallet.GetBalance(refreshBalance)
	if err != nil {
		return nil, err
	}

	if len(transaction) > 0 {
		// 1. Enforce transaction checks and verify transaction signature
		// NOTE for uphold provider we currently check against user provided publicKey
		//      thus this check does not protect us from a valid fake signature
		txInfo, err := userWallet.VerifyTransaction(transaction)
		if err != nil {
			return nil, err
		}
		if *txInfo.AltCurrency != altcurrency.BAT {
			return nil, errors.New("only grants submitted with BAT transactions are supported")
		}
		if txInfo.Probi.LessThan(decimal.Zero) {
			return nil, errors.New("included transaction cannot be for negative BAT")
		}
		if txInfo.Probi.LessThan(altcurrency.BAT.ToProbi(decimal.New(lowerTxLimit, 0))) {
			return nil, fmt.Errorf("included transaction must be for a minimum of %d BAT", lowerTxLimit)
		}
		if txInfo.Probi.GreaterThan(altcurrency.BAT.ToProbi(decimal.New(upperTxLimit, 0))) {
			return nil, fmt.Errorf("included transaction must be for a maxiumum of %d BAT", upperTxLimit)
		}
		if txInfo.Destination != SettlementDestination {
			return nil, errors.New("included transactions must have settlement as their destination")
		}

		txProbi = &txInfo.Probi
	}

	// 3. Sum until value is gt transaction amount
	var grants []Grant
	sumProbi := decimal.New(0, 1)
	for _, grant := range unredeemedGrants {
		if txProbi != nil {
			if sumProbi.GreaterThanOrEqual(*txProbi) {
				break
			}
		}
		if *grant.AltCurrency != altcurrency.BAT {
			return nil, errors.New("All grants must be in BAT")
		}
		sumProbi = sumProbi.Add(grant.Probi)
		grants = append(grants, grant)
	}

	if txProbi != nil && txProbi.GreaterThan(balance.SpendableProbi.Add(sumProbi)) {
		return nil, errors.New("wallet does not have enough funds to cover transaction")
	}

	// should be reasonable since we limit the redeem endpoint to a maximum of 1 simultaneous in-flight request
	ugpBalance, err := grantWallet.GetBalance(refreshBalance)
	if err != nil {
		return nil, err
	}

	if sumProbi.GreaterThan(ugpBalance.SpendableProbi) {
		safeMode = true
		raven.CaptureMessage("Hot wallet out of funds!!!", map[string]string{"out-of-funds": "true"})
		return nil, errors.New("ugp wallet lacks enough funds to fulfill grants")
	}

	if len(transaction) > 0 && testSubmit {
		var submitInfo *wallet.TransactionInfo
		// TODO remove this once we can retrieve publicKey info from uphold
		// NOTE We check the signature on the included transaction by submitting it but not confirming it
		submitInfo, err = userWallet.SubmitTransaction(transaction, false)
		if err != nil {
			if wallet.IsInvalidSignature(err) {
				return nil, errors.New("the included transaction was signed with the wrong publicKey")
			} else if !wallet.IsInsufficientBalance(err) {
				return nil, errors.New("error while test submitting the included transaction: " + err.Error())
			}
		}
		redeemTxInfo.ID = submitInfo.ID
	}

	// 4. Iterate through grants and check that:
	for _, grant := range grants {
		// the grant is mature
		if time.Now().Unix() < grant.MaturityTimestamp {
			return nil, errors.New("Grant is not yet redeemable as it is immature")
		}

		// the grant is not expired
		if time.Now().Unix() > grant.ExpiryTimestamp {
			return nil, errors.New("Grant is expired")
		}

		err = service.datastore.RedeemGrantForWallet(grant, walletInfo)
		if err != nil {
			return nil, err
		}

		redeemedGrantsCounter.With(prometheus.Labels{"promotionId": grant.PromotionID.String()}).Inc()
	}

	redeemTxInfo.Probi = sumProbi
	redeemTxInfo.Destination = walletInfo.ProviderID
	return &redeemTxInfo, nil
}

// GetRedeemedIDs returns a list of any grants that have already been redeemed
func (service *Service) GetRedeemedIDs(ctx context.Context, Grants []string) ([]string, error) {

	// 1. Check grant signatures and decode
	grants, err := DecodeGrants(grantPublicKey, Grants)
	if err != nil {
		return nil, err
	}
	grantCount := len(grants)
	results := make([]string, 0, grantCount)

	for _, grant := range grants {
		grantRedeemed, err := service.datastore.HasGrantBeenRedeemed(grant)
		if err != nil {
			return nil, err
		}
		if grantRedeemed {
			results = append(results, grant.GrantID.String())
		}
	}

	return results, nil
}

// RedeemGrantsResponse includes information about the transaction to settlement and the grant funds used
type RedeemGrantsResponse struct {
	wallet.TransactionInfo
	GrantTotal decimal.Decimal `json:"grantTotal"`
}

// Redeem the grants in the included response
func (service *Service) Redeem(ctx context.Context, req *RedeemGrantsRequest) (*RedeemGrantsResponse, error) {
	log := lg.Log(ctx)

	if RedemptionDisabled() {
		return nil, errors.New("Grant redemption has been disabled due to fail-safe condition")
	}

	grantFulfillmentInfo, err := service.Consume(ctx, req.WalletInfo, req.Transaction)
	if err != nil {
		return nil, err
	}

	if grantFulfillmentInfo == nil {
		return nil, nil
	}

	submitID := grantFulfillmentInfo.ID

	userWallet, err := provider.GetWallet(req.WalletInfo)
	if err != nil {
		log.Errorf("Could not get wallet %s from info after successful Consume", req.WalletInfo.ProviderID)
		raven.CaptureMessage("Could not get wallet after successful Consume", map[string]string{"providerID": req.WalletInfo.ProviderID})
		return nil, err
	}

	// fund user wallet with probi from grants
	_, err = grantWallet.Transfer(*grantFulfillmentInfo.AltCurrency, grantFulfillmentInfo.Probi, grantFulfillmentInfo.Destination)
	if err != nil {

		log.Errorf("Could not fund wallet %s after successful Consume", req.WalletInfo.ProviderID)
		raven.CaptureMessage("Could not fund wallet after successful Consume", map[string]string{"providerID": req.WalletInfo.ProviderID})
		return nil, err
	}

	// confirm settlement transaction previously sent to wallet provider
	var settlementInfo *wallet.TransactionInfo
	for tries := 5; tries >= 0; tries-- {
		// NOTE Consume (by way of VerifyTransaction) guards against transactions that seek to exploit parser differences
		// such as including additional fields that are not understood by this wallet provider implementation but may
		// be understood by the upstream wallet provider.
		settlementInfo, err = userWallet.ConfirmTransaction(submitID)
		if err == nil {
			break
		}
	}

	return &RedeemGrantsResponse{TransactionInfo: *settlementInfo, GrantTotal: grantFulfillmentInfo.Probi}, nil
}

// DrainGrantsRequest a request to drain a wallets grains to a linked uphold account
type DrainGrantsRequest struct {
	WalletInfo       wallet.Info `json:"wallet" valid:"required"`
	AnonymousAddress uuid.UUID   `json:"anonymousAddress" valid:"required"`
}

// Drain the grants for the wallet in the included response
func (service *Service) Drain(ctx context.Context, req *DrainGrantsRequest) error {
	log := lg.Log(ctx)

	if RedemptionDisabled() {
		return errors.New("Grant redemption has been disabled due to fail-safe condition")
	}

	grantFulfillmentInfo, err := service.Consume(ctx, req.WalletInfo, "")
	if err != nil {
		return err
	}

	// drain probi from grants into user wallet
	_, err = grantWallet.Transfer(*grantFulfillmentInfo.AltCurrency, grantFulfillmentInfo.Probi, req.AnonymousAddress.String())
	if err != nil {
		conn := service.redisPool.Get()
		defer closers.Panic(conn)
		b := GetBreaker(&conn)

		incErr := b.Increment()
		if incErr != nil {
			log.Errorf("Could not increment the breaker!!!")
			raven.CaptureMessage("Could not increment the breaker!!!", map[string]string{"breaker": "true"})
			safeMode = true
		}

		log.Errorf("Could not drain into wallet %s after successful Consume", req.WalletInfo.ProviderID)
		raven.CaptureMessage("Could not drain into wallet after successful Consume", map[string]string{"providerID": req.WalletInfo.ProviderID})
		return err
	}
	return nil
}
