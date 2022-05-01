package endpoints

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth"

	"github.com/go-kit/kit/endpoint"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
	"github.com/ibidathoillah/majoo-test/internal/globals"
	"github.com/ibidathoillah/majoo-test/lib/database"
	libHttp "github.com/ibidathoillah/majoo-test/lib/transport/http"
)

func GetAuth(service auth.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var data *models.AuthReponse
		payload := request.(*models.AuthGetAuth)

		validate := validator.New()
		err = validate.Struct(payload)
		if err == nil {
			err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
				data, err = service.GetAuth(ctx, payload)
				return err
			})
		}

		return libHttp.Response(ctx, data, nil), err
	}
}
