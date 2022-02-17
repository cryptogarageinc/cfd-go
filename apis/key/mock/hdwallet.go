// Code generated by MockGen. DO NOT EDIT.
// Source: hdwallet.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	key "github.com/cryptogarageinc/cfd-go/apis/key"
	types "github.com/cryptogarageinc/cfd-go/types"
	gomock "github.com/golang/mock/gomock"
)

// MockExtPubkeyApi is a mock of ExtPubkeyApi interface.
type MockExtPubkeyApi struct {
	ctrl     *gomock.Controller
	recorder *MockExtPubkeyApiMockRecorder
}

// MockExtPubkeyApiMockRecorder is the mock recorder for MockExtPubkeyApi.
type MockExtPubkeyApiMockRecorder struct {
	mock *MockExtPubkeyApi
}

// NewMockExtPubkeyApi creates a new mock instance.
func NewMockExtPubkeyApi(ctrl *gomock.Controller) *MockExtPubkeyApi {
	mock := &MockExtPubkeyApi{ctrl: ctrl}
	mock.recorder = &MockExtPubkeyApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtPubkeyApi) EXPECT() *MockExtPubkeyApiMockRecorder {
	return m.recorder
}

// ConvertToBip32 mocks base method.
func (m *MockExtPubkeyApi) ConvertToBip32(extPubkey *types.ExtPubkey) (*types.ExtPubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConvertToBip32", extPubkey)
	ret0, _ := ret[0].(*types.ExtPubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertToBip32 indicates an expected call of ConvertToBip32.
func (mr *MockExtPubkeyApiMockRecorder) ConvertToBip32(extPubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertToBip32", reflect.TypeOf((*MockExtPubkeyApi)(nil).ConvertToBip32), extPubkey)
}

// GetData mocks base method.
func (m *MockExtPubkeyApi) GetData(extPubkey *types.ExtPubkey) (*types.ExtkeyData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", extPubkey)
	ret0, _ := ret[0].(*types.ExtkeyData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockExtPubkeyApiMockRecorder) GetData(extPubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockExtPubkeyApi)(nil).GetData), extPubkey)
}

