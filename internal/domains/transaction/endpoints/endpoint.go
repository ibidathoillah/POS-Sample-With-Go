package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ibidathoillah/majoo-test/cmd/http/middleware"
	authModels "github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
	"github.com/ibidathoillah/majoo-test/internal/globals"
	"github.com/ibidathoillah/majoo-test/lib/database"
	libHttp "github.com/ibidathoillah/majoo-test/lib/transport/http"
	"github.com/ibidathoillah/majoo-test/lib/validator"
)

func FindAllOmzetReport(service transaction.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		payload := request.(*models.FindAllOmzetReport)
		var data []*models.TransactionOmzet

		user := ctx.Value(middleware.USERLOGIN_KEY).(authModels.AuthUserLogin)
		payload.UserID = user.UserID

		if payload.GroupBy == "" {
			payload.GroupBy = "merchant"
		}

		if payload.Page == 0 {
			payload.Page = 1
		}

		if payload.Limit == 0 {
			payload.Limit = 10
		}
		validate := validator.New()
		err = validate.Struct(payload)
		if err == nil {
			err = database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
				data, err = service.FindAllOmzetReport(ctx, *payload)
				return err
			})
		}

		return libHttp.Response(ctx, data, nil), err
	}
}
