// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// EmailFramework is an autogenerated mock type for the EmailFramework type
type EmailFramework struct {
	mock.Mock
}

type EmailFramework_Expecter struct {
	mock *mock.Mock
}

func (_m *EmailFramework) EXPECT() *EmailFramework_Expecter {
	return &EmailFramework_Expecter{mock: &_m.Mock}
}

// SendEmail provides a mock function with given fields: subject, content, to, attchFiles
func (_m *EmailFramework) SendEmail(subject string, content string, to []string, attchFiles []string) error {
	ret := _m.Called(subject, content, to, attchFiles)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, []string, []string) error); ok {
		r0 = rf(subject, content, to, attchFiles)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EmailFramework_SendEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmail'
type EmailFramework_SendEmail_Call struct {
	*mock.Call
}

// SendEmail is a helper method to define mock.On call
//   - subject string
//   - content string
//   - to []string
//   - attchFiles []string
func (_e *EmailFramework_Expecter) SendEmail(subject interface{}, content interface{}, to interface{}, attchFiles interface{}) *EmailFramework_SendEmail_Call {
	return &EmailFramework_SendEmail_Call{Call: _e.mock.On("SendEmail", subject, content, to, attchFiles)}
}

func (_c *EmailFramework_SendEmail_Call) Run(run func(subject string, content string, to []string, attchFiles []string)) *EmailFramework_SendEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].([]string), args[3].([]string))
	})
	return _c
}

func (_c *EmailFramework_SendEmail_Call) Return(_a0 error) *EmailFramework_SendEmail_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EmailFramework_SendEmail_Call) RunAndReturn(run func(string, string, []string, []string) error) *EmailFramework_SendEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewEmailFramework interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmailFramework creates a new instance of EmailFramework. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmailFramework(t mockConstructorTestingTNewEmailFramework) *EmailFramework {
	mock := &EmailFramework{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}