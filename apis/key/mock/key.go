// Code generated by MockGen. DO NOT EDIT.
// Source: key.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cryptogarageinc/cfd-go/types"
	gomock "go.uber.org/mock/gomock"
)

// MockPubkeyApi is a mock of PubkeyApi interface.
type MockPubkeyApi struct {
	ctrl     *gomock.Controller
	recorder *MockPubkeyApiMockRecorder
}

// MockPubkeyApiMockRecorder is the mock recorder for MockPubkeyApi.
type MockPubkeyApiMockRecorder struct {
	mock *MockPubkeyApi
}

// NewMockPubkeyApi creates a new mock instance.
func NewMockPubkeyApi(ctrl *gomock.Controller) *MockPubkeyApi {
	mock := &MockPubkeyApi{ctrl: ctrl}
	mock.recorder = &MockPubkeyApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPubkeyApi) EXPECT() *MockPubkeyApiMockRecorder {
	return m.recorder
}

// IsCompressed mocks base method.
func (m *MockPubkeyApi) IsCompressed(pubkey *types.Pubkey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCompressed", pubkey)
	ret0, _ := ret[0].(error)
	return ret0
}

// IsCompressed indicates an expected call of IsCompressed.
func (mr *MockPubkeyApiMockRecorder) IsCompressed(pubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCompressed", reflect.TypeOf((*MockPubkeyApi)(nil).IsCompressed), pubkey)
}

// Verify mocks base method.
func (m *MockPubkeyApi) Verify(pubkey *types.Pubkey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", pubkey)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockPubkeyApiMockRecorder) Verify(pubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockPubkeyApi)(nil).Verify), pubkey)
}

