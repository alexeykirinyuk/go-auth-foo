package routes

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/confirm"
	"github.com/volatiletech/authboss/v3/lock"
	"net/http"
)

const userTemplateBasePath = baseTemplatePath + "user/"

func configureUserRouter(boss *authboss.Authboss, mux *chi.Mux, storage data.IUserStorage) {
	mux.Group(func(r chi.Router) {
		r.Use(authboss.Middleware2(boss, authboss.RequireNone, authboss.RespondUnauthorized), lock.Middleware(boss), confirm.Middleware(boss))
		r.Use(roleMiddleware(boss, models.RoleMember))

		setUpRoutes(r, storage)
	})
}

func setUpRoutes(mux chi.Router, storage data.IUserStorage) {
	mux.MethodFunc("GET", "/user", func(w http.ResponseWriter, r *http.Request) {
		items, err := storage.GetAll()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		Render(w, r, userTemplateBasePath+"view.tpl", items)
	})
	mux.MethodFunc("GET", "/user/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			InternalServerError(w, r)
			return
		}

		item, err := storage.GetById(id)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		Render(w, r, userTemplateBasePath+"update.tpl", item)
	})
	mux.MethodFunc("POST", "/user/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			InternalServerError(w, r)
			return
		}

		err = r.ParseForm()
		if err != nil {
			InternalServerError(w, r)
			return
		}
		role := r.FormValue("role")

		if role != models.RoleMember && role != models.RoleAdmin {
			InternalServerError(w, r)
			return
		}

		user, err := storage.GetById(id)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		user.Role = role
		err = storage.Update(user)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllUser(w, r)
	})
}

func RedirectToAllUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/user", http.StatusMovedPermanently)
}
func roleMiddleware(b *authboss.Authboss, role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, err := b.LoadCurrentUser(&r)
			if err != nil {
				// TODO
				panic(err)
			}

			user, ok := u.(*models.User)
			if !ok {
				// TODO
				panic(fmt.Errorf("test error"))
			}
			if user.Role != role {
				// TODO
				panic(fmt.Errorf("test error"))
			}

			next.ServeHTTP(w, r)
		})
	}
}
