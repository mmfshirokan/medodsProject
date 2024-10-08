// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/mmfshirokan/medodsProject/internal/model"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, rft
func (_m *Repository) Add(ctx context.Context, rft model.RefreshToken) error {
	ret := _m.Called(ctx, rft)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.RefreshToken) error); ok {
		r0 = rf(ctx, rft)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Repository_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - rft model.RefreshToken
func (_e *Repository_Expecter) Add(ctx interface{}, rft interface{}) *Repository_Add_Call {
	return &Repository_Add_Call{Call: _e.mock.On("Add", ctx, rft)}
}

func (_c *Repository_Add_Call) Run(run func(ctx context.Context, rft model.RefreshToken)) *Repository_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.RefreshToken))
	})
	return _c
}

func (_c *Repository_Add_Call) Return(_a0 error) *Repository_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Add_Call) RunAndReturn(run func(context.Context, model.RefreshToken) error) *Repository_Add_Call {
	_c.Call.Return(run)
	return _c
}

// AddUsr provides a mock function with given fields: ctx, usr
func (_m *Repository) AddUsr(ctx context.Context, usr model.User) error {
	ret := _m.Called(ctx, usr)

	if len(ret) == 0 {
		panic("no return value specified for AddUsr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(ctx, usr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_AddUsr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUsr'
type Repository_AddUsr_Call struct {
	*mock.Call
}

// AddUsr is a helper method to define mock.On call
//   - ctx context.Context
//   - usr model.User
func (_e *Repository_Expecter) AddUsr(ctx interface{}, usr interface{}) *Repository_AddUsr_Call {
	return &Repository_AddUsr_Call{Call: _e.mock.On("AddUsr", ctx, usr)}
}

func (_c *Repository_AddUsr_Call) Run(run func(ctx context.Context, usr model.User)) *Repository_AddUsr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.User))
	})
	return _c
}

func (_c *Repository_AddUsr_Call) Return(_a0 error) *Repository_AddUsr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_AddUsr_Call) RunAndReturn(run func(context.Context, model.User) error) *Repository_AddUsr_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, rftID
func (_m *Repository) Delete(ctx context.Context, rftID uuid.UUID) error {
	ret := _m.Called(ctx, rftID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, rftID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Repository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - rftID uuid.UUID
func (_e *Repository_Expecter) Delete(ctx interface{}, rftID interface{}) *Repository_Delete_Call {
	return &Repository_Delete_Call{Call: _e.mock.On("Delete", ctx, rftID)}
}

func (_c *Repository_Delete_Call) Run(run func(ctx context.Context, rftID uuid.UUID)) *Repository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_Delete_Call) Return(_a0 error) *Repository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Delete_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *Repository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetHash provides a mock function with given fields: ctx, rftID
func (_m *Repository) GetHash(ctx context.Context, rftID uuid.UUID) (string, error) {
	ret := _m.Called(ctx, rftID)

	if len(ret) == 0 {
		panic("no return value specified for GetHash")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (string, error)); ok {
		return rf(ctx, rftID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) string); ok {
		r0 = rf(ctx, rftID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, rftID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHash'
type Repository_GetHash_Call struct {
	*mock.Call
}

// GetHash is a helper method to define mock.On call
//   - ctx context.Context
//   - rftID uuid.UUID
func (_e *Repository_Expecter) GetHash(ctx interface{}, rftID interface{}) *Repository_GetHash_Call {
	return &Repository_GetHash_Call{Call: _e.mock.On("GetHash", ctx, rftID)}
}

func (_c *Repository_GetHash_Call) Run(run func(ctx context.Context, rftID uuid.UUID)) *Repository_GetHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_GetHash_Call) Return(_a0 string, _a1 error) *Repository_GetHash_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetHash_Call) RunAndReturn(run func(context.Context, uuid.UUID) (string, error)) *Repository_GetHash_Call {
	_c.Call.Return(run)
	return _c
}

// GetPwd provides a mock function with given fields: ctx, id
func (_m *Repository) GetPwd(ctx context.Context, id uuid.UUID) (string, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPwd")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (string, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetPwd_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPwd'
type Repository_GetPwd_Call struct {
	*mock.Call
}

// GetPwd is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Repository_Expecter) GetPwd(ctx interface{}, id interface{}) *Repository_GetPwd_Call {
	return &Repository_GetPwd_Call{Call: _e.mock.On("GetPwd", ctx, id)}
}

func (_c *Repository_GetPwd_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Repository_GetPwd_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_GetPwd_Call) Return(_a0 string, _a1 error) *Repository_GetPwd_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetPwd_Call) RunAndReturn(run func(context.Context, uuid.UUID) (string, error)) *Repository_GetPwd_Call {
	_c.Call.Return(run)
	return _c
}

// GetWithUserID provides a mock function with given fields: ctx, id
func (_m *Repository) GetWithUserID(ctx context.Context, id uuid.UUID) ([]model.RefreshToken, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetWithUserID")
	}

	var r0 []model.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]model.RefreshToken, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []model.RefreshToken); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.RefreshToken)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetWithUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWithUserID'
type Repository_GetWithUserID_Call struct {
	*mock.Call
}

// GetWithUserID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *Repository_Expecter) GetWithUserID(ctx interface{}, id interface{}) *Repository_GetWithUserID_Call {
	return &Repository_GetWithUserID_Call{Call: _e.mock.On("GetWithUserID", ctx, id)}
}

func (_c *Repository_GetWithUserID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *Repository_GetWithUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_GetWithUserID_Call) Return(_a0 []model.RefreshToken, _a1 error) *Repository_GetWithUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetWithUserID_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]model.RefreshToken, error)) *Repository_GetWithUserID_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
