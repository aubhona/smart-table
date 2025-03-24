// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// PostCustomerV1OrderCreateResponseObject is an autogenerated mock type for the PostCustomerV1OrderCreateResponseObject type
type PostCustomerV1OrderCreateResponseObject struct {
	mock.Mock
}

type PostCustomerV1OrderCreateResponseObject_Expecter struct {
	mock *mock.Mock
}

func (_m *PostCustomerV1OrderCreateResponseObject) EXPECT() *PostCustomerV1OrderCreateResponseObject_Expecter {
	return &PostCustomerV1OrderCreateResponseObject_Expecter{mock: &_m.Mock}
}

// VisitPostCustomerV1OrderCreateResponse provides a mock function with given fields: w
func (_m *PostCustomerV1OrderCreateResponseObject) VisitPostCustomerV1OrderCreateResponse(w http.ResponseWriter) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for VisitPostCustomerV1OrderCreateResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisitPostCustomerV1OrderCreateResponse'
type PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call struct {
	*mock.Call
}

// VisitPostCustomerV1OrderCreateResponse is a helper method to define mock.On call
//   - w http.ResponseWriter
func (_e *PostCustomerV1OrderCreateResponseObject_Expecter) VisitPostCustomerV1OrderCreateResponse(w interface{}) *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call {
	return &PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call{Call: _e.mock.On("VisitPostCustomerV1OrderCreateResponse", w)}
}

func (_c *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call) Run(run func(w http.ResponseWriter)) *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter))
	})
	return _c
}

func (_c *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call) Return(_a0 error) *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call) RunAndReturn(run func(http.ResponseWriter) error) *PostCustomerV1OrderCreateResponseObject_VisitPostCustomerV1OrderCreateResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewPostCustomerV1OrderCreateResponseObject creates a new instance of PostCustomerV1OrderCreateResponseObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostCustomerV1OrderCreateResponseObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostCustomerV1OrderCreateResponseObject {
	mock := &PostCustomerV1OrderCreateResponseObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
