// Package oapi-codegen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package strictserver

import (
	"time"
)

// CreateUser defines model for CreateUser.
type CreateUser struct {
	Email    *string `json:"email,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// User defines model for User.
type User struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Email     *string    `json:"email,omitempty"`
	Id        int64      `json:"id"`
	Lastname  *string    `json:"lastname,omitempty"`
	Name      *string    `json:"name,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody = CreateUser

// PutApiV1UsersUserIdJSONRequestBody defines body for PutApiV1UsersUserId for application/json ContentType.
type PutApiV1UsersUserIdJSONRequestBody = User