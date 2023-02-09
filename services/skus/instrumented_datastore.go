package skus

// Code generated by gowrap. DO NOT EDIT.
// template: ../../.prom-gowrap.tmpl
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p github.com/brave-intl/bat-go/services/skus -i Datastore -t ../../.prom-gowrap.tmpl -o instrumented_datastore.go -l ""

import (
	"context"
	"time"

	"github.com/brave-intl/bat-go/libs/inputs"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

// DatastoreWithPrometheus implements Datastore interface with all methods wrapped
// with Prometheus metrics
type DatastoreWithPrometheus struct {
	base         Datastore
	instanceName string
}

var datastoreDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "skus_datastore_duration_seconds",
		Help:       "datastore runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewDatastoreWithPrometheus returns an instance of the Datastore decorated with prometheus summary metric
func NewDatastoreWithPrometheus(base Datastore, instanceName string) DatastoreWithPrometheus {
	return DatastoreWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// AppendOrderMetadata implements Datastore
func (_d DatastoreWithPrometheus) AppendOrderMetadata(ctx context.Context, up1 *uuid.UUID, s1 string, s2 string) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "AppendOrderMetadata", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.AppendOrderMetadata(ctx, up1, s1, s2)
}

// AppendOrderMetadataInt implements Datastore
func (_d DatastoreWithPrometheus) AppendOrderMetadataInt(ctx context.Context, up1 *uuid.UUID, s1 string, i1 int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "AppendOrderMetadataInt", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.AppendOrderMetadataInt(ctx, up1, s1, i1)
}

// AreTimeLimitedV2CredsSubmitted implements Datastore
func (_d DatastoreWithPrometheus) AreTimeLimitedV2CredsSubmitted(ctx context.Context, blindedCreds ...string) (b1 bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "AreTimeLimitedV2CredsSubmitted", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.AreTimeLimitedV2CredsSubmitted(ctx, blindedCreds...)
}

// BeginTx implements Datastore
func (_d DatastoreWithPrometheus) BeginTx() (tp1 *sqlx.Tx, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "BeginTx", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.BeginTx()
}

// CheckExpiredCheckoutSession implements Datastore
func (_d DatastoreWithPrometheus) CheckExpiredCheckoutSession(u1 uuid.UUID) (b1 bool, s1 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "CheckExpiredCheckoutSession", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CheckExpiredCheckoutSession(u1)
}

// CommitVote implements Datastore
func (_d DatastoreWithPrometheus) CommitVote(ctx context.Context, vr VoteRecord, tx *sqlx.Tx) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "CommitVote", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CommitVote(ctx, vr, tx)
}

// CreateKey implements Datastore
func (_d DatastoreWithPrometheus) CreateKey(merchant string, name string, encryptedSecretKey string, nonce string) (kp1 *Key, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateKey(merchant, name, encryptedSecretKey, nonce)
}

// CreateOrder implements Datastore
func (_d DatastoreWithPrometheus) CreateOrder(totalPrice decimal.Decimal, merchantID string, status string, currency string, location string, validFor *time.Duration, orderItems []OrderItem, allowedPaymentMethods *Methods) (op1 *Order, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateOrder(totalPrice, merchantID, status, currency, location, validFor, orderItems, allowedPaymentMethods)
}

// CreateTransaction implements Datastore
func (_d DatastoreWithPrometheus) CreateTransaction(orderID uuid.UUID, externalTransactionID string, status string, currency string, kind string, amount decimal.Decimal) (tp1 *Transaction, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateTransaction", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateTransaction(orderID, externalTransactionID, status, currency, kind, amount)
}

// DeleteKey implements Datastore
func (_d DatastoreWithPrometheus) DeleteKey(id uuid.UUID, delaySeconds int) (kp1 *Key, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteKey(id, delaySeconds)
}

// DeleteOrderCreds implements Datastore
func (_d DatastoreWithPrometheus) DeleteOrderCreds(orderID uuid.UUID, isSigned bool) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteOrderCreds", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteOrderCreds(orderID, isSigned)
}

