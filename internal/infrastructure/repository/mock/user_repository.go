// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock

import (
	context "context"

	domain "github.com/abc-valera/flugo-api/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: c, user
func (_m *UserRepository) CreateUser(c context.Context, user *domain.User) error {
	ret := _m.Called(c, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) error); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - c context.Context
//   - user *domain.User
func (_e *UserRepository_Expecter) CreateUser(c interface{}, user interface{}) *UserRepository_CreateUser_Call {
	return &UserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", c, user)}
}

func (_c *UserRepository_CreateUser_Call) Run(run func(c context.Context, user *domain.User)) *UserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*domain.User))
	})
	return _c
}

func (_c *UserRepository_CreateUser_Call) Return(_a0 error) *UserRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_CreateUser_Call) RunAndReturn(run func(context.Context, *domain.User) error) *UserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: c, username
func (_m *UserRepository) DeleteUser(c context.Context, username string) error {
	ret := _m.Called(c, username)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserRepository_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - c context.Context
//   - username string
func (_e *UserRepository_Expecter) DeleteUser(c interface{}, username interface{}) *UserRepository_DeleteUser_Call {
	return &UserRepository_DeleteUser_Call{Call: _e.mock.On("DeleteUser", c, username)}
}

func (_c *UserRepository_DeleteUser_Call) Run(run func(c context.Context, username string)) *UserRepository_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_DeleteUser_Call) Return(_a0 error) *UserRepository_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_DeleteUser_Call) RunAndReturn(run func(context.Context, string) error) *UserRepository_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: c, email
func (_m *UserRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ret := _m.Called(c, email)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type UserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - c context.Context
//   - email string
func (_e *UserRepository_Expecter) GetUserByEmail(c interface{}, email interface{}) *UserRepository_GetUserByEmail_Call {
	return &UserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", c, email)}
}

func (_c *UserRepository_GetUserByEmail_Call) Run(run func(c context.Context, email string)) *UserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_GetUserByEmail_Call) Return(_a0 *domain.User, _a1 error) *UserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*domain.User, error)) *UserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: c, username
func (_m *UserRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ret := _m.Called(c, username)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type UserRepository_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - c context.Context
//   - username string
func (_e *UserRepository_Expecter) GetUserByUsername(c interface{}, username interface{}) *UserRepository_GetUserByUsername_Call {
	return &UserRepository_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", c, username)}
}

func (_c *UserRepository_GetUserByUsername_Call) Run(run func(c context.Context, username string)) *UserRepository_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_GetUserByUsername_Call) Return(_a0 *domain.User, _a1 error) *UserRepository_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetUserByUsername_Call) RunAndReturn(run func(context.Context, string) (*domain.User, error)) *UserRepository_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// SearchUsersByUsername provides a mock function with given fields: c, username, params
func (_m *UserRepository) SearchUsersByUsername(c context.Context, username string, params *domain.SelectParams) (domain.Users, error) {
	ret := _m.Called(c, username, params)

	var r0 domain.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.SelectParams) (domain.Users, error)); ok {
		return rf(c, username, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.SelectParams) domain.Users); ok {
		r0 = rf(c, username, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *domain.SelectParams) error); ok {
		r1 = rf(c, username, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_SearchUsersByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchUsersByUsername'
type UserRepository_SearchUsersByUsername_Call struct {
	*mock.Call
}

// SearchUsersByUsername is a helper method to define mock.On call
//   - c context.Context
//   - username string
//   - params *domain.SelectParams
func (_e *UserRepository_Expecter) SearchUsersByUsername(c interface{}, username interface{}, params interface{}) *UserRepository_SearchUsersByUsername_Call {
	return &UserRepository_SearchUsersByUsername_Call{Call: _e.mock.On("SearchUsersByUsername", c, username, params)}
}

func (_c *UserRepository_SearchUsersByUsername_Call) Run(run func(c context.Context, username string, params *domain.SelectParams)) *UserRepository_SearchUsersByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*domain.SelectParams))
	})
	return _c
}

func (_c *UserRepository_SearchUsersByUsername_Call) Return(_a0 domain.Users, _a1 error) *UserRepository_SearchUsersByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_SearchUsersByUsername_Call) RunAndReturn(run func(context.Context, string, *domain.SelectParams) (domain.Users, error)) *UserRepository_SearchUsersByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserBio provides a mock function with given fields: c, username, bio
func (_m *UserRepository) UpdateUserBio(c context.Context, username string, bio string) error {
	ret := _m.Called(c, username, bio)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, username, bio)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_UpdateUserBio_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserBio'
