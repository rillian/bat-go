// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/skus/datastore.go

// Package skus is a generated GoMock package.
package skus

import (
	context "context"
	reflect "reflect"
	time "time"

	inputs "github.com/brave-intl/bat-go/libs/inputs"
	v4 "github.com/golang-migrate/migrate/v4"
	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	go_uuid "github.com/satori/go.uuid"
	decimal "github.com/shopspring/decimal"
)

// MockDatastore is a mock of Datastore interface.
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore.
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance.
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// AppendOrderMetadata mocks base method.
func (m *MockDatastore) AppendOrderMetadata(arg0 context.Context, arg1 *go_uuid.UUID, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendOrderMetadata", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendOrderMetadata indicates an expected call of AppendOrderMetadata.
func (mr *MockDatastoreMockRecorder) AppendOrderMetadata(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendOrderMetadata", reflect.TypeOf((*MockDatastore)(nil).AppendOrderMetadata), arg0, arg1, arg2, arg3)
}

// BeginTx mocks base method.
func (m *MockDatastore) BeginTx() (*sqlx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginTx")
	ret0, _ := ret[0].(*sqlx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginTx indicates an expected call of BeginTx.
func (mr *MockDatastoreMockRecorder) BeginTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginTx", reflect.TypeOf((*MockDatastore)(nil).BeginTx))
}

// CheckExpiredCheckoutSession mocks base method.
func (m *MockDatastore) CheckExpiredCheckoutSession(arg0 go_uuid.UUID) (bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckExpiredCheckoutSession", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CheckExpiredCheckoutSession indicates an expected call of CheckExpiredCheckoutSession.
func (mr *MockDatastoreMockRecorder) CheckExpiredCheckoutSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckExpiredCheckoutSession", reflect.TypeOf((*MockDatastore)(nil).CheckExpiredCheckoutSession), arg0)
}

// CommitVote mocks base method.
func (m *MockDatastore) CommitVote(ctx context.Context, vr VoteRecord, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitVote", ctx, vr, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitVote indicates an expected call of CommitVote.
func (mr *MockDatastoreMockRecorder) CommitVote(ctx, vr, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitVote", reflect.TypeOf((*MockDatastore)(nil).CommitVote), ctx, vr, tx)
}

// CreateKey mocks base method.
func (m *MockDatastore) CreateKey(merchant, name, encryptedSecretKey, nonce string) (*Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateKey", merchant, name, encryptedSecretKey, nonce)
	ret0, _ := ret[0].(*Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateKey indicates an expected call of CreateKey.
func (mr *MockDatastoreMockRecorder) CreateKey(merchant, name, encryptedSecretKey, nonce interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateKey", reflect.TypeOf((*MockDatastore)(nil).CreateKey), merchant, name, encryptedSecretKey, nonce)
}

// CreateOrder mocks base method.
func (m *MockDatastore) CreateOrder(totalPrice decimal.Decimal, merchantID, status, currency, location string, validFor *time.Duration, orderItems []OrderItem, allowedPaymentMethods *Methods) (*Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", totalPrice, merchantID, status, currency, location, validFor, orderItems, allowedPaymentMethods)
	ret0, _ := ret[0].(*Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockDatastoreMockRecorder) CreateOrder(totalPrice, merchantID, status, currency, location, validFor, orderItems, allowedPaymentMethods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockDatastore)(nil).CreateOrder), totalPrice, merchantID, status, currency, location, validFor, orderItems, allowedPaymentMethods)
}

// CreateTransaction mocks base method.
func (m *MockDatastore) CreateTransaction(orderID go_uuid.UUID, externalTransactionID, status, currency, kind string, amount decimal.Decimal) (*Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", orderID, externalTransactionID, status, currency, kind, amount)
	ret0, _ := ret[0].(*Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockDatastoreMockRecorder) CreateTransaction(orderID, externalTransactionID, status, currency, kind, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockDatastore)(nil).CreateTransaction), orderID, externalTransactionID, status, currency, kind, amount)
}

// DeleteKey mocks base method.
func (m *MockDatastore) DeleteKey(id go_uuid.UUID, delaySeconds int) (*Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteKey", id, delaySeconds)
	ret0, _ := ret[0].(*Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteKey indicates an expected call of DeleteKey.
func (mr *MockDatastoreMockRecorder) DeleteKey(id, delaySeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteKey", reflect.TypeOf((*MockDatastore)(nil).DeleteKey), id, delaySeconds)
}

// DeleteOrderCreds mocks base method.
func (m *MockDatastore) DeleteOrderCreds(orderID go_uuid.UUID, isSigned bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrderCreds", orderID, isSigned)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrderCreds indicates an expected call of DeleteOrderCreds.
func (mr *MockDatastoreMockRecorder) DeleteOrderCreds(orderID, isSigned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrderCreds", reflect.TypeOf((*MockDatastore)(nil).DeleteOrderCreds), orderID, isSigned)
}

// GetIssuer mocks base method.
func (m *MockDatastore) GetIssuer(merchantID string) (*Issuer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuer", merchantID)
	ret0, _ := ret[0].(*Issuer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuer indicates an expected call of GetIssuer.
func (mr *MockDatastoreMockRecorder) GetIssuer(merchantID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuer", reflect.TypeOf((*MockDatastore)(nil).GetIssuer), merchantID)
}

// GetIssuerByPublicKey mocks base method.
func (m *MockDatastore) GetIssuerByPublicKey(publicKey string) (*Issuer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuerByPublicKey", publicKey)
	ret0, _ := ret[0].(*Issuer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuerByPublicKey indicates an expected call of GetIssuerByPublicKey.
func (mr *MockDatastoreMockRecorder) GetIssuerByPublicKey(publicKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuerByPublicKey", reflect.TypeOf((*MockDatastore)(nil).GetIssuerByPublicKey), publicKey)
}

// GetKey mocks base method.
func (m *MockDatastore) GetKey(id go_uuid.UUID, showExpired bool) (*Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", id, showExpired)
	ret0, _ := ret[0].(*Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey.
func (mr *MockDatastoreMockRecorder) GetKey(id, showExpired interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockDatastore)(nil).GetKey), id, showExpired)
}

// GetKeysByMerchant mocks base method.
func (m *MockDatastore) GetKeysByMerchant(merchant string, showExpired bool) (*[]Key, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeysByMerchant", merchant, showExpired)
	ret0, _ := ret[0].(*[]Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKeysByMerchant indicates an expected call of GetKeysByMerchant.
func (mr *MockDatastoreMockRecorder) GetKeysByMerchant(merchant, showExpired interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeysByMerchant", reflect.TypeOf((*MockDatastore)(nil).GetKeysByMerchant), merchant, showExpired)
}

// GetOrder mocks base method.
func (m *MockDatastore) GetOrder(orderID go_uuid.UUID) (*Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", orderID)
	ret0, _ := ret[0].(*Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockDatastoreMockRecorder) GetOrder(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockDatastore)(nil).GetOrder), orderID)
}

// GetOrderByExternalID mocks base method.
func (m *MockDatastore) GetOrderByExternalID(externalID string) (*Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByExternalID", externalID)
	ret0, _ := ret[0].(*Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByExternalID indicates an expected call of GetOrderByExternalID.
func (mr *MockDatastoreMockRecorder) GetOrderByExternalID(externalID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByExternalID", reflect.TypeOf((*MockDatastore)(nil).GetOrderByExternalID), externalID)
}

// GetOrderCreds mocks base method.
func (m *MockDatastore) GetOrderCreds(orderID go_uuid.UUID, isSigned bool) (*[]OrderCreds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderCreds", orderID, isSigned)
	ret0, _ := ret[0].(*[]OrderCreds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderCreds indicates an expected call of GetOrderCreds.
func (mr *MockDatastoreMockRecorder) GetOrderCreds(orderID, isSigned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderCreds", reflect.TypeOf((*MockDatastore)(nil).GetOrderCreds), orderID, isSigned)
}

// GetOrderCredsByItemID mocks base method.
func (m *MockDatastore) GetOrderCredsByItemID(orderID, itemID go_uuid.UUID, isSigned bool) (*OrderCreds, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderCredsByItemID", orderID, itemID, isSigned)
	ret0, _ := ret[0].(*OrderCreds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderCredsByItemID indicates an expected call of GetOrderCredsByItemID.
func (mr *MockDatastoreMockRecorder) GetOrderCredsByItemID(orderID, itemID, isSigned interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderCredsByItemID", reflect.TypeOf((*MockDatastore)(nil).GetOrderCredsByItemID), orderID, itemID, isSigned)
}

// GetPagedMerchantTransactions mocks base method.
func (m *MockDatastore) GetPagedMerchantTransactions(ctx context.Context, merchantID go_uuid.UUID, pagination *inputs.Pagination) (*[]Transaction, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPagedMerchantTransactions", ctx, merchantID, pagination)
	ret0, _ := ret[0].(*[]Transaction)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPagedMerchantTransactions indicates an expected call of GetPagedMerchantTransactions.
func (mr *MockDatastoreMockRecorder) GetPagedMerchantTransactions(ctx, merchantID, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPagedMerchantTransactions", reflect.TypeOf((*MockDatastore)(nil).GetPagedMerchantTransactions), ctx, merchantID, pagination)
}

// GetSumForTransactions mocks base method.
func (m *MockDatastore) GetSumForTransactions(orderID go_uuid.UUID) (decimal.Decimal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSumForTransactions", orderID)
	ret0, _ := ret[0].(decimal.Decimal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSumForTransactions indicates an expected call of GetSumForTransactions.
func (mr *MockDatastoreMockRecorder) GetSumForTransactions(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSumForTransactions", reflect.TypeOf((*MockDatastore)(nil).GetSumForTransactions), orderID)
}

// GetTransaction mocks base method.
func (m *MockDatastore) GetTransaction(externalTransactionID string) (*Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", externalTransactionID)
	ret0, _ := ret[0].(*Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockDatastoreMockRecorder) GetTransaction(externalTransactionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockDatastore)(nil).GetTransaction), externalTransactionID)
}

// GetTransactions mocks base method.
func (m *MockDatastore) GetTransactions(orderID go_uuid.UUID) (*[]Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", orderID)
	ret0, _ := ret[0].(*[]Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockDatastoreMockRecorder) GetTransactions(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockDatastore)(nil).GetTransactions), orderID)
}

// GetUncommittedVotesForUpdate mocks base method.
func (m *MockDatastore) GetUncommittedVotesForUpdate(ctx context.Context) (*sqlx.Tx, []*VoteRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUncommittedVotesForUpdate", ctx)
	ret0, _ := ret[0].(*sqlx.Tx)
	ret1, _ := ret[1].([]*VoteRecord)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUncommittedVotesForUpdate indicates an expected call of GetUncommittedVotesForUpdate.
func (mr *MockDatastoreMockRecorder) GetUncommittedVotesForUpdate(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUncommittedVotesForUpdate", reflect.TypeOf((*MockDatastore)(nil).GetUncommittedVotesForUpdate), ctx)
}

// InsertIssuer mocks base method.
func (m *MockDatastore) InsertIssuer(issuer *Issuer) (*Issuer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertIssuer", issuer)
	ret0, _ := ret[0].(*Issuer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertIssuer indicates an expected call of InsertIssuer.
func (mr *MockDatastoreMockRecorder) InsertIssuer(issuer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertIssuer", reflect.TypeOf((*MockDatastore)(nil).InsertIssuer), issuer)
}

// InsertOrderCreds mocks base method.
func (m *MockDatastore) InsertOrderCreds(creds *OrderCreds) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrderCreds", creds)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrderCreds indicates an expected call of InsertOrderCreds.
func (mr *MockDatastoreMockRecorder) InsertOrderCreds(creds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrderCreds", reflect.TypeOf((*MockDatastore)(nil).InsertOrderCreds), creds)
}

// InsertVote mocks base method.
func (m *MockDatastore) InsertVote(ctx context.Context, vr VoteRecord) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertVote", ctx, vr)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertVote indicates an expected call of InsertVote.
func (mr *MockDatastoreMockRecorder) InsertVote(ctx, vr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertVote", reflect.TypeOf((*MockDatastore)(nil).InsertVote), ctx, vr)
}

// IsStripeSub mocks base method.
func (m *MockDatastore) IsStripeSub(arg0 go_uuid.UUID) (bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsStripeSub", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsStripeSub indicates an expected call of IsStripeSub.
func (mr *MockDatastoreMockRecorder) IsStripeSub(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStripeSub", reflect.TypeOf((*MockDatastore)(nil).IsStripeSub), arg0)
}

// MarkVoteErrored mocks base method.
func (m *MockDatastore) MarkVoteErrored(ctx context.Context, vr VoteRecord, tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkVoteErrored", ctx, vr, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkVoteErrored indicates an expected call of MarkVoteErrored.
func (mr *MockDatastoreMockRecorder) MarkVoteErrored(ctx, vr, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkVoteErrored", reflect.TypeOf((*MockDatastore)(nil).MarkVoteErrored), ctx, vr, tx)
}

// Migrate mocks base method.
func (m *MockDatastore) Migrate(arg0 ...uint) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Migrate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Migrate indicates an expected call of Migrate.
func (mr *MockDatastoreMockRecorder) Migrate(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Migrate", reflect.TypeOf((*MockDatastore)(nil).Migrate), arg0...)
}

// NewMigrate mocks base method.
func (m *MockDatastore) NewMigrate() (*v4.Migrate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewMigrate")
	ret0, _ := ret[0].(*v4.Migrate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewMigrate indicates an expected call of NewMigrate.
func (mr *MockDatastoreMockRecorder) NewMigrate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewMigrate", reflect.TypeOf((*MockDatastore)(nil).NewMigrate))
}

// RawDB mocks base method.
func (m *MockDatastore) RawDB() *sqlx.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawDB")
	ret0, _ := ret[0].(*sqlx.DB)
	return ret0
}

// RawDB indicates an expected call of RawDB.
func (mr *MockDatastoreMockRecorder) RawDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawDB", reflect.TypeOf((*MockDatastore)(nil).RawDB))
}

// RenewOrder mocks base method.
func (m *MockDatastore) RenewOrder(ctx context.Context, orderID go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewOrder", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenewOrder indicates an expected call of RenewOrder.
func (mr *MockDatastoreMockRecorder) RenewOrder(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewOrder", reflect.TypeOf((*MockDatastore)(nil).RenewOrder), ctx, orderID)
}

// RollbackTx mocks base method.
func (m *MockDatastore) RollbackTx(tx *sqlx.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RollbackTx", tx)
}

// RollbackTx indicates an expected call of RollbackTx.
func (mr *MockDatastoreMockRecorder) RollbackTx(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTx", reflect.TypeOf((*MockDatastore)(nil).RollbackTx), tx)
}

// RollbackTxAndHandle mocks base method.
func (m *MockDatastore) RollbackTxAndHandle(tx *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RollbackTxAndHandle", tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// RollbackTxAndHandle indicates an expected call of RollbackTxAndHandle.
func (mr *MockDatastoreMockRecorder) RollbackTxAndHandle(tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RollbackTxAndHandle", reflect.TypeOf((*MockDatastore)(nil).RollbackTxAndHandle), tx)
}

// RunNextOrderJob mocks base method.
func (m *MockDatastore) RunNextOrderJob(ctx context.Context, worker OrderWorker) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunNextOrderJob", ctx, worker)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunNextOrderJob indicates an expected call of RunNextOrderJob.
func (mr *MockDatastoreMockRecorder) RunNextOrderJob(ctx, worker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunNextOrderJob", reflect.TypeOf((*MockDatastore)(nil).RunNextOrderJob), ctx, worker)
}

// SetOrderPaid mocks base method.
func (m *MockDatastore) SetOrderPaid(arg0 context.Context, arg1 *go_uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrderPaid", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetOrderPaid indicates an expected call of SetOrderPaid.
func (mr *MockDatastoreMockRecorder) SetOrderPaid(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrderPaid", reflect.TypeOf((*MockDatastore)(nil).SetOrderPaid), arg0, arg1)
}

// SetOrderTrialDays mocks base method.
func (m *MockDatastore) SetOrderTrialDays(ctx context.Context, orderID *go_uuid.UUID, days int64) (*Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetOrderTrialDays", ctx, orderID, days)
	ret0, _ := ret[0].(*Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetOrderTrialDays indicates an expected call of SetOrderTrialDays.
func (mr *MockDatastoreMockRecorder) SetOrderTrialDays(ctx, orderID, days interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOrderTrialDays", reflect.TypeOf((*MockDatastore)(nil).SetOrderTrialDays), ctx, orderID, days)
}

// UpdateOrder mocks base method.
func (m *MockDatastore) UpdateOrder(orderID go_uuid.UUID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", orderID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockDatastoreMockRecorder) UpdateOrder(orderID, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockDatastore)(nil).UpdateOrder), orderID, status)
}

// UpdateOrderMetadata mocks base method.
func (m *MockDatastore) UpdateOrderMetadata(orderID go_uuid.UUID, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrderMetadata", orderID, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrderMetadata indicates an expected call of UpdateOrderMetadata.
func (mr *MockDatastoreMockRecorder) UpdateOrderMetadata(orderID, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrderMetadata", reflect.TypeOf((*MockDatastore)(nil).UpdateOrderMetadata), orderID, key, value)
}

// UpdateTransaction mocks base method.
func (m *MockDatastore) UpdateTransaction(orderID go_uuid.UUID, externalTransactionID, status, currency, kind string, amount decimal.Decimal) (*Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", orderID, externalTransactionID, status, currency, kind, amount)
	ret0, _ := ret[0].(*Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockDatastoreMockRecorder) UpdateTransaction(orderID, externalTransactionID, status, currency, kind, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockDatastore)(nil).UpdateTransaction), orderID, externalTransactionID, status, currency, kind, amount)
}
