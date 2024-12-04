// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ids "github.com/MetalBlockchain/metalgo/ids"

	interfaces "github.com/MetalBlockchain/coreth/interfaces"

	mock "github.com/stretchr/testify/mock"

	types "github.com/MetalBlockchain/coreth/core/types"
)

// EthClient is an autogenerated mock type for the EthClient type
type EthClient struct {
	mock.Mock
}

// AcceptedCallContract provides a mock function with given fields: _a0, _a1
func (_m *EthClient) AcceptedCallContract(_a0 context.Context, _a1 interfaces.CallMsg) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.CallMsg) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.CallMsg) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptedCodeAt provides a mock function with given fields: _a0, _a1
func (_m *EthClient) AcceptedCodeAt(_a0 context.Context, _a1 common.Address) ([]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) []byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AcceptedNonceAt provides a mock function with given fields: _a0, _a1
func (_m *EthClient) AcceptedNonceAt(_a0 context.Context, _a1 common.Address) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AssetBalanceAt provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *EthClient) AssetBalanceAt(_a0 context.Context, _a1 common.Address, _a2 ids.ID, _a3 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, ids.ID, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, ids.ID, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BalanceAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) BalanceAt(_a0 context.Context, _a1 common.Address, _a2 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByHash provides a mock function with given fields: _a0, _a1
func (_m *EthClient) BlockByHash(_a0 context.Context, _a1 common.Hash) (*types.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockByNumber provides a mock function with given fields: _a0, _a1
func (_m *EthClient) BlockByNumber(_a0 context.Context, _a1 *big.Int) (*types.Block, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Block); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BlockNumber provides a mock function with given fields: _a0
func (_m *EthClient) BlockNumber(_a0 context.Context) (uint64, error) {
	ret := _m.Called(_a0)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CallContract provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) CallContract(_a0 context.Context, _a1 interfaces.CallMsg, _a2 *big.Int) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.CallMsg, *big.Int) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.CallMsg, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *EthClient) Close() {
	_m.Called()
}

// CodeAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) CodeAt(_a0 context.Context, _a1 common.Address, _a2 *big.Int) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: _a0, _a1
func (_m *EthClient) EstimateGas(_a0 context.Context, _a1 interfaces.CallMsg) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.CallMsg) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.CallMsg) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: _a0, _a1
func (_m *EthClient) FilterLogs(_a0 context.Context, _a1 interfaces.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.FilterQuery) []types.Log); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.FilterQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: _a0, _a1
func (_m *EthClient) HeaderByNumber(_a0 context.Context, _a1 *big.Int) (*types.Header, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonceAt provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) NonceAt(_a0 context.Context, _a1 common.Address, _a2 *big.Int) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: _a0, _a1
func (_m *EthClient) SendTransaction(_a0 context.Context, _a1 *types.Transaction) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeFilterLogs provides a mock function with given fields: _a0, _a1, _a2
func (_m *EthClient) SubscribeFilterLogs(_a0 context.Context, _a1 interfaces.FilterQuery, _a2 chan<- types.Log) (interfaces.Subscription, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 interfaces.Subscription
	if rf, ok := ret.Get(0).(func(context.Context, interfaces.FilterQuery, chan<- types.Log) interfaces.Subscription); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.Subscription)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interfaces.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: _a0
func (_m *EthClient) SuggestGasPrice(_a0 context.Context) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasTipCap provides a mock function with given fields: _a0
func (_m *EthClient) SuggestGasTipCap(_a0 context.Context) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: _a0, _a1
func (_m *EthClient) TransactionReceipt(_a0 context.Context, _a1 common.Hash) (*types.Receipt, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
