// Code generated by MockGen. DO NOT EDIT.
// Source: script.go
//
// Generated by this command:
//
//	mockgen -source script.go -destination mock/script.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cryptogarageinc/cfd-go/types"
	gomock "go.uber.org/mock/gomock"
)

// MockScriptApi is a mock of ScriptApi interface.
type MockScriptApi struct {
	ctrl     *gomock.Controller
	recorder *MockScriptApiMockRecorder
}

// MockScriptApiMockRecorder is the mock recorder for MockScriptApi.
type MockScriptApiMockRecorder struct {
	mock *MockScriptApi
}

// NewMockScriptApi creates a new mock instance.
func NewMockScriptApi(ctrl *gomock.Controller) *MockScriptApi {
	mock := &MockScriptApi{ctrl: ctrl}
	mock.recorder = &MockScriptApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScriptApi) EXPECT() *MockScriptApiMockRecorder {
	return m.recorder
}

// AnalyzeLockingScript mocks base method.
func (m *MockScriptApi) AnalyzeLockingScript(script *types.Script) (types.HashType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnalyzeLockingScript", script)
	ret0, _ := ret[0].(types.HashType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnalyzeLockingScript indicates an expected call of AnalyzeLockingScript.
func (mr *MockScriptApiMockRecorder) AnalyzeLockingScript(script any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnalyzeLockingScript", reflect.TypeOf((*MockScriptApi)(nil).AnalyzeLockingScript), script)
}

// CreateFromAsm mocks base method.
func (m *MockScriptApi) CreateFromAsm(asm string) (*types.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromAsm", asm)
	ret0, _ := ret[0].(*types.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromAsm indicates an expected call of CreateFromAsm.
func (mr *MockScriptApiMockRecorder) CreateFromAsm(asm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromAsm", reflect.TypeOf((*MockScriptApi)(nil).CreateFromAsm), asm)
}

// CreateFromAsmStrings mocks base method.
func (m *MockScriptApi) CreateFromAsmStrings(asmStrings []string) (*types.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFromAsmStrings", asmStrings)
	ret0, _ := ret[0].(*types.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFromAsmStrings indicates an expected call of CreateFromAsmStrings.
func (mr *MockScriptApiMockRecorder) CreateFromAsmStrings(asmStrings any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFromAsmStrings", reflect.TypeOf((*MockScriptApi)(nil).CreateFromAsmStrings), asmStrings)
}

// CreateMultisig mocks base method.
func (m *MockScriptApi) CreateMultisig(pubkeys []*types.Pubkey, requireSigNum uint32) (*types.Script, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMultisig", pubkeys, requireSigNum)
	ret0, _ := ret[0].(*types.Script)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMultisig indicates an expected call of CreateMultisig.
func (mr *MockScriptApiMockRecorder) CreateMultisig(pubkeys, requireSigNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultisig", reflect.TypeOf((*MockScriptApi)(nil).CreateMultisig), pubkeys, requireSigNum)
}

// IsCheckHashType mocks base method.
func (m *MockScriptApi) IsCheckHashType(hashType types.HashType, script *types.Script) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCheckHashType", hashType, script)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsCheckHashType indicates an expected call of IsCheckHashType.
func (mr *MockScriptApiMockRecorder) IsCheckHashType(hashType, script any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCheckHashType", reflect.TypeOf((*MockScriptApi)(nil).IsCheckHashType), hashType, script)
}

// Parse mocks base method.
func (m *MockScriptApi) Parse(script *types.Script) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Parse", script)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Parse indicates an expected call of Parse.
func (mr *MockScriptApiMockRecorder) Parse(script any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Parse", reflect.TypeOf((*MockScriptApi)(nil).Parse), script)
}

// ParseMultisig mocks base method.
func (m *MockScriptApi) ParseMultisig(script *types.Script) ([]*types.Pubkey, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseMultisig", script)
	ret0, _ := ret[0].([]*types.Pubkey)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ParseMultisig indicates an expected call of ParseMultisig.
func (mr *MockScriptApiMockRecorder) ParseMultisig(script any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseMultisig", reflect.TypeOf((*MockScriptApi)(nil).ParseMultisig), script)
}
