package http

import (
	"net/http"

	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/ibidathoillah/majoo-test/internal/domains/user"
	"github.com/ibidathoillah/majoo-test/internal/domains/user/endpoints"
	"github.com/ibidathoillah/majoo-test/internal/domains/user/models"
	"github.com/ibidathoillah/majoo-test/lib/server"
)

func CreateUser(service user.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.CreateUser(service), server.HTTPOption{
			DecodeModel: &models.User{},
		}, opts).ServeHTTP(w, r)
	}
}

func RegisterUser(service user.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.RegisterUser(service), server.HTTPOption{
			DecodeModel: &models.UserRegister{},
		}, opts).ServeHTTP(w, r)
	}
}

func LoginUser(service user.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.LoginUser(service), server.HTTPOption{
			DecodeModel: &models.UserLogin{},
		}, opts).ServeHTTP(w, r)
	}
}
