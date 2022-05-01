package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth/models"
)

const USERLOGIN_KEY = "USER_LOGIN"

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through. It's just fine
// until you decide to write something similar and customize your client response.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userLoginStruct := models.AuthUserLogin{}

		userLogin, err := auth.TokenAuth.Decode(jwtauth.TokenFromHeader(r))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userLoginMap, err := userLogin.AsMap(ctx)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userLoginString, err := json.Marshal(userLoginMap)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		json.Unmarshal(userLoginString, &userLoginStruct)
		ctx = context.WithValue(ctx, USERLOGIN_KEY, userLoginStruct)

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
