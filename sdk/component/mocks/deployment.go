// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	proto "github.com/golang/protobuf/proto"
	mock "github.com/stretchr/testify/mock"
)

// Deployment is an autogenerated mock type for the Deployment type
type Deployment struct {
	mock.Mock
}

// Proto provides a mock function with given fields:
func (_m *Deployment) Proto() proto.Message {
	ret := _m.Called()

	var r0 proto.Message
	if rf, ok := ret.Get(0).(func() proto.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(proto.Message)
		}
	}

	return r0
}

// String provides a mock function with given fields:
func (_m *Deployment) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
