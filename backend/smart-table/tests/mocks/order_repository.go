// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/smart-table/src/domains/customer/domain"
	mock "github.com/stretchr/testify/mock"

	utils "github.com/smart-table/src/utils"

	uuid "github.com/google/uuid"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

type OrderRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *OrderRepository) EXPECT() *OrderRepository_Expecter {
	return &OrderRepository_Expecter{mock: &_m.Mock}
}

// Begin provides a mock function with given fields: ctx
func (_m *OrderRepository) Begin(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type OrderRepository_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OrderRepository_Expecter) Begin(ctx interface{}) *OrderRepository_Begin_Call {
	return &OrderRepository_Begin_Call{Call: _e.mock.On("Begin", ctx)}
}

func (_c *OrderRepository_Begin_Call) Run(run func(ctx context.Context)) *OrderRepository_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OrderRepository_Begin_Call) Return(_a0 error) *OrderRepository_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderRepository_Begin_Call) RunAndReturn(run func(context.Context) error) *OrderRepository_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: ctx
func (_m *OrderRepository) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type OrderRepository_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - ctx context.Context
func (_e *OrderRepository_Expecter) Commit(ctx interface{}) *OrderRepository_Commit_Call {
	return &OrderRepository_Commit_Call{Call: _e.mock.On("Commit", ctx)}
}

func (_c *OrderRepository_Commit_Call) Run(run func(ctx context.Context)) *OrderRepository_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *OrderRepository_Commit_Call) Return(_a0 error) *OrderRepository_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderRepository_Commit_Call) RunAndReturn(run func(context.Context) error) *OrderRepository_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// FindActiveOrderByTableID provides a mock function with given fields: ctx, tableID
func (_m *OrderRepository) FindActiveOrderByTableID(ctx context.Context, tableID string) (utils.SharedRef[domain.Order], error) {
	ret := _m.Called(ctx, tableID)

	if len(ret) == 0 {
		panic("no return value specified for FindActiveOrderByTableID")
	}

	var r0 utils.SharedRef[domain.Order]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (utils.SharedRef[domain.Order], error)); ok {
		return rf(ctx, tableID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) utils.SharedRef[domain.Order]); ok {
		r0 = rf(ctx, tableID)
	} else {
		r0 = ret.Get(0).(utils.SharedRef[domain.Order])
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, tableID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_FindActiveOrderByTableID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindActiveOrderByTableID'
type OrderRepository_FindActiveOrderByTableID_Call struct {
	*mock.Call
}

// FindActiveOrderByTableID is a helper method to define mock.On call
//   - ctx context.Context
//   - tableID string
func (_e *OrderRepository_Expecter) FindActiveOrderByTableID(ctx interface{}, tableID interface{}) *OrderRepository_FindActiveOrderByTableID_Call {
	return &OrderRepository_FindActiveOrderByTableID_Call{Call: _e.mock.On("FindActiveOrderByTableID", ctx, tableID)}
}

func (_c *OrderRepository_FindActiveOrderByTableID_Call) Run(run func(ctx context.Context, tableID string)) *OrderRepository_FindActiveOrderByTableID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *OrderRepository_FindActiveOrderByTableID_Call) Return(_a0 utils.SharedRef[domain.Order], _a1 error) *OrderRepository_FindActiveOrderByTableID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderRepository_FindActiveOrderByTableID_Call) RunAndReturn(run func(context.Context, string) (utils.SharedRef[domain.Order], error)) *OrderRepository_FindActiveOrderByTableID_Call {
	_c.Call.Return(run)
	return _c
}

