package http

import (
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/endpoints"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
	"github.com/ibidathoillah/majoo-test/lib/server"
)

func FindAllOmzetReport(service transaction.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.FindAllOmzetReport(service), server.HTTPOption{
			DecodeModel: &models.FindAllOmzetReport{},
		}, opts).ServeHTTP(w, r)
	}
}
