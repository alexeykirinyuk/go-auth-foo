package libs

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	"github.com/go-chi/chi"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/confirm"
	"github.com/volatiletech/authboss/v3/lock"
	"net/http"
)

func ConfigureAuthMiddleware(r chi.Router, boss *authboss.Authboss, roles ...string)  {
	r.Use(authboss.Middleware2(boss, authboss.RequireNone, authboss.RespondUnauthorized))
	r.Use(lock.Middleware(boss))
	r.Use(confirm.Middleware(boss))
	r.Use(roleMiddleware(boss, roles))
}

func roleMiddleware(b *authboss.Authboss, roles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, err := b.LoadCurrentUser(&r)
			if err != nil {
				// TODO
				panic(err)
			}

			user, ok := u.(*auth.User)
			if !ok {
				// TODO
				panic(fmt.Errorf("test error"))
			}

			if !hasRole(user, roles) {
				// TODO
				panic(fmt.Errorf("test error"))
			}

			next.ServeHTTP(w, r)
		})
	}
}

func hasRole(user *auth.User, roles []string) bool {
	for _, role := range roles {
		if user.Role == role {
			return true
		}
	}

	return false
}