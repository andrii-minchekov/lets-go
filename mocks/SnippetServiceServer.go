// Code generated by mockery v2.6.0. DO NOT EDIT.

package mocks

import (
	context "context"

	generated "github.com/andrii-minchekov/lets-go/app/impl/grpc"
	mock "github.com/stretchr/testify/mock"
)

// SnippetServiceServer is an autogenerated mock type for the SnippetServiceServer type
type SnippetServiceServer struct {
	mock.Mock
}

// CreateSnippet provides a mock function with given fields: _a0, _a1
func (_m *SnippetServiceServer) CreateSnippet(_a0 context.Context, _a1 *generated.CreateSnippetRequest) (*generated.CreateSnippetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *generated.CreateSnippetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *generated.CreateSnippetRequest) *generated.CreateSnippetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*generated.CreateSnippetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *generated.CreateSnippetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
