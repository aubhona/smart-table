// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// PostAdminV1UserSignInResponseObject is an autogenerated mock type for the PostAdminV1UserSignInResponseObject type
type PostAdminV1UserSignInResponseObject struct {
	mock.Mock
}

type PostAdminV1UserSignInResponseObject_Expecter struct {
	mock *mock.Mock
}

func (_m *PostAdminV1UserSignInResponseObject) EXPECT() *PostAdminV1UserSignInResponseObject_Expecter {
	return &PostAdminV1UserSignInResponseObject_Expecter{mock: &_m.Mock}
}

// VisitPostAdminV1UserSignInResponse provides a mock function with given fields: w
func (_m *PostAdminV1UserSignInResponseObject) VisitPostAdminV1UserSignInResponse(w http.ResponseWriter) error {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for VisitPostAdminV1UserSignInResponse")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter) error); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VisitPostAdminV1UserSignInResponse'
type PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call struct {
	*mock.Call
}

// VisitPostAdminV1UserSignInResponse is a helper method to define mock.On call
//   - w http.ResponseWriter
func (_e *PostAdminV1UserSignInResponseObject_Expecter) VisitPostAdminV1UserSignInResponse(w interface{}) *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call {
	return &PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call{Call: _e.mock.On("VisitPostAdminV1UserSignInResponse", w)}
}

func (_c *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call) Run(run func(w http.ResponseWriter)) *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter))
	})
	return _c
}

func (_c *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call) Return(_a0 error) *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call) RunAndReturn(run func(http.ResponseWriter) error) *PostAdminV1UserSignInResponseObject_VisitPostAdminV1UserSignInResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewPostAdminV1UserSignInResponseObject creates a new instance of PostAdminV1UserSignInResponseObject. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPostAdminV1UserSignInResponseObject(t interface {
	mock.TestingT
	Cleanup(func())
}) *PostAdminV1UserSignInResponseObject {
	mock := &PostAdminV1UserSignInResponseObject{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
