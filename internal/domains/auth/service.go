package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/go-kit/kit/log"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/user"
	"github.com/ibidathoillah/majoo-test/lib/errors"
	"github.com/ibidathoillah/majoo-test/lib/map_json"
)

var TokenAuth *jwtauth.JWTAuth = jwtauth.New("HS256", []byte("secret"), nil)

type service struct {
	actor       string
	logger      log.Logger
	userService user.UseCase
}

func (s *service) GetAuth(ctx context.Context, payload *models.AuthGetAuth) (*models.AuthReponse, error) {
	user, err := s.userService.GetUserLogin(ctx, payload.Username, payload.Password)
	if err != nil {
		return nil, errors.NewError(err, http.StatusUnauthorized, "Wrong user_name or password")
	}

	userMap, err := map_json.StructToMap(models.AuthUserLogin{
		UserID:   &user.ID,
		Username: user.Username,
		Expired:  time.Now().AddDate(0, 0, 2)})
	if err != nil {
		return nil, err
	}
	_, tokenString, err := TokenAuth.Encode(userMap)
	if err != nil {
		return nil, err
	}

	return &models.AuthReponse{
		Authorization: models.Bearer,
		Token:         tokenString,
	}, nil
}

func NewAuthService(
	logger log.Logger,
	userService user.UseCase,
) UseCase {
	service := service{
		actor:       "Auth",
		logger:      nil,
		userService: userService,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}
