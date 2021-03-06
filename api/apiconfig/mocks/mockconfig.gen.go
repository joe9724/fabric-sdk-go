// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hyperledger/fabric-sdk-go/api/apiconfig (interfaces: Config)

package mock_apiconfig

import (
	x509 "crypto/x509"
	gomock "github.com/golang/mock/gomock"
	apiconfig "github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	factory "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/bccsp/factory"
	time "time"
)

// MockConfig is a mock of Config interface
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return _m.recorder
}

// CAClientCertFile mocks base method
func (_m *MockConfig) CAClientCertFile(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "CAClientCertFile", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CAClientCertFile indicates an expected call of CAClientCertFile
func (_mr *MockConfigMockRecorder) CAClientCertFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CAClientCertFile", arg0)
}

// CAClientKeyFile mocks base method
func (_m *MockConfig) CAClientKeyFile(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "CAClientKeyFile", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CAClientKeyFile indicates an expected call of CAClientKeyFile
func (_mr *MockConfigMockRecorder) CAClientKeyFile(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CAClientKeyFile", arg0)
}

// CAConfig mocks base method
func (_m *MockConfig) CAConfig(_param0 string) (*apiconfig.CAConfig, error) {
	ret := _m.ctrl.Call(_m, "CAConfig", _param0)
	ret0, _ := ret[0].(*apiconfig.CAConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CAConfig indicates an expected call of CAConfig
func (_mr *MockConfigMockRecorder) CAConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CAConfig", arg0)
}

// CAKeyStorePath mocks base method
func (_m *MockConfig) CAKeyStorePath() string {
	ret := _m.ctrl.Call(_m, "CAKeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CAKeyStorePath indicates an expected call of CAKeyStorePath
func (_mr *MockConfigMockRecorder) CAKeyStorePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CAKeyStorePath")
}

// CAServerCertFiles mocks base method
func (_m *MockConfig) CAServerCertFiles(_param0 string) ([]string, error) {
	ret := _m.ctrl.Call(_m, "CAServerCertFiles", _param0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CAServerCertFiles indicates an expected call of CAServerCertFiles
func (_mr *MockConfigMockRecorder) CAServerCertFiles(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CAServerCertFiles", arg0)
}

// CSPConfig mocks base method
func (_m *MockConfig) CSPConfig() *factory.FactoryOpts {
	ret := _m.ctrl.Call(_m, "CSPConfig")
	ret0, _ := ret[0].(*factory.FactoryOpts)
	return ret0
}

// CSPConfig indicates an expected call of CSPConfig
func (_mr *MockConfigMockRecorder) CSPConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CSPConfig")
}

// Client mocks base method
func (_m *MockConfig) Client() (*apiconfig.ClientConfig, error) {
	ret := _m.ctrl.Call(_m, "Client")
	ret0, _ := ret[0].(*apiconfig.ClientConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Client indicates an expected call of Client
func (_mr *MockConfigMockRecorder) Client() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Client")
}

// CryptoConfigPath mocks base method
func (_m *MockConfig) CryptoConfigPath() string {
	ret := _m.ctrl.Call(_m, "CryptoConfigPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CryptoConfigPath indicates an expected call of CryptoConfigPath
func (_mr *MockConfigMockRecorder) CryptoConfigPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CryptoConfigPath")
}

// Ephemeral mocks base method
func (_m *MockConfig) Ephemeral() bool {
	ret := _m.ctrl.Call(_m, "Ephemeral")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ephemeral indicates an expected call of Ephemeral
func (_mr *MockConfigMockRecorder) Ephemeral() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ephemeral")
}

// IsSecurityEnabled mocks base method
func (_m *MockConfig) IsSecurityEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsSecurityEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsSecurityEnabled indicates an expected call of IsSecurityEnabled
func (_mr *MockConfigMockRecorder) IsSecurityEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsSecurityEnabled")
}

// IsTLSEnabled mocks base method
func (_m *MockConfig) IsTLSEnabled() bool {
	ret := _m.ctrl.Call(_m, "IsTLSEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTLSEnabled indicates an expected call of IsTLSEnabled
func (_mr *MockConfigMockRecorder) IsTLSEnabled() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsTLSEnabled")
}

// KeyStorePath mocks base method
func (_m *MockConfig) KeyStorePath() string {
	ret := _m.ctrl.Call(_m, "KeyStorePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// KeyStorePath indicates an expected call of KeyStorePath
func (_mr *MockConfigMockRecorder) KeyStorePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KeyStorePath")
}

// MspID mocks base method
func (_m *MockConfig) MspID(_param0 string) (string, error) {
	ret := _m.ctrl.Call(_m, "MspID", _param0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MspID indicates an expected call of MspID
func (_mr *MockConfigMockRecorder) MspID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MspID", arg0)
}

// NetworkConfig mocks base method
func (_m *MockConfig) NetworkConfig() (*apiconfig.NetworkConfig, error) {
	ret := _m.ctrl.Call(_m, "NetworkConfig")
	ret0, _ := ret[0].(*apiconfig.NetworkConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkConfig indicates an expected call of NetworkConfig
func (_mr *MockConfigMockRecorder) NetworkConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NetworkConfig")
}

// OrdererConfig mocks base method
func (_m *MockConfig) OrdererConfig(_param0 string) (*apiconfig.OrdererConfig, error) {
	ret := _m.ctrl.Call(_m, "OrdererConfig", _param0)
	ret0, _ := ret[0].(*apiconfig.OrdererConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrdererConfig indicates an expected call of OrdererConfig
func (_mr *MockConfigMockRecorder) OrdererConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OrdererConfig", arg0)
}

// OrderersConfig mocks base method
func (_m *MockConfig) OrderersConfig() ([]apiconfig.OrdererConfig, error) {
	ret := _m.ctrl.Call(_m, "OrderersConfig")
	ret0, _ := ret[0].([]apiconfig.OrdererConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderersConfig indicates an expected call of OrderersConfig
func (_mr *MockConfigMockRecorder) OrderersConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OrderersConfig")
}

// PeerConfig mocks base method
func (_m *MockConfig) PeerConfig(_param0 string, _param1 string) (*apiconfig.PeerConfig, error) {
	ret := _m.ctrl.Call(_m, "PeerConfig", _param0, _param1)
	ret0, _ := ret[0].(*apiconfig.PeerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerConfig indicates an expected call of PeerConfig
func (_mr *MockConfigMockRecorder) PeerConfig(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PeerConfig", arg0, arg1)
}

// PeersConfig mocks base method
func (_m *MockConfig) PeersConfig(_param0 string) ([]apiconfig.PeerConfig, error) {
	ret := _m.ctrl.Call(_m, "PeersConfig", _param0)
	ret0, _ := ret[0].([]apiconfig.PeerConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeersConfig indicates an expected call of PeersConfig
func (_mr *MockConfigMockRecorder) PeersConfig(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PeersConfig", arg0)
}

// RandomOrdererConfig mocks base method
func (_m *MockConfig) RandomOrdererConfig() (*apiconfig.OrdererConfig, error) {
	ret := _m.ctrl.Call(_m, "RandomOrdererConfig")
	ret0, _ := ret[0].(*apiconfig.OrdererConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RandomOrdererConfig indicates an expected call of RandomOrdererConfig
func (_mr *MockConfigMockRecorder) RandomOrdererConfig() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RandomOrdererConfig")
}

// SecurityAlgorithm mocks base method
func (_m *MockConfig) SecurityAlgorithm() string {
	ret := _m.ctrl.Call(_m, "SecurityAlgorithm")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityAlgorithm indicates an expected call of SecurityAlgorithm
func (_mr *MockConfigMockRecorder) SecurityAlgorithm() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecurityAlgorithm")
}

// SecurityLevel mocks base method
func (_m *MockConfig) SecurityLevel() int {
	ret := _m.ctrl.Call(_m, "SecurityLevel")
	ret0, _ := ret[0].(int)
	return ret0
}

// SecurityLevel indicates an expected call of SecurityLevel
func (_mr *MockConfigMockRecorder) SecurityLevel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecurityLevel")
}

// SecurityProvider mocks base method
func (_m *MockConfig) SecurityProvider() string {
	ret := _m.ctrl.Call(_m, "SecurityProvider")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProvider indicates an expected call of SecurityProvider
func (_mr *MockConfigMockRecorder) SecurityProvider() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecurityProvider")
}

// SecurityProviderLabel mocks base method
func (_m *MockConfig) SecurityProviderLabel() string {
	ret := _m.ctrl.Call(_m, "SecurityProviderLabel")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderLabel indicates an expected call of SecurityProviderLabel
func (_mr *MockConfigMockRecorder) SecurityProviderLabel() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecurityProviderLabel")
}

// SecurityProviderPin mocks base method
func (_m *MockConfig) SecurityProviderPin() string {
	ret := _m.ctrl.Call(_m, "SecurityProviderPin")
	ret0, _ := ret[0].(string)
	return ret0
}

// SecurityProviderPin indicates an expected call of SecurityProviderPin
func (_mr *MockConfigMockRecorder) SecurityProviderPin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SecurityProviderPin")
}

// SetTLSCACertPool mocks base method
func (_m *MockConfig) SetTLSCACertPool(_param0 *x509.CertPool) {
	_m.ctrl.Call(_m, "SetTLSCACertPool", _param0)
}

// SetTLSCACertPool indicates an expected call of SetTLSCACertPool
func (_mr *MockConfigMockRecorder) SetTLSCACertPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTLSCACertPool", arg0)
}

// SoftVerify mocks base method
func (_m *MockConfig) SoftVerify() bool {
	ret := _m.ctrl.Call(_m, "SoftVerify")
	ret0, _ := ret[0].(bool)
	return ret0
}

// SoftVerify indicates an expected call of SoftVerify
func (_mr *MockConfigMockRecorder) SoftVerify() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SoftVerify")
}

// TLSCACertPool mocks base method
func (_m *MockConfig) TLSCACertPool(_param0 string) (*x509.CertPool, error) {
	ret := _m.ctrl.Call(_m, "TLSCACertPool", _param0)
	ret0, _ := ret[0].(*x509.CertPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TLSCACertPool indicates an expected call of TLSCACertPool
func (_mr *MockConfigMockRecorder) TLSCACertPool(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TLSCACertPool", arg0)
}

// TimeoutOrDefault mocks base method
func (_m *MockConfig) TimeoutOrDefault(_param0 apiconfig.ConnectionType) time.Duration {
	ret := _m.ctrl.Call(_m, "TimeoutOrDefault", _param0)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeoutOrDefault indicates an expected call of TimeoutOrDefault
func (_mr *MockConfigMockRecorder) TimeoutOrDefault(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TimeoutOrDefault", arg0)
}
