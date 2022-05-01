package user

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"
	"github.com/ibidathoillah/majoo-test/lib/types"
)

type UseCase interface {
	CreateUser(ctx context.Context, model *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, model *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID *int64) error
	GetUser(ctx context.Context, userID *int64) (*models.User, error)
	GetUserLogin(ctx context.Context, user_name string, password *types.Password) (*models.User, error)
}
