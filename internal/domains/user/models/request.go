package models

import (
	"github.com/ibidathoillah/majoo-test/lib/types"
)

type UserRegister struct {
	Name     string          `db:"name" json:"name" validate:"required"`
	Username string          `db:"user_name" json:"user_name" validate:"required"`
	Password *types.Password `db:"password" json:"password" validate:"required"`
}

type UserLogin struct {
	Username string          `db:"user_name" json:"user_name" validate:"required"`
	Password *types.Password `db:"password" json:"password" validate:"required"`
}
