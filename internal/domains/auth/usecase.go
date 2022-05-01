package auth

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
)

type UseCase interface {
	GetAuth(ctx context.Context, model *models.AuthGetAuth) (*models.AuthReponse, error)
}
