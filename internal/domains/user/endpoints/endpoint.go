package endpoints

import (
	"context"

	"github.com/ibidathoillah/majoo-test/lib/database"
	"github.com/ibidathoillah/majoo-test/lib/validator"

	"github.com/ibidathoillah/majoo-test/internal/domains/user"
	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"

	"github.com/go-kit/kit/endpoint"
	"github.com/ibidathoillah/majoo-test/internal/globals"
	libHttp "github.com/ibidathoillah/majoo-test/lib/transport/http"
)

func CreateUser(service user.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var data *models.User
		payload := request.(*models.User)

		validate := validator.New()
		err = validate.Struct(payload)
		if err == nil {
			err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
				data, err = service.CreateUser(ctx, payload)
				return err
			})

		}

		return libHttp.Response(ctx, data, nil), err
	}
}

func RegisterUser(service user.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var data *models.User
		payload := request.(*models.UserRegister)

		register := models.User{
			Name:     payload.Name,
			Username: payload.Username,
			Password: payload.Password,
		}

		validate := validator.New()
		err = validate.Struct(register)
		if err == nil {
			err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
				data, err = service.CreateUser(ctx, &register)
				return err
			})
		}

		return libHttp.Response(ctx, data, nil), err
	}
}

func LoginUser(service user.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var data *models.User
		payload := request.(*models.UserLogin)

		err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
			data, err = service.GetUserLogin(ctx, payload.Username, payload.Password)
			return err
		})

		return libHttp.Response(ctx, data, nil), err
	}
}