// DeleteTimeLimitedV2OrderCredsByOrder implements Datastore
func (_d DatastoreWithPrometheus) DeleteTimeLimitedV2OrderCredsByOrder(orderID uuid.UUID) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteTimeLimitedV2OrderCredsByOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteTimeLimitedV2OrderCredsByOrder(orderID)
}

// GetIssuer implements Datastore
func (_d DatastoreWithPrometheus) GetIssuer(merchantID string) (ip1 *Issuer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetIssuer", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetIssuer(merchantID)
}

// GetIssuerByPublicKey implements Datastore
func (_d DatastoreWithPrometheus) GetIssuerByPublicKey(publicKey string) (ip1 *Issuer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetIssuerByPublicKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetIssuerByPublicKey(publicKey)
}

// GetKey implements Datastore
func (_d DatastoreWithPrometheus) GetKey(id uuid.UUID, showExpired bool) (kp1 *Key, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetKey", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetKey(id, showExpired)
}

// GetKeysByMerchant implements Datastore
func (_d DatastoreWithPrometheus) GetKeysByMerchant(merchant string, showExpired bool) (kap1 *[]Key, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetKeysByMerchant", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetKeysByMerchant(merchant, showExpired)
}

// GetOrder implements Datastore
func (_d DatastoreWithPrometheus) GetOrder(orderID uuid.UUID) (op1 *Order, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOrder(orderID)
}

// GetOrderByExternalID implements Datastore
func (_d DatastoreWithPrometheus) GetOrderByExternalID(externalID string) (op1 *Order, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOrderByExternalID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOrderByExternalID(externalID)
}

// GetOrderCreds implements Datastore
func (_d DatastoreWithPrometheus) GetOrderCreds(orderID uuid.UUID, isSigned bool) (oa1 []OrderCreds, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOrderCreds", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOrderCreds(orderID, isSigned)
}

// GetOrderCredsByItemID implements Datastore
func (_d DatastoreWithPrometheus) GetOrderCredsByItemID(orderID uuid.UUID, itemID uuid.UUID, isSigned bool) (op1 *OrderCreds, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOrderCredsByItemID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOrderCredsByItemID(orderID, itemID, isSigned)
}

// GetOrderItem implements Datastore
func (_d DatastoreWithPrometheus) GetOrderItem(ctx context.Context, itemID uuid.UUID) (op1 *OrderItem, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOrderItem", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOrderItem(ctx, itemID)
}

// GetOutboxMovAvgDurationSeconds implements Datastore
func (_d DatastoreWithPrometheus) GetOutboxMovAvgDurationSeconds() (i1 int64, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetOutboxMovAvgDurationSeconds", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetOutboxMovAvgDurationSeconds()
}

// GetPagedMerchantTransactions implements Datastore
func (_d DatastoreWithPrometheus) GetPagedMerchantTransactions(ctx context.Context, merchantID uuid.UUID, pagination *inputs.Pagination) (tap1 *[]Transaction, i1 int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetPagedMerchantTransactions", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetPagedMerchantTransactions(ctx, merchantID, pagination)
}

// GetSigningOrderRequestOutboxByOrder implements Datastore
func (_d DatastoreWithPrometheus) GetSigningOrderRequestOutboxByOrder(ctx context.Context, orderID uuid.UUID) (sa1 []SigningOrderRequestOutbox, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetSigningOrderRequestOutboxByOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetSigningOrderRequestOutboxByOrder(ctx, orderID)
}

// GetSigningOrderRequestOutboxByOrderItem implements Datastore
func (_d DatastoreWithPrometheus) GetSigningOrderRequestOutboxByOrderItem(ctx context.Context, itemID uuid.UUID) (sa1 []SigningOrderRequestOutbox, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetSigningOrderRequestOutboxByOrderItem", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetSigningOrderRequestOutboxByOrderItem(ctx, itemID)
}

