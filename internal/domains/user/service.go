package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ibidathoillah/majoo-test/lib/errors"
	"github.com/ibidathoillah/majoo-test/lib/types"

	"github.com/go-kit/kit/log"
	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/user/repositories"
)

type service struct {
	actor      string
	logger     log.Logger
	repository repositories.Interface
}

func (s *service) CreateUser(ctx context.Context, payload *models.User) (*models.User, error) {

	err := payload.Password.Encrypt()
	if err != nil {
		return nil, err
	}

	user, err := s.repository.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.NewError(fmt.Errorf("Username is already registered"), http.StatusConflict, "user_is_already_exist")
	}

	userModel, err := s.repository.CreateUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (s *service) UpdateUser(ctx context.Context, payload *models.User) (*models.User, error) {
	userModel, err := s.repository.UpdateUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (s *service) DeleteUser(ctx context.Context, userID *int64) error {

	err := s.repository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUser(ctx context.Context, userID *int64) (*models.User, error) {

	userModel, err := s.repository.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return userModel, nil
}

func (s *service) GetUserLogin(ctx context.Context, userName string, password *types.Password) (*models.User, error) {

	userLogin, err := s.repository.GetUserByUsername(ctx, userName)
	if err != nil {
		return nil, err
	}

	if userLogin != nil && password != nil && userLogin.Password.Validate([]byte(string(*password))) {
		userLogin.Password = nil
		return userLogin, nil
	}

	return nil, errors.NewError(err, http.StatusUnauthorized, "get_user_login_failed")
}

func NewUserService(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor:      "USER",
		logger:     nil,
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}
