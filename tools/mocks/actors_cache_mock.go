// Code generated by mockery v2.53.4. DO NOT EDIT.

package mocks

import (
	address "github.com/filecoin-project/go-address"
	backoff "github.com/zondax/golem/pkg/zhttpclient/backoff"

	common "github.com/zondax/fil-parser/actors/cache/impl/common"

	context "context"

	fil_parsertypes "github.com/zondax/fil-parser/types"

	logger "github.com/zondax/golem/pkg/logger"

	metrics "github.com/zondax/fil-parser/actors/cache/metrics"

	mock "github.com/stretchr/testify/mock"

	types "github.com/filecoin-project/lotus/chain/types"
)

// IActorsCache is an autogenerated mock type for the IActorsCache type
type IActorsCache struct {
	mock.Mock
}

// BackFill provides a mock function with no fields
func (_m *IActorsCache) BackFill() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BackFill")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearBadAddressCache provides a mock function with no fields
func (_m *IActorsCache) ClearBadAddressCache() {
	_m.Called()
}

// GetActorCode provides a mock function with given fields: add, key, onChainOnly
func (_m *IActorsCache) GetActorCode(add address.Address, key types.TipSetKey, onChainOnly bool) (string, error) {
	ret := _m.Called(add, key, onChainOnly)

	if len(ret) == 0 {
		panic("no return value specified for GetActorCode")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(address.Address, types.TipSetKey, bool) (string, error)); ok {
		return rf(add, key, onChainOnly)
	}
	if rf, ok := ret.Get(0).(func(address.Address, types.TipSetKey, bool) string); ok {
		r0 = rf(add, key, onChainOnly)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(address.Address, types.TipSetKey, bool) error); ok {
		r1 = rf(add, key, onChainOnly)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEVMSelectorSig provides a mock function with given fields: ctx, selectorHash
func (_m *IActorsCache) GetEVMSelectorSig(ctx context.Context, selectorHash string) (string, error) {
	ret := _m.Called(ctx, selectorHash)

	if len(ret) == 0 {
		panic("no return value specified for GetEVMSelectorSig")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, selectorHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, selectorHash)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, selectorHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRobustAddress provides a mock function with given fields: add
func (_m *IActorsCache) GetRobustAddress(add address.Address) (string, error) {
	ret := _m.Called(add)

	if len(ret) == 0 {
		panic("no return value specified for GetRobustAddress")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(address.Address) (string, error)); ok {
		return rf(add)
	}
	if rf, ok := ret.Get(0).(func(address.Address) string); ok {
		r0 = rf(add)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(address.Address) error); ok {
		r1 = rf(add)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetShortAddress provides a mock function with given fields: add
func (_m *IActorsCache) GetShortAddress(add address.Address) (string, error) {
	ret := _m.Called(add)

	if len(ret) == 0 {
		panic("no return value specified for GetShortAddress")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(address.Address) (string, error)); ok {
		return rf(add)
	}
	if rf, ok := ret.Get(0).(func(address.Address) string); ok {
		r0 = rf(add)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(address.Address) error); ok {
		r1 = rf(add)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImplementationType provides a mock function with no fields
func (_m *IActorsCache) ImplementationType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ImplementationType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsGenesisActor provides a mock function with given fields: addr
func (_m *IActorsCache) IsGenesisActor(addr string) bool {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for IsGenesisActor")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSystemActor provides a mock function with given fields: addr
func (_m *IActorsCache) IsSystemActor(addr string) bool {
	ret := _m.Called(addr)

	if len(ret) == 0 {
		panic("no return value specified for IsSystemActor")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewImpl provides a mock function with given fields: source, _a1, _a2, _a3
func (_m *IActorsCache) NewImpl(source common.DataSource, _a1 *logger.Logger, _a2 *metrics.ActorsCacheMetricsClient, _a3 *backoff.BackOff) error {
	ret := _m.Called(source, _a1, _a2, _a3)

	if len(ret) == 0 {
		panic("no return value specified for NewImpl")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(common.DataSource, *logger.Logger, *metrics.ActorsCacheMetricsClient, *backoff.BackOff) error); ok {
		r0 = rf(source, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreAddressInfo provides a mock function with given fields: info
func (_m *IActorsCache) StoreAddressInfo(info fil_parsertypes.AddressInfo) {
	_m.Called(info)
}

// StoreEVMSelectorSig provides a mock function with given fields: ctx, selectorHash, selectorSig
func (_m *IActorsCache) StoreEVMSelectorSig(ctx context.Context, selectorHash string, selectorSig string) error {
	ret := _m.Called(ctx, selectorHash, selectorSig)

	if len(ret) == 0 {
		panic("no return value specified for StoreEVMSelectorSig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, selectorHash, selectorSig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIActorsCache creates a new instance of IActorsCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIActorsCache(t interface {
	mock.TestingT
	Cleanup(func())
}) *IActorsCache {
	mock := &IActorsCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
