// Code generated by MockGen. DO NOT EDIT.
// Source: api_key.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	domain "github.com/bucketeer-io/bucketeer/pkg/account/domain"
	mysql "github.com/bucketeer-io/bucketeer/pkg/storage/v2/mysql"
	account "github.com/bucketeer-io/bucketeer/proto/account"
)

// MockAPIKeyStorage is a mock of APIKeyStorage interface.
type MockAPIKeyStorage struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyStorageMockRecorder
}

// MockAPIKeyStorageMockRecorder is the mock recorder for MockAPIKeyStorage.
type MockAPIKeyStorageMockRecorder struct {
	mock *MockAPIKeyStorage
}

// NewMockAPIKeyStorage creates a new mock instance.
func NewMockAPIKeyStorage(ctrl *gomock.Controller) *MockAPIKeyStorage {
	mock := &MockAPIKeyStorage{ctrl: ctrl}
	mock.recorder = &MockAPIKeyStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyStorage) EXPECT() *MockAPIKeyStorageMockRecorder {
	return m.recorder
}

// CreateAPIKey mocks base method.
func (m *MockAPIKeyStorage) CreateAPIKey(ctx context.Context, k *domain.APIKey, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAPIKey", ctx, k, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAPIKey indicates an expected call of CreateAPIKey.
func (mr *MockAPIKeyStorageMockRecorder) CreateAPIKey(ctx, k, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAPIKey", reflect.TypeOf((*MockAPIKeyStorage)(nil).CreateAPIKey), ctx, k, environmentNamespace)
}

// GetAPIKey mocks base method.
func (m *MockAPIKeyStorage) GetAPIKey(ctx context.Context, id, environmentNamespace string) (*domain.APIKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAPIKey", ctx, id, environmentNamespace)
	ret0, _ := ret[0].(*domain.APIKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAPIKey indicates an expected call of GetAPIKey.
func (mr *MockAPIKeyStorageMockRecorder) GetAPIKey(ctx, id, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAPIKey", reflect.TypeOf((*MockAPIKeyStorage)(nil).GetAPIKey), ctx, id, environmentNamespace)
}

// ListAPIKeys mocks base method.
func (m *MockAPIKeyStorage) ListAPIKeys(ctx context.Context, whereParts []mysql.WherePart, orders []*mysql.Order, limit, offset int) ([]*account.APIKey, int, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAPIKeys", ctx, whereParts, orders, limit, offset)
	ret0, _ := ret[0].([]*account.APIKey)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// ListAPIKeys indicates an expected call of ListAPIKeys.
func (mr *MockAPIKeyStorageMockRecorder) ListAPIKeys(ctx, whereParts, orders, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAPIKeys", reflect.TypeOf((*MockAPIKeyStorage)(nil).ListAPIKeys), ctx, whereParts, orders, limit, offset)
}

// UpdateAPIKey mocks base method.
func (m *MockAPIKeyStorage) UpdateAPIKey(ctx context.Context, k *domain.APIKey, environmentNamespace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAPIKey", ctx, k, environmentNamespace)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAPIKey indicates an expected call of UpdateAPIKey.
func (mr *MockAPIKeyStorageMockRecorder) UpdateAPIKey(ctx, k, environmentNamespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAPIKey", reflect.TypeOf((*MockAPIKeyStorage)(nil).UpdateAPIKey), ctx, k, environmentNamespace)
}