type UserRepository_UpdateUserBio_Call struct {
	*mock.Call
}

// UpdateUserBio is a helper method to define mock.On call
//   - c context.Context
//   - username string
//   - bio string
func (_e *UserRepository_Expecter) UpdateUserBio(c interface{}, username interface{}, bio interface{}) *UserRepository_UpdateUserBio_Call {
	return &UserRepository_UpdateUserBio_Call{Call: _e.mock.On("UpdateUserBio", c, username, bio)}
}

func (_c *UserRepository_UpdateUserBio_Call) Run(run func(c context.Context, username string, bio string)) *UserRepository_UpdateUserBio_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserRepository_UpdateUserBio_Call) Return(_a0 error) *UserRepository_UpdateUserBio_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_UpdateUserBio_Call) RunAndReturn(run func(context.Context, string, string) error) *UserRepository_UpdateUserBio_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserFullname provides a mock function with given fields: c, username, fullname
func (_m *UserRepository) UpdateUserFullname(c context.Context, username string, fullname string) error {
	ret := _m.Called(c, username, fullname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, username, fullname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_UpdateUserFullname_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserFullname'
type UserRepository_UpdateUserFullname_Call struct {
	*mock.Call
}

// UpdateUserFullname is a helper method to define mock.On call
//   - c context.Context
//   - username string
//   - fullname string
func (_e *UserRepository_Expecter) UpdateUserFullname(c interface{}, username interface{}, fullname interface{}) *UserRepository_UpdateUserFullname_Call {
	return &UserRepository_UpdateUserFullname_Call{Call: _e.mock.On("UpdateUserFullname", c, username, fullname)}
}

func (_c *UserRepository_UpdateUserFullname_Call) Run(run func(c context.Context, username string, fullname string)) *UserRepository_UpdateUserFullname_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserRepository_UpdateUserFullname_Call) Return(_a0 error) *UserRepository_UpdateUserFullname_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_UpdateUserFullname_Call) RunAndReturn(run func(context.Context, string, string) error) *UserRepository_UpdateUserFullname_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserHashedPassword provides a mock function with given fields: c, username, hashedPassword
func (_m *UserRepository) UpdateUserHashedPassword(c context.Context, username string, hashedPassword string) error {
	ret := _m.Called(c, username, hashedPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, username, hashedPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_UpdateUserHashedPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserHashedPassword'
type UserRepository_UpdateUserHashedPassword_Call struct {
	*mock.Call
}

// UpdateUserHashedPassword is a helper method to define mock.On call
//   - c context.Context
//   - username string
//   - hashedPassword string
func (_e *UserRepository_Expecter) UpdateUserHashedPassword(c interface{}, username interface{}, hashedPassword interface{}) *UserRepository_UpdateUserHashedPassword_Call {
	return &UserRepository_UpdateUserHashedPassword_Call{Call: _e.mock.On("UpdateUserHashedPassword", c, username, hashedPassword)}
}

func (_c *UserRepository_UpdateUserHashedPassword_Call) Run(run func(c context.Context, username string, hashedPassword string)) *UserRepository_UpdateUserHashedPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserRepository_UpdateUserHashedPassword_Call) Return(_a0 error) *UserRepository_UpdateUserHashedPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_UpdateUserHashedPassword_Call) RunAndReturn(run func(context.Context, string, string) error) *UserRepository_UpdateUserHashedPassword_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserStatus provides a mock function with given fields: c, username, status
func (_m *UserRepository) UpdateUserStatus(c context.Context, username string, status string) error {
	ret := _m.Called(c, username, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(c, username, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_UpdateUserStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserStatus'
type UserRepository_UpdateUserStatus_Call struct {
	*mock.Call
}

// UpdateUserStatus is a helper method to define mock.On call
//   - c context.Context
//   - username string
//   - status string
func (_e *UserRepository_Expecter) UpdateUserStatus(c interface{}, username interface{}, status interface{}) *UserRepository_UpdateUserStatus_Call {
	return &UserRepository_UpdateUserStatus_Call{Call: _e.mock.On("UpdateUserStatus", c, username, status)}
}

func (_c *UserRepository_UpdateUserStatus_Call) Run(run func(c context.Context, username string, status string)) *UserRepository_UpdateUserStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserRepository_UpdateUserStatus_Call) Return(_a0 error) *UserRepository_UpdateUserStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_UpdateUserStatus_Call) RunAndReturn(run func(context.Context, string, string) error) *UserRepository_UpdateUserStatus_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}