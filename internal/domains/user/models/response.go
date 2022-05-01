package models

import (
	"time"

	"github.com/ibidathoillah/majoo-test/lib/types"
)

type User struct {
	ID        int64           `db:"id,omitempty" json:"id"`
	Name      string          `db:"name" json:"name" validate:"required"`
	Username  string          `db:"user_name" json:"user_name" validate:"required,slug"`
	Password  *types.Password `db:"password" json:"password,omitempty" validate:"required"`
	CreatedBy int64           `db:"created_by" json:"created_by"`
	UpdatedBy int64           `db:"updated_by" json:"updated_by"`
	CreatedAt *time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time      `db:"updated_at" json:"updated_at"`
}
