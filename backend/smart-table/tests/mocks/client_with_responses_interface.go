// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	viewsLUcustomer "github.com/smart-table/src/views/codegen/customer"
)

// ClientWithResponsesInterface is an autogenerated mock type for the ClientWithResponsesInterface type
type ClientWithResponsesInterface struct {
	mock.Mock
}

type ClientWithResponsesInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ClientWithResponsesInterface) EXPECT() *ClientWithResponsesInterface_Expecter {
	return &ClientWithResponsesInterface_Expecter{mock: &_m.Mock}
}

// PostCustomerV1OrderCreateWithBodyWithResponse provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *ClientWithResponsesInterface) PostCustomerV1OrderCreateWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PostCustomerV1OrderCreateWithBodyWithResponse")
	}

	var r0 *viewsLUcustomer.PostCustomerV1OrderCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) *viewsLUcustomer.PostCustomerV1OrderCreateResponse); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viewsLUcustomer.PostCustomerV1OrderCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCustomerV1OrderCreateWithBodyWithResponse'
type ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call struct {
	*mock.Call
}

// PostCustomerV1OrderCreateWithBodyWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - contentType string
//   - body io.Reader
//   - reqEditors ...viewsLUcustomer.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) PostCustomerV1OrderCreateWithBodyWithResponse(ctx interface{}, contentType interface{}, body interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call {
	return &ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call{Call: _e.mock.On("PostCustomerV1OrderCreateWithBodyWithResponse",
		append([]interface{}{ctx, contentType, body}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call) Run(run func(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn)) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]viewsLUcustomer.RequestEditorFn, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(viewsLUcustomer.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call) Return(_a0 *viewsLUcustomer.PostCustomerV1OrderCreateResponse, _a1 error) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call) RunAndReturn(run func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error)) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithBodyWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// PostCustomerV1OrderCreateWithResponse provides a mock function with given fields: ctx, body, reqEditors
