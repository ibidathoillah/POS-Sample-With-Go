package http

import (
	"net/http"

	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/ibidathoillah/majoo-test/cmd/container"
	"github.com/ibidathoillah/majoo-test/cmd/http/middleware"
	authHttpTransport "github.com/ibidathoillah/majoo-test/internal/domains/auth/transports/http"
	transactionHttpTransport "github.com/ibidathoillah/majoo-test/internal/domains/transaction/transports/http"
	userHttpTransport "github.com/ibidathoillah/majoo-test/internal/domains/user/transports/http"
	libServer "github.com/ibidathoillah/majoo-test/lib/server"
)

func MakeHandler(
	router *chi.Mux,
	container container.Container) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(libServer.ErrorEncoder),
	}

	router.Group(func(r chi.Router) {

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", authHttpTransport.GetAuth(container.AuthService, opts))
		})

		r.Route("/user", func(r chi.Router) {
			r.Post("/register", userHttpTransport.RegisterUser(container.UserService, opts))
		})

		r.Route("/transactions", func(r chi.Router) {
			r.Use(middleware.Authenticator)
			r.Get("/omzets", transactionHttpTransport.FindAllOmzetReport(container.TransactionService, opts))
		})
	})

	return router
}