// FindOrder provides a mock function with given fields: ctx, orderUUID
func (_m *OrderRepository) FindOrder(ctx context.Context, orderUUID uuid.UUID) (utils.SharedRef[domain.Order], error) {
	ret := _m.Called(ctx, orderUUID)

	if len(ret) == 0 {
		panic("no return value specified for FindOrder")
	}

	var r0 utils.SharedRef[domain.Order]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (utils.SharedRef[domain.Order], error)); ok {
		return rf(ctx, orderUUID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) utils.SharedRef[domain.Order]); ok {
		r0 = rf(ctx, orderUUID)
	} else {
		r0 = ret.Get(0).(utils.SharedRef[domain.Order])
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, orderUUID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_FindOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrder'
type OrderRepository_FindOrder_Call struct {
	*mock.Call
}

// FindOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - orderUUID uuid.UUID
func (_e *OrderRepository_Expecter) FindOrder(ctx interface{}, orderUUID interface{}) *OrderRepository_FindOrder_Call {
	return &OrderRepository_FindOrder_Call{Call: _e.mock.On("FindOrder", ctx, orderUUID)}
}

func (_c *OrderRepository_FindOrder_Call) Run(run func(ctx context.Context, orderUUID uuid.UUID)) *OrderRepository_FindOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *OrderRepository_FindOrder_Call) Return(_a0 utils.SharedRef[domain.Order], _a1 error) *OrderRepository_FindOrder_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderRepository_FindOrder_Call) RunAndReturn(run func(context.Context, uuid.UUID) (utils.SharedRef[domain.Order], error)) *OrderRepository_FindOrder_Call {
	_c.Call.Return(run)
	return _c
}

// FindOrders provides a mock function with given fields: ctx, orderUUIDs
func (_m *OrderRepository) FindOrders(ctx context.Context, orderUUIDs []uuid.UUID) ([]utils.SharedRef[domain.Order], error) {
	ret := _m.Called(ctx, orderUUIDs)

	if len(ret) == 0 {
		panic("no return value specified for FindOrders")
	}

	var r0 []utils.SharedRef[domain.Order]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]utils.SharedRef[domain.Order], error)); ok {
		return rf(ctx, orderUUIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []utils.SharedRef[domain.Order]); ok {
		r0 = rf(ctx, orderUUIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]utils.SharedRef[domain.Order])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, orderUUIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderRepository_FindOrders_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOrders'
type OrderRepository_FindOrders_Call struct {
	*mock.Call
}

// FindOrders is a helper method to define mock.On call
//   - ctx context.Context
//   - orderUUIDs []uuid.UUID
func (_e *OrderRepository_Expecter) FindOrders(ctx interface{}, orderUUIDs interface{}) *OrderRepository_FindOrders_Call {
	return &OrderRepository_FindOrders_Call{Call: _e.mock.On("FindOrders", ctx, orderUUIDs)}
}

func (_c *OrderRepository_FindOrders_Call) Run(run func(ctx context.Context, orderUUIDs []uuid.UUID)) *OrderRepository_FindOrders_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uuid.UUID))
	})
	return _c
}

func (_c *OrderRepository_FindOrders_Call) Return(_a0 []utils.SharedRef[domain.Order], _a1 error) *OrderRepository_FindOrders_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *OrderRepository_FindOrders_Call) RunAndReturn(run func(context.Context, []uuid.UUID) ([]utils.SharedRef[domain.Order], error)) *OrderRepository_FindOrders_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, order
func (_m *OrderRepository) Save(ctx context.Context, order utils.SharedRef[domain.Order]) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, utils.SharedRef[domain.Order]) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type OrderRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - order utils.SharedRef[domain.Order]
func (_e *OrderRepository_Expecter) Save(ctx interface{}, order interface{}) *OrderRepository_Save_Call {
	return &OrderRepository_Save_Call{Call: _e.mock.On("Save", ctx, order)}
}

func (_c *OrderRepository_Save_Call) Run(run func(ctx context.Context, order utils.SharedRef[domain.Order])) *OrderRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(utils.SharedRef[domain.Order]))
	})
	return _c
}

func (_c *OrderRepository_Save_Call) Return(_a0 error) *OrderRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *OrderRepository_Save_Call) RunAndReturn(run func(context.Context, utils.SharedRef[domain.Order]) error) *OrderRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
