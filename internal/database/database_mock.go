// Code generated by mockery v2.32.3. DO NOT EDIT.

package database

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

type MockDatabase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDatabase) EXPECT() *MockDatabase_Expecter {
	return &MockDatabase_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockDatabase) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockDatabase_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockDatabase_Expecter) Close() *MockDatabase_Close_Call {
	return &MockDatabase_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockDatabase_Close_Call) Run(run func()) *MockDatabase_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDatabase_Close_Call) Return(_a0 error) *MockDatabase_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Close_Call) RunAndReturn(run func() error) *MockDatabase_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CreateDocument provides a mock function with given fields: ctx, id, content
func (_m *MockDatabase) CreateDocument(ctx context.Context, id string, content string) error {
	ret := _m.Called(ctx, id, content)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_CreateDocument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDocument'
type MockDatabase_CreateDocument_Call struct {
	*mock.Call
}

// CreateDocument is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - content string
func (_e *MockDatabase_Expecter) CreateDocument(ctx interface{}, id interface{}, content interface{}) *MockDatabase_CreateDocument_Call {
	return &MockDatabase_CreateDocument_Call{Call: _e.mock.On("CreateDocument", ctx, id, content)}
}

func (_c *MockDatabase_CreateDocument_Call) Run(run func(ctx context.Context, id string, content string)) *MockDatabase_CreateDocument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockDatabase_CreateDocument_Call) Return(_a0 error) *MockDatabase_CreateDocument_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_CreateDocument_Call) RunAndReturn(run func(context.Context, string, string) error) *MockDatabase_CreateDocument_Call {
	_c.Call.Return(run)
	return _c
}

// GetDocument provides a mock function with given fields: ctx, id
func (_m *MockDatabase) GetDocument(ctx context.Context, id string) (Document, error) {
	ret := _m.Called(ctx, id)

	var r0 Document
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (Document, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) Document); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(Document)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDatabase_GetDocument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDocument'
type MockDatabase_GetDocument_Call struct {
	*mock.Call
}

// GetDocument is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockDatabase_Expecter) GetDocument(ctx interface{}, id interface{}) *MockDatabase_GetDocument_Call {
	return &MockDatabase_GetDocument_Call{Call: _e.mock.On("GetDocument", ctx, id)}
}

func (_c *MockDatabase_GetDocument_Call) Run(run func(ctx context.Context, id string)) *MockDatabase_GetDocument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockDatabase_GetDocument_Call) Return(_a0 Document, _a1 error) *MockDatabase_GetDocument_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDatabase_GetDocument_Call) RunAndReturn(run func(context.Context, string) (Document, error)) *MockDatabase_GetDocument_Call {
	_c.Call.Return(run)
	return _c
}

// Migrate provides a mock function with given fields: ctx
func (_m *MockDatabase) Migrate(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDatabase_Migrate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Migrate'
type MockDatabase_Migrate_Call struct {
	*mock.Call
}

// Migrate is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockDatabase_Expecter) Migrate(ctx interface{}) *MockDatabase_Migrate_Call {
	return &MockDatabase_Migrate_Call{Call: _e.mock.On("Migrate", ctx)}
}

func (_c *MockDatabase_Migrate_Call) Run(run func(ctx context.Context)) *MockDatabase_Migrate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockDatabase_Migrate_Call) Return(_a0 error) *MockDatabase_Migrate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDatabase_Migrate_Call) RunAndReturn(run func(context.Context) error) *MockDatabase_Migrate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDatabase creates a new instance of MockDatabase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDatabase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDatabase {
	mock := &MockDatabase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
