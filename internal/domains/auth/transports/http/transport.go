package http

import (
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth/endpoints"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
	"github.com/ibidathoillah/majoo-test/lib/server"
)

func GetAuth(service auth.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.GetAuth(service), server.HTTPOption{
			DecodeModel: &models.AuthGetAuth{},
		}, opts).ServeHTTP(w, r)
	}
}