func (_m *ClientWithResponsesInterface) PostCustomerV1OrderCreateWithResponse(ctx context.Context, body viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PostCustomerV1OrderCreateWithResponse")
	}

	var r0 *viewsLUcustomer.PostCustomerV1OrderCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error)); ok {
		return rf(ctx, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) *viewsLUcustomer.PostCustomerV1OrderCreateResponse); ok {
		r0 = rf(ctx, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viewsLUcustomer.PostCustomerV1OrderCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) error); ok {
		r1 = rf(ctx, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCustomerV1OrderCreateWithResponse'
type ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call struct {
	*mock.Call
}

// PostCustomerV1OrderCreateWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - body viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody
//   - reqEditors ...viewsLUcustomer.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) PostCustomerV1OrderCreateWithResponse(ctx interface{}, body interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call {
	return &ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call{Call: _e.mock.On("PostCustomerV1OrderCreateWithResponse",
		append([]interface{}{ctx, body}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call) Run(run func(ctx context.Context, body viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, reqEditors ...viewsLUcustomer.RequestEditorFn)) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]viewsLUcustomer.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(viewsLUcustomer.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call) Return(_a0 *viewsLUcustomer.PostCustomerV1OrderCreateResponse, _a1 error) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call) RunAndReturn(run func(context.Context, viewsLUcustomer.PostCustomerV1OrderCreateJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCreateResponse, error)) *ClientWithResponsesInterface_PostCustomerV1OrderCreateWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// PostCustomerV1OrderCustomerSignInWithBodyWithResponse provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *ClientWithResponsesInterface) PostCustomerV1OrderCustomerSignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PostCustomerV1OrderCustomerSignInWithBodyWithResponse")
	}

	var r0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCustomerV1OrderCustomerSignInWithBodyWithResponse'
type ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call struct {
	*mock.Call
}

// PostCustomerV1OrderCustomerSignInWithBodyWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - contentType string
//   - body io.Reader
//   - reqEditors ...viewsLUcustomer.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) PostCustomerV1OrderCustomerSignInWithBodyWithResponse(ctx interface{}, contentType interface{}, body interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call {
	return &ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call{Call: _e.mock.On("PostCustomerV1OrderCustomerSignInWithBodyWithResponse",
		append([]interface{}{ctx, contentType, body}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call) Run(run func(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]viewsLUcustomer.RequestEditorFn, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(viewsLUcustomer.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call) Return(_a0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, _a1 error) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call) RunAndReturn(run func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithBodyWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// PostCustomerV1OrderCustomerSignInWithResponse provides a mock function with given fields: ctx, body, reqEditors
func (_m *ClientWithResponsesInterface) PostCustomerV1OrderCustomerSignInWithResponse(ctx context.Context, body viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PostCustomerV1OrderCustomerSignInWithResponse")
	}

	var r0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error)); ok {
		return rf(ctx, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse); ok {
		r0 = rf(ctx, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) error); ok {
		r1 = rf(ctx, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCustomerV1OrderCustomerSignInWithResponse'
type ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call struct {
	*mock.Call
}

// PostCustomerV1OrderCustomerSignInWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - body viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody
//   - reqEditors ...viewsLUcustomer.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) PostCustomerV1OrderCustomerSignInWithResponse(ctx interface{}, body interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call {
	return &ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call{Call: _e.mock.On("PostCustomerV1OrderCustomerSignInWithResponse",
		append([]interface{}{ctx, body}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call) Run(run func(ctx context.Context, body viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, reqEditors ...viewsLUcustomer.RequestEditorFn)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]viewsLUcustomer.RequestEditorFn, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(viewsLUcustomer.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call) Return(_a0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, _a1 error) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call) RunAndReturn(run func(context.Context, viewsLUcustomer.PostCustomerV1OrderCustomerSignInJSONRequestBody, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignInResponse, error)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignInWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// PostCustomerV1OrderCustomerSignUpWithBodyWithResponse provides a mock function with given fields: ctx, contentType, body, reqEditors
func (_m *ClientWithResponsesInterface) PostCustomerV1OrderCustomerSignUpWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, contentType, body)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PostCustomerV1OrderCustomerSignUpWithBodyWithResponse")
	}

	var r0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse, error)); ok {
		return rf(ctx, contentType, body, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) *viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse); ok {
		r0 = rf(ctx, contentType, body, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) error); ok {
		r1 = rf(ctx, contentType, body, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PostCustomerV1OrderCustomerSignUpWithBodyWithResponse'
type ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call struct {
	*mock.Call
}

// PostCustomerV1OrderCustomerSignUpWithBodyWithResponse is a helper method to define mock.On call
//   - ctx context.Context
//   - contentType string
//   - body io.Reader
//   - reqEditors ...viewsLUcustomer.RequestEditorFn
func (_e *ClientWithResponsesInterface_Expecter) PostCustomerV1OrderCustomerSignUpWithBodyWithResponse(ctx interface{}, contentType interface{}, body interface{}, reqEditors ...interface{}) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call {
	return &ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call{Call: _e.mock.On("PostCustomerV1OrderCustomerSignUpWithBodyWithResponse",
		append([]interface{}{ctx, contentType, body}, reqEditors...)...)}
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call) Run(run func(ctx context.Context, contentType string, body io.Reader, reqEditors ...viewsLUcustomer.RequestEditorFn)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]viewsLUcustomer.RequestEditorFn, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(viewsLUcustomer.RequestEditorFn)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(io.Reader), variadicArgs...)
	})
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call) Return(_a0 *viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse, _a1 error) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call) RunAndReturn(run func(context.Context, string, io.Reader, ...viewsLUcustomer.RequestEditorFn) (*viewsLUcustomer.PostCustomerV1OrderCustomerSignUpResponse, error)) *ClientWithResponsesInterface_PostCustomerV1OrderCustomerSignUpWithBodyWithResponse_Call {
	_c.Call.Return(run)
	return _c
}

// NewClientWithResponsesInterface creates a new instance of ClientWithResponsesInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientWithResponsesInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientWithResponsesInterface {
	mock := &ClientWithResponsesInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