// GetSigningOrderRequestOutboxByRequestID implements Datastore
func (_d DatastoreWithPrometheus) GetSigningOrderRequestOutboxByRequestID(ctx context.Context, requestID uuid.UUID) (sp1 *SigningOrderRequestOutbox, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetSigningOrderRequestOutboxByRequestID", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetSigningOrderRequestOutboxByRequestID(ctx, requestID)
}

// GetSumForTransactions implements Datastore
func (_d DatastoreWithPrometheus) GetSumForTransactions(orderID uuid.UUID) (d1 decimal.Decimal, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetSumForTransactions", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetSumForTransactions(orderID)
}

// GetTimeLimitedV2OrderCredsByOrder implements Datastore
func (_d DatastoreWithPrometheus) GetTimeLimitedV2OrderCredsByOrder(orderID uuid.UUID) (tp1 *TimeLimitedV2Creds, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetTimeLimitedV2OrderCredsByOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetTimeLimitedV2OrderCredsByOrder(orderID)
}

// GetTimeLimitedV2OrderCredsByOrderItem implements Datastore
func (_d DatastoreWithPrometheus) GetTimeLimitedV2OrderCredsByOrderItem(itemID uuid.UUID) (tp1 *TimeLimitedV2Creds, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetTimeLimitedV2OrderCredsByOrderItem", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetTimeLimitedV2OrderCredsByOrderItem(itemID)
}

// GetTransaction implements Datastore
func (_d DatastoreWithPrometheus) GetTransaction(externalTransactionID string) (tp1 *Transaction, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetTransaction", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetTransaction(externalTransactionID)
}

// GetTransactions implements Datastore
func (_d DatastoreWithPrometheus) GetTransactions(orderID uuid.UUID) (tap1 *[]Transaction, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetTransactions", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetTransactions(orderID)
}

// GetUncommittedVotesForUpdate implements Datastore
func (_d DatastoreWithPrometheus) GetUncommittedVotesForUpdate(ctx context.Context) (tp1 *sqlx.Tx, vpa1 []*VoteRecord, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "GetUncommittedVotesForUpdate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetUncommittedVotesForUpdate(ctx)
}

// InsertIssuer implements Datastore
func (_d DatastoreWithPrometheus) InsertIssuer(issuer *Issuer) (ip1 *Issuer, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertIssuer", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertIssuer(issuer)
}

// InsertOrderCredsTx implements Datastore
func (_d DatastoreWithPrometheus) InsertOrderCredsTx(ctx context.Context, tx *sqlx.Tx, creds *OrderCreds) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertOrderCredsTx", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertOrderCredsTx(ctx, tx, creds)
}

// InsertSignedOrderCredentialsTx implements Datastore
func (_d DatastoreWithPrometheus) InsertSignedOrderCredentialsTx(ctx context.Context, tx *sqlx.Tx, signedOrderResult *SigningOrderResult) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertSignedOrderCredentialsTx", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertSignedOrderCredentialsTx(ctx, tx, signedOrderResult)
}

// InsertSigningOrderRequestOutbox implements Datastore
func (_d DatastoreWithPrometheus) InsertSigningOrderRequestOutbox(ctx context.Context, requestID uuid.UUID, orderID uuid.UUID, itemID uuid.UUID, signingOrderRequest SigningOrderRequest) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertSigningOrderRequestOutbox", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertSigningOrderRequestOutbox(ctx, requestID, orderID, itemID, signingOrderRequest)
}

// InsertTimeLimitedV2OrderCredsTx implements Datastore
func (_d DatastoreWithPrometheus) InsertTimeLimitedV2OrderCredsTx(ctx context.Context, tx *sqlx.Tx, tlv2 TimeAwareSubIssuedCreds) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertTimeLimitedV2OrderCredsTx", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertTimeLimitedV2OrderCredsTx(ctx, tx, tlv2)
}

// InsertVote implements Datastore
func (_d DatastoreWithPrometheus) InsertVote(ctx context.Context, vr VoteRecord) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "InsertVote", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.InsertVote(ctx, vr)
}

