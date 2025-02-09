// Code generated by MockGen. DO NOT EDIT.
// Source: ./libs/clients/ratios/client.go

// Package mock_ratios is a generated GoMock package.
package mock_ratios

import (
	context "context"
	reflect "reflect"

	ratios "github.com/brave-intl/bat-go/libs/clients/ratios"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// FetchRate mocks base method.
func (m *MockClient) FetchRate(ctx context.Context, base, currency string) (*ratios.RateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchRate", ctx, base, currency)
	ret0, _ := ret[0].(*ratios.RateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchRate indicates an expected call of FetchRate.
func (mr *MockClientMockRecorder) FetchRate(ctx, base, currency interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchRate", reflect.TypeOf((*MockClient)(nil).FetchRate), ctx, base, currency)
}
