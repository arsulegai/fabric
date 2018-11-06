// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import common "github.com/hyperledger/fabric/protos/common"
import mock "github.com/stretchr/testify/mock"

// ChainPuller is an autogenerated mock type for the ChainPuller type
type ChainPuller struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ChainPuller) Close() {
	_m.Called()
}

// PullBlock provides a mock function with given fields: seq
func (_m *ChainPuller) PullBlock(seq uint64) *common.Block {
	ret := _m.Called(seq)

	var r0 *common.Block
	if rf, ok := ret.Get(0).(func(uint64) *common.Block); ok {
		r0 = rf(seq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*common.Block)
		}
	}

	return r0
}