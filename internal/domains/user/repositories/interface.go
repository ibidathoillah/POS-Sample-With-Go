package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"
)

type Interface interface {
	CreateUser(ctx context.Context, model *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, model *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userID *int64) error
	GetUser(ctx context.Context, userID *int64) (*models.User, error)
	GetUserByUsername(ctx context.Context, userName string) (*models.User, error)
}