// VerifyEcSignature mocks base method.
func (m *MockPubkeyApi) VerifyEcSignature(pubkey *types.Pubkey, sighash, signature string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEcSignature", pubkey, sighash, signature)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyEcSignature indicates an expected call of VerifyEcSignature.
func (mr *MockPubkeyApiMockRecorder) VerifyEcSignature(pubkey, sighash, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEcSignature", reflect.TypeOf((*MockPubkeyApi)(nil).VerifyEcSignature), pubkey, sighash, signature)
}

// VerifyMessage mocks base method.
func (m *MockPubkeyApi) VerifyMessage(pubkey *types.Pubkey, signature, message, magic string) (*types.Pubkey, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyMessage", pubkey, signature, message, magic)
	ret0, _ := ret[0].(*types.Pubkey)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VerifyMessage indicates an expected call of VerifyMessage.
func (mr *MockPubkeyApiMockRecorder) VerifyMessage(pubkey, signature, message, magic interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyMessage", reflect.TypeOf((*MockPubkeyApi)(nil).VerifyMessage), pubkey, signature, message, magic)
}

// MockPrivkeyApi is a mock of PrivkeyApi interface.
type MockPrivkeyApi struct {
	ctrl     *gomock.Controller
	recorder *MockPrivkeyApiMockRecorder
}

// MockPrivkeyApiMockRecorder is the mock recorder for MockPrivkeyApi.
type MockPrivkeyApiMockRecorder struct {
	mock *MockPrivkeyApi
}

// NewMockPrivkeyApi creates a new mock instance.
func NewMockPrivkeyApi(ctrl *gomock.Controller) *MockPrivkeyApi {
	mock := &MockPrivkeyApi{ctrl: ctrl}
	mock.recorder = &MockPrivkeyApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivkeyApi) EXPECT() *MockPrivkeyApiMockRecorder {
	return m.recorder
}

// CreateEcSignature mocks base method.
func (m *MockPrivkeyApi) CreateEcSignature(privkey *types.Privkey, sighash *types.ByteData, sighashType *types.SigHashType) (*types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEcSignature", privkey, sighash, sighashType)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEcSignature indicates an expected call of CreateEcSignature.
func (mr *MockPrivkeyApiMockRecorder) CreateEcSignature(privkey, sighash, sighashType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEcSignature", reflect.TypeOf((*MockPrivkeyApi)(nil).CreateEcSignature), privkey, sighash, sighashType)
}

// CreateEcSignatureGrindR mocks base method.
func (m *MockPrivkeyApi) CreateEcSignatureGrindR(privkey *types.Privkey, sighash *types.ByteData, sighashType *types.SigHashType, grindR bool) (*types.ByteData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEcSignatureGrindR", privkey, sighash, sighashType, grindR)
	ret0, _ := ret[0].(*types.ByteData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEcSignatureGrindR indicates an expected call of CreateEcSignatureGrindR.
func (mr *MockPrivkeyApiMockRecorder) CreateEcSignatureGrindR(privkey, sighash, sighashType, grindR interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEcSignatureGrindR", reflect.TypeOf((*MockPrivkeyApi)(nil).CreateEcSignatureGrindR), privkey, sighash, sighashType, grindR)
}

// GetPrivkeyFromWif mocks base method.
func (m *MockPrivkeyApi) GetPrivkeyFromWif(wif string) (*types.Privkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivkeyFromWif", wif)
	ret0, _ := ret[0].(*types.Privkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivkeyFromWif indicates an expected call of GetPrivkeyFromWif.
func (mr *MockPrivkeyApiMockRecorder) GetPrivkeyFromWif(wif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivkeyFromWif", reflect.TypeOf((*MockPrivkeyApi)(nil).GetPrivkeyFromWif), wif)
}

// GetPubkey mocks base method.
func (m *MockPrivkeyApi) GetPubkey(privkey *types.Privkey) (*types.Pubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkey", privkey)
	ret0, _ := ret[0].(*types.Pubkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubkey indicates an expected call of GetPubkey.
func (mr *MockPrivkeyApiMockRecorder) GetPubkey(privkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkey", reflect.TypeOf((*MockPrivkeyApi)(nil).GetPubkey), privkey)
}

// GetWifFromHex mocks base method.
func (m *MockPrivkeyApi) GetWifFromHex(privkeyHex string) (*types.Privkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWifFromHex", privkeyHex)
	ret0, _ := ret[0].(*types.Privkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWifFromHex indicates an expected call of GetWifFromHex.
func (mr *MockPrivkeyApiMockRecorder) GetWifFromHex(privkeyHex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWifFromHex", reflect.TypeOf((*MockPrivkeyApi)(nil).GetWifFromHex), privkeyHex)
}

// GetWifFromHexWithCompressedPubkey mocks base method.
func (m *MockPrivkeyApi) GetWifFromHexWithCompressedPubkey(privkeyHex string, compressedPubkey bool) (*types.Privkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWifFromHexWithCompressedPubkey", privkeyHex, compressedPubkey)
	ret0, _ := ret[0].(*types.Privkey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWifFromHexWithCompressedPubkey indicates an expected call of GetWifFromHexWithCompressedPubkey.
func (mr *MockPrivkeyApiMockRecorder) GetWifFromHexWithCompressedPubkey(privkeyHex, compressedPubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWifFromHexWithCompressedPubkey", reflect.TypeOf((*MockPrivkeyApi)(nil).GetWifFromHexWithCompressedPubkey), privkeyHex, compressedPubkey)
}

// HasWif mocks base method.
func (m *MockPrivkeyApi) HasWif(wif string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasWif", wif)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasWif indicates an expected call of HasWif.
func (mr *MockPrivkeyApiMockRecorder) HasWif(wif interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasWif", reflect.TypeOf((*MockPrivkeyApi)(nil).HasWif), wif)
}

// SignMessage mocks base method.
func (m *MockPrivkeyApi) SignMessage(privkey *types.Privkey, message, magic string, isOutputBase64 bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignMessage", privkey, message, magic, isOutputBase64)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignMessage indicates an expected call of SignMessage.
func (mr *MockPrivkeyApiMockRecorder) SignMessage(privkey, message, magic, isOutputBase64 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignMessage", reflect.TypeOf((*MockPrivkeyApi)(nil).SignMessage), privkey, message, magic, isOutputBase64)
}