// GetExtPubkeyByPath mocks base method.
func (m *MockExtPubkeyApi) GetExtPubkeyByPath(extPubkey *types.ExtPubkey, bip32Path string) (*types.ExtPubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPubkeyByPath", extPubkey, bip32Path)
	ret0, _ := ret[0].(*types.ExtPubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPubkeyByPath indicates an expected call of GetExtPubkeyByPath.
func (mr *MockExtPubkeyApiMockRecorder) GetExtPubkeyByPath(extPubkey, bip32Path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPubkeyByPath", reflect.TypeOf((*MockExtPubkeyApi)(nil).GetExtPubkeyByPath), extPubkey, bip32Path)
}

// GetPubkey mocks base method.
func (m *MockExtPubkeyApi) GetPubkey(extPubkey *types.ExtPubkey) (*types.Pubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkey", extPubkey)
	ret0, _ := ret[0].(*types.Pubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubkey indicates an expected call of GetPubkey.
func (mr *MockExtPubkeyApiMockRecorder) GetPubkey(extPubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkey", reflect.TypeOf((*MockExtPubkeyApi)(nil).GetPubkey), extPubkey)
}

// MockExtPrivkeyApi is a mock of ExtPrivkeyApi interface.
type MockExtPrivkeyApi struct {
	ctrl     *gomock.Controller
	recorder *MockExtPrivkeyApiMockRecorder
}

// MockExtPrivkeyApiMockRecorder is the mock recorder for MockExtPrivkeyApi.
type MockExtPrivkeyApiMockRecorder struct {
	mock *MockExtPrivkeyApi
}

// NewMockExtPrivkeyApi creates a new mock instance.
func NewMockExtPrivkeyApi(ctrl *gomock.Controller) *MockExtPrivkeyApi {
	mock := &MockExtPrivkeyApi{ctrl: ctrl}
	mock.recorder = &MockExtPrivkeyApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExtPrivkeyApi) EXPECT() *MockExtPrivkeyApiMockRecorder {
	return m.recorder
}

// GetData mocks base method.
func (m *MockExtPrivkeyApi) GetData(extPrivkey *types.ExtPrivkey) (*types.ExtkeyData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData", extPrivkey)
	ret0, _ := ret[0].(*types.ExtkeyData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetData indicates an expected call of GetData.
func (mr *MockExtPrivkeyApiMockRecorder) GetData(extPrivkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockExtPrivkeyApi)(nil).GetData), extPrivkey)
}

// GetExtPrivkeyByPath mocks base method.
func (m *MockExtPrivkeyApi) GetExtPrivkeyByPath(extPrivkey *types.ExtPrivkey, bip32Path string) (*types.ExtPrivkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPrivkeyByPath", extPrivkey, bip32Path)
	ret0, _ := ret[0].(*types.ExtPrivkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPrivkeyByPath indicates an expected call of GetExtPrivkeyByPath.
func (mr *MockExtPrivkeyApiMockRecorder) GetExtPrivkeyByPath(extPrivkey, bip32Path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPrivkeyByPath", reflect.TypeOf((*MockExtPrivkeyApi)(nil).GetExtPrivkeyByPath), extPrivkey, bip32Path)
}

// GetExtPubkey mocks base method.
func (m *MockExtPrivkeyApi) GetExtPubkey(extPrivkey *types.ExtPrivkey) (*types.ExtPubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPubkey", extPrivkey)
	ret0, _ := ret[0].(*types.ExtPubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPubkey indicates an expected call of GetExtPubkey.
func (mr *MockExtPrivkeyApiMockRecorder) GetExtPubkey(extPrivkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPubkey", reflect.TypeOf((*MockExtPrivkeyApi)(nil).GetExtPubkey), extPrivkey)
}

// GetPrivkey mocks base method.
func (m *MockExtPrivkeyApi) GetPrivkey(extPrivkey *types.ExtPrivkey) (*types.Privkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivkey", extPrivkey)
	ret0, _ := ret[0].(*types.Privkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivkey indicates an expected call of GetPrivkey.
func (mr *MockExtPrivkeyApiMockRecorder) GetPrivkey(extPrivkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivkey", reflect.TypeOf((*MockExtPrivkeyApi)(nil).GetPrivkey), extPrivkey)
}

// GetPubkey mocks base method.
func (m *MockExtPrivkeyApi) GetPubkey(extPrivkey *types.ExtPrivkey) (*types.Pubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkey", extPrivkey)
	ret0, _ := ret[0].(*types.Pubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubkey indicates an expected call of GetPubkey.
func (mr *MockExtPrivkeyApiMockRecorder) GetPubkey(extPrivkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkey", reflect.TypeOf((*MockExtPrivkeyApi)(nil).GetPubkey), extPrivkey)
}

// MockHdWalletApi is a mock of HdWalletApi interface.
type MockHdWalletApi struct {
	ctrl     *gomock.Controller
	recorder *MockHdWalletApiMockRecorder
}

// MockHdWalletApiMockRecorder is the mock recorder for MockHdWalletApi.
type MockHdWalletApiMockRecorder struct {
	mock *MockHdWalletApi
}

// NewMockHdWalletApi creates a new mock instance.
func NewMockHdWalletApi(ctrl *gomock.Controller) *MockHdWalletApi {
	mock := &MockHdWalletApi{ctrl: ctrl}
	mock.recorder = &MockHdWalletApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHdWalletApi) EXPECT() *MockHdWalletApiMockRecorder {
	return m.recorder
}

// CreateBip32Path mocks base method.
func (m *MockHdWalletApi) CreateBip32Path(items ...key.Bip32Item) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range items {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateBip32Path", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBip32Path indicates an expected call of CreateBip32Path.
func (mr *MockHdWalletApiMockRecorder) CreateBip32Path(items ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBip32Path", reflect.TypeOf((*MockHdWalletApi)(nil).CreateBip32Path), items...)
}

// GetExtPrivkey mocks base method.
func (m *MockHdWalletApi) GetExtPrivkey(seed *types.ByteData) (*types.ExtPrivkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPrivkey", seed)
	ret0, _ := ret[0].(*types.ExtPrivkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPrivkey indicates an expected call of GetExtPrivkey.
func (mr *MockHdWalletApiMockRecorder) GetExtPrivkey(seed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPrivkey", reflect.TypeOf((*MockHdWalletApi)(nil).GetExtPrivkey), seed)
}

// GetExtPrivkeyByPath mocks base method.
func (m *MockHdWalletApi) GetExtPrivkeyByPath(seed *types.ByteData, bip32Path string) (*types.ExtPrivkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPrivkeyByPath", seed, bip32Path)
	ret0, _ := ret[0].(*types.ExtPrivkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPrivkeyByPath indicates an expected call of GetExtPrivkeyByPath.
func (mr *MockHdWalletApiMockRecorder) GetExtPrivkeyByPath(seed, bip32Path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPrivkeyByPath", reflect.TypeOf((*MockHdWalletApi)(nil).GetExtPrivkeyByPath), seed, bip32Path)
}

// GetExtPubkeyByPath mocks base method.
func (m *MockHdWalletApi) GetExtPubkeyByPath(seed *types.ByteData, bip32Path string) (*types.ExtPubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtPubkeyByPath", seed, bip32Path)
	ret0, _ := ret[0].(*types.ExtPubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtPubkeyByPath indicates an expected call of GetExtPubkeyByPath.
func (mr *MockHdWalletApiMockRecorder) GetExtPubkeyByPath(seed, bip32Path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtPubkeyByPath", reflect.TypeOf((*MockHdWalletApi)(nil).GetExtPubkeyByPath), seed, bip32Path)
}

// GetMnemonicFromEntropy mocks base method.
func (m *MockHdWalletApi) GetMnemonicFromEntropy(entropy *types.ByteData, language string) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMnemonicFromEntropy", entropy, language)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMnemonicFromEntropy indicates an expected call of GetMnemonicFromEntropy.
func (mr *MockHdWalletApiMockRecorder) GetMnemonicFromEntropy(entropy, language interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMnemonicFromEntropy", reflect.TypeOf((*MockHdWalletApi)(nil).GetMnemonicFromEntropy), entropy, language)
}

// GetMnemonicFromEntropyEng mocks base method.
func (m *MockHdWalletApi) GetMnemonicFromEntropyEng(entropy *types.ByteData) (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMnemonicFromEntropyEng", entropy)
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMnemonicFromEntropyEng indicates an expected call of GetMnemonicFromEntropyEng.
func (mr *MockHdWalletApiMockRecorder) GetMnemonicFromEntropyEng(entropy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMnemonicFromEntropyEng", reflect.TypeOf((*MockHdWalletApi)(nil).GetMnemonicFromEntropyEng), entropy)
}

// GetSeedFromMnemonic mocks base method.
func (m *MockHdWalletApi) GetSeedFromMnemonic(mnemonic []string, language string) (*types.ByteData, *types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedFromMnemonic", mnemonic, language)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(*types.ByteData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSeedFromMnemonic indicates an expected call of GetSeedFromMnemonic.
func (mr *MockHdWalletApiMockRecorder) GetSeedFromMnemonic(mnemonic, language interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedFromMnemonic", reflect.TypeOf((*MockHdWalletApi)(nil).GetSeedFromMnemonic), mnemonic, language)
}

// GetSeedFromMnemonicAndPassphrase mocks base method.
func (m *MockHdWalletApi) GetSeedFromMnemonicAndPassphrase(mnemonic []string, language, passphrase string) (*types.ByteData, *types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedFromMnemonicAndPassphrase", mnemonic, language, passphrase)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(*types.ByteData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSeedFromMnemonicAndPassphrase indicates an expected call of GetSeedFromMnemonicAndPassphrase.
func (mr *MockHdWalletApiMockRecorder) GetSeedFromMnemonicAndPassphrase(mnemonic, language, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedFromMnemonicAndPassphrase", reflect.TypeOf((*MockHdWalletApi)(nil).GetSeedFromMnemonicAndPassphrase), mnemonic, language, passphrase)
}

// GetSeedFromMnemonicEng mocks base method.
func (m *MockHdWalletApi) GetSeedFromMnemonicEng(mnemonic []string) (*types.ByteData, *types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedFromMnemonicEng", mnemonic)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(*types.ByteData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSeedFromMnemonicEng indicates an expected call of GetSeedFromMnemonicEng.
func (mr *MockHdWalletApiMockRecorder) GetSeedFromMnemonicEng(mnemonic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedFromMnemonicEng", reflect.TypeOf((*MockHdWalletApi)(nil).GetSeedFromMnemonicEng), mnemonic)
}

// GetSeedFromMnemonicEngAndPassphrase mocks base method.
func (m *MockHdWalletApi) GetSeedFromMnemonicEngAndPassphrase(mnemonic []string, passphrase string) (*types.ByteData, *types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeedFromMnemonicEngAndPassphrase", mnemonic, passphrase)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(*types.ByteData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSeedFromMnemonicEngAndPassphrase indicates an expected call of GetSeedFromMnemonicEngAndPassphrase.
func (mr *MockHdWalletApiMockRecorder) GetSeedFromMnemonicEngAndPassphrase(mnemonic, passphrase interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeedFromMnemonicEngAndPassphrase", reflect.TypeOf((*MockHdWalletApi)(nil).GetSeedFromMnemonicEngAndPassphrase), mnemonic, passphrase)
}
