// Copyright 2023 Versity Software
// This file is licensed under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package auth

import (
	"errors"
	"fmt"

	"github.com/versity/versitygw/s3err"
)

type Role string

const (
	RoleUser     Role = "user"
	RoleAdmin    Role = "admin"
	RoleUserPlus Role = "userplus"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin:
		return true
	case RoleUser:
		return true
	case RoleUserPlus:
		return true
	default:
		return false
	}
}

// Account is a gateway IAM account
type Account struct {
	Access    string `json:"access"`
	Secret    string `json:"secret"`
	Role      Role   `json:"role"`
	UserID    int    `json:"userID"`
	GroupID   int    `json:"groupID"`
	ProjectID int    `json:"projectID"`
}

type ListUserAccountsResult struct {
	Accounts []Account
}

// Mutable props, which could be changed when updating an IAM account
type MutableProps struct {
	Secret    *string `json:"secret"`
	Role      Role    `json:"role"`
	UserID    *int    `json:"userID"`
	GroupID   *int    `json:"groupID"`
	ProjectID *int    `json:"projectID"`
}

func (m MutableProps) Validate() error {
	if m.Role != "" && !m.Role.IsValid() {
		return s3err.GetAPIError(s3err.ErrAdminInvalidUserRole)
	}

	return nil
}

func updateAcc(acc *Account, props MutableProps) {
	if props.Secret != nil {
		acc.Secret = *props.Secret
	}
	if props.GroupID != nil {
		acc.GroupID = *props.GroupID
	}
	if props.UserID != nil {
		acc.UserID = *props.UserID
	}
	if props.ProjectID != nil {
		acc.ProjectID = *props.ProjectID
	}
	if props.Role != "" {
		acc.Role = props.Role
	}
}

// IAMService is the interface for all IAM service implementations
//
//go:generate moq -out ../s3api/controllers/iam_moq_test.go -pkg controllers . IAMService
type IAMService interface {
	CreateAccount(account Account) error
	GetUserAccount(access string) (Account, error)
	UpdateUserAccount(access string, props MutableProps) error
	DeleteUserAccount(access string) error
	ListUserAccounts() ([]Account, error)
	Shutdown() error
}

var (
	// ErrUserExists is returned when the user already exists
	ErrUserExists = errors.New("user already exists")
	// ErrNoSuchUser is returned when the user does not exist
	ErrNoSuchUser = errors.New("user not found")
)

type Opts struct {
	RootAccount Account
	Dir         string
}

func New(o *Opts) (IAMService, error) {
	switch {
	case o.Dir != "":
		svc, err := NewInternal(o.RootAccount, o.Dir)
		if err != nil {
			return nil, err
		}
		fmt.Printf("initializing internal IAM with %q\n", o.Dir)
		return svc, nil
	default:
		fmt.Println("No IAM service configured, enabling single account mode")
		return NewIAMServiceSingle(o.RootAccount), nil
	}
}