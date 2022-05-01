package models

import (
	"time"

	"github.com/ibidathoillah/majoo-test/lib/types"
)

type AuthorizationType string

const (
	Bearer AuthorizationType = "Bearer"
)

type AuthReponse struct {
	Authorization AuthorizationType `json:"authorization"`
	Token         string            `json:"token"`
}

type AuthGetAuth struct {
	Username string          `json:"user_name" validate:"required"`
	Password *types.Password `json:"password" validate:"required"`
}

type AuthUserLogin struct {
	UserID         *int64    `json:"user_id"`
	OrganizationID int64     `json:"organization_id"`
	Username       string    `json:"user_name"`
	Expired        time.Time `json:"expired"`
}
