package model

import "crud-go/src/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string
	EncryptPassword()
	CheckPasswordHash(string) bool
	GetJSONValue() (string, error)
	GenerateToken() (string, *rest_err.RestErr)

	SetID(string)
	SetPassword(string)
}
