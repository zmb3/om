// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ConfigureOpsmanService struct {
	EnableRBACStub        func(api.RBACSettings) error
	enableRBACMutex       sync.RWMutex
	enableRBACArgsForCall []struct {
		arg1 api.RBACSettings
	}
	enableRBACReturns struct {
		result1 error
	}
	enableRBACReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateBannerStub        func(api.BannerSettings) error
	updateBannerMutex       sync.RWMutex
	updateBannerArgsForCall []struct {
		arg1 api.BannerSettings
	}
	updateBannerReturns struct {
		result1 error
	}
	updateBannerReturnsOnCall map[int]struct {
		result1 error
	}
	UpdatePivnetTokenStub        func(api.PivnetSettings) error
	updatePivnetTokenMutex       sync.RWMutex
	updatePivnetTokenArgsForCall []struct {
		arg1 api.PivnetSettings
	}
	updatePivnetTokenReturns struct {
		result1 error
	}
	updatePivnetTokenReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSSLCertificateStub        func(api.SSLCertificateInput) error
	updateSSLCertificateMutex       sync.RWMutex
	updateSSLCertificateArgsForCall []struct {
		arg1 api.SSLCertificateInput
	}
	updateSSLCertificateReturns struct {
		result1 error
	}
	updateSSLCertificateReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateSyslogSettingsStub        func(api.SyslogSettings) error
	updateSyslogSettingsMutex       sync.RWMutex
	updateSyslogSettingsArgsForCall []struct {
		arg1 api.SyslogSettings
	}
	updateSyslogSettingsReturns struct {
		result1 error
	}
	updateSyslogSettingsReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ConfigureOpsmanService) EnableRBAC(arg1 api.RBACSettings) error {
	fake.enableRBACMutex.Lock()
	ret, specificReturn := fake.enableRBACReturnsOnCall[len(fake.enableRBACArgsForCall)]
	fake.enableRBACArgsForCall = append(fake.enableRBACArgsForCall, struct {
		arg1 api.RBACSettings
	}{arg1})
	fake.recordInvocation("EnableRBAC", []interface{}{arg1})
	fake.enableRBACMutex.Unlock()
	if fake.EnableRBACStub != nil {
		return fake.EnableRBACStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.enableRBACReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) EnableRBACCallCount() int {
	fake.enableRBACMutex.RLock()
	defer fake.enableRBACMutex.RUnlock()
	return len(fake.enableRBACArgsForCall)
}

func (fake *ConfigureOpsmanService) EnableRBACCalls(stub func(api.RBACSettings) error) {
	fake.enableRBACMutex.Lock()
	defer fake.enableRBACMutex.Unlock()
	fake.EnableRBACStub = stub
}

func (fake *ConfigureOpsmanService) EnableRBACArgsForCall(i int) api.RBACSettings {
	fake.enableRBACMutex.RLock()
	defer fake.enableRBACMutex.RUnlock()
	argsForCall := fake.enableRBACArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) EnableRBACReturns(result1 error) {
	fake.enableRBACMutex.Lock()
	defer fake.enableRBACMutex.Unlock()
	fake.EnableRBACStub = nil
	fake.enableRBACReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) EnableRBACReturnsOnCall(i int, result1 error) {
	fake.enableRBACMutex.Lock()
	defer fake.enableRBACMutex.Unlock()
	fake.EnableRBACStub = nil
	if fake.enableRBACReturnsOnCall == nil {
		fake.enableRBACReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.enableRBACReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateBanner(arg1 api.BannerSettings) error {
	fake.updateBannerMutex.Lock()
	ret, specificReturn := fake.updateBannerReturnsOnCall[len(fake.updateBannerArgsForCall)]
	fake.updateBannerArgsForCall = append(fake.updateBannerArgsForCall, struct {
		arg1 api.BannerSettings
	}{arg1})
	fake.recordInvocation("UpdateBanner", []interface{}{arg1})
	fake.updateBannerMutex.Unlock()
	if fake.UpdateBannerStub != nil {
		return fake.UpdateBannerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateBannerReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdateBannerCallCount() int {
	fake.updateBannerMutex.RLock()
	defer fake.updateBannerMutex.RUnlock()
	return len(fake.updateBannerArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdateBannerCalls(stub func(api.BannerSettings) error) {
	fake.updateBannerMutex.Lock()
	defer fake.updateBannerMutex.Unlock()
	fake.UpdateBannerStub = stub
}

func (fake *ConfigureOpsmanService) UpdateBannerArgsForCall(i int) api.BannerSettings {
	fake.updateBannerMutex.RLock()
	defer fake.updateBannerMutex.RUnlock()
	argsForCall := fake.updateBannerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdateBannerReturns(result1 error) {
	fake.updateBannerMutex.Lock()
	defer fake.updateBannerMutex.Unlock()
	fake.UpdateBannerStub = nil
	fake.updateBannerReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateBannerReturnsOnCall(i int, result1 error) {
	fake.updateBannerMutex.Lock()
	defer fake.updateBannerMutex.Unlock()
	fake.UpdateBannerStub = nil
	if fake.updateBannerReturnsOnCall == nil {
		fake.updateBannerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateBannerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdatePivnetToken(arg1 api.PivnetSettings) error {
	fake.updatePivnetTokenMutex.Lock()
	ret, specificReturn := fake.updatePivnetTokenReturnsOnCall[len(fake.updatePivnetTokenArgsForCall)]
	fake.updatePivnetTokenArgsForCall = append(fake.updatePivnetTokenArgsForCall, struct {
		arg1 api.PivnetSettings
	}{arg1})
	fake.recordInvocation("UpdatePivnetToken", []interface{}{arg1})
	fake.updatePivnetTokenMutex.Unlock()
	if fake.UpdatePivnetTokenStub != nil {
		return fake.UpdatePivnetTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updatePivnetTokenReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenCallCount() int {
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	return len(fake.updatePivnetTokenArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenCalls(stub func(api.PivnetSettings) error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = stub
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenArgsForCall(i int) api.PivnetSettings {
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	argsForCall := fake.updatePivnetTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenReturns(result1 error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = nil
	fake.updatePivnetTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdatePivnetTokenReturnsOnCall(i int, result1 error) {
	fake.updatePivnetTokenMutex.Lock()
	defer fake.updatePivnetTokenMutex.Unlock()
	fake.UpdatePivnetTokenStub = nil
	if fake.updatePivnetTokenReturnsOnCall == nil {
		fake.updatePivnetTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updatePivnetTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificate(arg1 api.SSLCertificateInput) error {
	fake.updateSSLCertificateMutex.Lock()
	ret, specificReturn := fake.updateSSLCertificateReturnsOnCall[len(fake.updateSSLCertificateArgsForCall)]
	fake.updateSSLCertificateArgsForCall = append(fake.updateSSLCertificateArgsForCall, struct {
		arg1 api.SSLCertificateInput
	}{arg1})
	fake.recordInvocation("UpdateSSLCertificate", []interface{}{arg1})
	fake.updateSSLCertificateMutex.Unlock()
	if fake.UpdateSSLCertificateStub != nil {
		return fake.UpdateSSLCertificateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSSLCertificateReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateCallCount() int {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	return len(fake.updateSSLCertificateArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateCalls(stub func(api.SSLCertificateInput) error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = stub
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateArgsForCall(i int) api.SSLCertificateInput {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	argsForCall := fake.updateSSLCertificateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateReturns(result1 error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = nil
	fake.updateSSLCertificateReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSSLCertificateReturnsOnCall(i int, result1 error) {
	fake.updateSSLCertificateMutex.Lock()
	defer fake.updateSSLCertificateMutex.Unlock()
	fake.UpdateSSLCertificateStub = nil
	if fake.updateSSLCertificateReturnsOnCall == nil {
		fake.updateSSLCertificateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSSLCertificateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettings(arg1 api.SyslogSettings) error {
	fake.updateSyslogSettingsMutex.Lock()
	ret, specificReturn := fake.updateSyslogSettingsReturnsOnCall[len(fake.updateSyslogSettingsArgsForCall)]
	fake.updateSyslogSettingsArgsForCall = append(fake.updateSyslogSettingsArgsForCall, struct {
		arg1 api.SyslogSettings
	}{arg1})
	fake.recordInvocation("UpdateSyslogSettings", []interface{}{arg1})
	fake.updateSyslogSettingsMutex.Unlock()
	if fake.UpdateSyslogSettingsStub != nil {
		return fake.UpdateSyslogSettingsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateSyslogSettingsReturns
	return fakeReturns.result1
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettingsCallCount() int {
	fake.updateSyslogSettingsMutex.RLock()
	defer fake.updateSyslogSettingsMutex.RUnlock()
	return len(fake.updateSyslogSettingsArgsForCall)
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettingsCalls(stub func(api.SyslogSettings) error) {
	fake.updateSyslogSettingsMutex.Lock()
	defer fake.updateSyslogSettingsMutex.Unlock()
	fake.UpdateSyslogSettingsStub = stub
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettingsArgsForCall(i int) api.SyslogSettings {
	fake.updateSyslogSettingsMutex.RLock()
	defer fake.updateSyslogSettingsMutex.RUnlock()
	argsForCall := fake.updateSyslogSettingsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettingsReturns(result1 error) {
	fake.updateSyslogSettingsMutex.Lock()
	defer fake.updateSyslogSettingsMutex.Unlock()
	fake.UpdateSyslogSettingsStub = nil
	fake.updateSyslogSettingsReturns = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) UpdateSyslogSettingsReturnsOnCall(i int, result1 error) {
	fake.updateSyslogSettingsMutex.Lock()
	defer fake.updateSyslogSettingsMutex.Unlock()
	fake.UpdateSyslogSettingsStub = nil
	if fake.updateSyslogSettingsReturnsOnCall == nil {
		fake.updateSyslogSettingsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSyslogSettingsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ConfigureOpsmanService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enableRBACMutex.RLock()
	defer fake.enableRBACMutex.RUnlock()
	fake.updateBannerMutex.RLock()
	defer fake.updateBannerMutex.RUnlock()
	fake.updatePivnetTokenMutex.RLock()
	defer fake.updatePivnetTokenMutex.RUnlock()
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	fake.updateSyslogSettingsMutex.RLock()
	defer fake.updateSyslogSettingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ConfigureOpsmanService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