// IsStripeSub implements Datastore
func (_d DatastoreWithPrometheus) IsStripeSub(u1 uuid.UUID) (b1 bool, s1 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "IsStripeSub", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.IsStripeSub(u1)
}

// MarkVoteErrored implements Datastore
func (_d DatastoreWithPrometheus) MarkVoteErrored(ctx context.Context, vr VoteRecord, tx *sqlx.Tx) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "MarkVoteErrored", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.MarkVoteErrored(ctx, vr, tx)
}

// Migrate implements Datastore
func (_d DatastoreWithPrometheus) Migrate(p1 ...uint) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "Migrate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Migrate(p1...)
}

// NewMigrate implements Datastore
func (_d DatastoreWithPrometheus) NewMigrate() (mp1 *migrate.Migrate, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "NewMigrate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.NewMigrate()
}

// RawDB implements Datastore
func (_d DatastoreWithPrometheus) RawDB() (dp1 *sqlx.DB) {
	_since := time.Now()
	defer func() {
		result := "ok"
		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RawDB", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawDB()
}

// RenewOrder implements Datastore
func (_d DatastoreWithPrometheus) RenewOrder(ctx context.Context, orderID uuid.UUID) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RenewOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RenewOrder(ctx, orderID)
}

// RollbackTx implements Datastore
func (_d DatastoreWithPrometheus) RollbackTx(tx *sqlx.Tx) {
	_since := time.Now()
	defer func() {
		result := "ok"
		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RollbackTx", result).Observe(time.Since(_since).Seconds())
	}()
	_d.base.RollbackTx(tx)
	return
}

// RollbackTxAndHandle implements Datastore
func (_d DatastoreWithPrometheus) RollbackTxAndHandle(tx *sqlx.Tx) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RollbackTxAndHandle", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RollbackTxAndHandle(tx)
}

// RunNextOrderJob implements Datastore
func (_d DatastoreWithPrometheus) RunNextOrderJob(ctx context.Context, worker OrderWorker) (b1 bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "RunNextOrderJob", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RunNextOrderJob(ctx, worker)
}

// SendSigningRequest implements Datastore
func (_d DatastoreWithPrometheus) SendSigningRequest(ctx context.Context, signingRequestWriter SigningRequestWriter) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "SendSigningRequest", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SendSigningRequest(ctx, signingRequestWriter)
}

// SetOrderPaid implements Datastore
func (_d DatastoreWithPrometheus) SetOrderPaid(ctx context.Context, up1 *uuid.UUID) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "SetOrderPaid", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SetOrderPaid(ctx, up1)
}

// SetOrderTrialDays implements Datastore
func (_d DatastoreWithPrometheus) SetOrderTrialDays(ctx context.Context, orderID *uuid.UUID, days int64) (op1 *Order, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "SetOrderTrialDays", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SetOrderTrialDays(ctx, orderID, days)
}

// UpdateOrder implements Datastore
func (_d DatastoreWithPrometheus) UpdateOrder(orderID uuid.UUID, status string) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateOrder", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateOrder(orderID, status)
}

// UpdateOrderMetadata implements Datastore
func (_d DatastoreWithPrometheus) UpdateOrderMetadata(orderID uuid.UUID, key string, value string) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateOrderMetadata", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateOrderMetadata(orderID, key, value)
}

// UpdateSigningOrderRequestOutboxTx implements Datastore
func (_d DatastoreWithPrometheus) UpdateSigningOrderRequestOutboxTx(ctx context.Context, tx *sqlx.Tx, requestID uuid.UUID, completedAt time.Time) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateSigningOrderRequestOutboxTx", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateSigningOrderRequestOutboxTx(ctx, tx, requestID, completedAt)
}

// UpdateTransaction implements Datastore
func (_d DatastoreWithPrometheus) UpdateTransaction(orderID uuid.UUID, externalTransactionID string, status string, currency string, kind string, amount decimal.Decimal) (tp1 *Transaction, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		datastoreDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateTransaction", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateTransaction(orderID, externalTransactionID, status, currency, kind, amount)
}
