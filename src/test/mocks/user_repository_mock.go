// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/repository/user_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	rest_err "crud-go/src/configuration/rest_err"
	model "crud-go/src/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", userDomain)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), userDomain)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(arg0 string) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), arg0)
}

// FindUserByEmail mocks base method.
func (m *MockUserRepository) FindUserByEmail(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepositoryMockRecorder) FindUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindUserByEmail), arg0)
}

// FindUserById mocks base method.
func (m *MockUserRepository) FindUserById(arg0 string) (model.UserDomainInterface, *rest_err.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", arg0)
	ret0, _ := ret[0].(model.UserDomainInterface)
	ret1, _ := ret[1].(*rest_err.RestErr)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockUserRepositoryMockRecorder) FindUserById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockUserRepository)(nil).FindUserById), arg0)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", userId, userDomain)
	ret0, _ := ret[0].(*rest_err.RestErr)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(userId, userDomain interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), userId, userDomain)
}