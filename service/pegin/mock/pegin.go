// Code generated by MockGen. DO NOT EDIT.
// Source: pegin.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cryptogarageinc/cfd-go/types"
	gomock "go.uber.org/mock/gomock"
)

// MockPegin is a mock of Pegin interface.
type MockPegin struct {
	ctrl     *gomock.Controller
	recorder *MockPeginMockRecorder
}

// MockPeginMockRecorder is the mock recorder for MockPegin.
type MockPeginMockRecorder struct {
	mock *MockPegin
}

// NewMockPegin creates a new mock instance.
func NewMockPegin(ctrl *gomock.Controller) *MockPegin {
	mock := &MockPegin{ctrl: ctrl}
	mock.recorder = &MockPeginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPegin) EXPECT() *MockPeginMockRecorder {
	return m.recorder
}

// CreatePeginAddress mocks base method.
func (m *MockPegin) CreatePeginAddress(addressType types.AddressType, pubkey *types.Pubkey, fedpegScript *types.Script) (*types.Address, *types.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeginAddress", addressType, pubkey, fedpegScript)
	ret0, _ := ret[0].(*types.Address)
	ret1, _ := ret[1].(*types.Script)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePeginAddress indicates an expected call of CreatePeginAddress.
func (mr *MockPeginMockRecorder) CreatePeginAddress(addressType, pubkey, fedpegScript interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeginAddress", reflect.TypeOf((*MockPegin)(nil).CreatePeginAddress), addressType, pubkey, fedpegScript)
}

// CreatePeginTransaction mocks base method.
func (m *MockPegin) CreatePeginTransaction(peginOutPoint *types.OutPoint, peginData *types.InputPeginData, utxoList []*types.ElementsUtxoData, sendList []*types.InputConfidentialTxOut, changeAddress *string, option *types.PeginTxOption) (*types.ConfidentialTx, *types.ConfidentialTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeginTransaction", peginOutPoint, peginData, utxoList, sendList, changeAddress, option)
	ret0, _ := ret[0].(*types.ConfidentialTx)
	ret1, _ := ret[1].(*types.ConfidentialTx)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreatePeginTransaction indicates an expected call of CreatePeginTransaction.
func (mr *MockPeginMockRecorder) CreatePeginTransaction(peginOutPoint, peginData, utxoList, sendList, changeAddress, option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeginTransaction", reflect.TypeOf((*MockPegin)(nil).CreatePeginTransaction), peginOutPoint, peginData, utxoList, sendList, changeAddress, option)
}

// GetPeginUtxoData mocks base method.
func (m *MockPegin) GetPeginUtxoData(proposalTx *types.ConfidentialTx, peginOutPoint *types.OutPoint, pubkey *types.Pubkey) (*types.ElementsUtxoData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeginUtxoData", proposalTx, peginOutPoint, pubkey)
	ret0, _ := ret[0].(*types.ElementsUtxoData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeginUtxoData indicates an expected call of GetPeginUtxoData.
func (mr *MockPeginMockRecorder) GetPeginUtxoData(proposalTx, peginOutPoint, pubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeginUtxoData", reflect.TypeOf((*MockPegin)(nil).GetPeginUtxoData), proposalTx, peginOutPoint, pubkey)
}

// GetPubkeyFromAccountExtPubkey mocks base method.
func (m *MockPegin) GetPubkeyFromAccountExtPubkey(accountExtPubkey *types.ExtPubkey, bip32Path string) (*types.Pubkey, *types.ExtPubkey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkeyFromAccountExtPubkey", accountExtPubkey, bip32Path)
	ret0, _ := ret[0].(*types.Pubkey)
	ret1, _ := ret[1].(*types.ExtPubkey)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPubkeyFromAccountExtPubkey indicates an expected call of GetPubkeyFromAccountExtPubkey.
func (mr *MockPeginMockRecorder) GetPubkeyFromAccountExtPubkey(accountExtPubkey, bip32Path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkeyFromAccountExtPubkey", reflect.TypeOf((*MockPegin)(nil).GetPubkeyFromAccountExtPubkey), accountExtPubkey, bip32Path)
}

// VerifyPubkeySignature mocks base method.
func (m *MockPegin) VerifyPubkeySignature(proposalTx *types.ConfidentialTx, utxoData *types.ElementsUtxoData, signature *types.ByteData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyPubkeySignature", proposalTx, utxoData, signature)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyPubkeySignature indicates an expected call of VerifyPubkeySignature.
func (mr *MockPeginMockRecorder) VerifyPubkeySignature(proposalTx, utxoData, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyPubkeySignature", reflect.TypeOf((*MockPegin)(nil).VerifyPubkeySignature), proposalTx, utxoData, signature)
}
