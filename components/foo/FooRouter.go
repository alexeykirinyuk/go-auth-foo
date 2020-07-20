package foo

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
	"net/http"
)

const fooTemplateBasePath = "components/foo/templates"

func ConfigureRouter(mux *chi.Mux, boss *authboss.Authboss, dbProvider data.IDatabaseProvider) {
	mux.Group(func(r chi.Router) {
		libs.ConfigureAuthMiddleware(mux, boss, auth.RoleMember, auth.RoleAdmin)
		configure(mux, dbProvider)
	})
}

func configure(r chi.Router, dbProvider data.IDatabaseProvider) {
	r.MethodFunc("GET", "/foo", func(w http.ResponseWriter, r *http.Request) {
		fooStorage := newStorage(dbProvider)

		items, err := fooStorage.getAll()
		if err != nil {
			panic(err)
		}

		libs.Render(w, r, fooTemplateBasePath+"view.tpl", items)
	})
	r.MethodFunc("GET", "/foo/create", func(w http.ResponseWriter, r *http.Request) {
		libs.Render(w, r, "views/foo/create.tpl", nil)
	})
	r.MethodFunc("POST", "/foo/create", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		foo := extractFooFromFormData(r)

		errs, ok := service.create(foo)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllFoo(w, r)
	})
	r.MethodFunc("POST", "/foo/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := extractIdFromRouteParameters(w, r)
		if !ok {
			return
		}

		errs, ok := service.delete(id)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllFoo(w, r)
	})
	r.MethodFunc("GET", "/foo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := extractIdFromRouteParameters(w, r)
		if !ok {
			return
		}

		foo, errs, ok := service.get(id)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		libs.Render(w, r, fooTemplateBasePath+"update.tpl", foo)
	})
	r.MethodFunc("POST", "/foo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := extractIdFromRouteParameters(w, r)
		if !ok {
			return
		}

		foo := extractFooFromFormData(r)
		foo.Id = id

		errs, ok := service.update(foo)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllFoo(w, r)
	})
}

func extractFooFromFormData(r *http.Request) foo {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	return foo{
		Id:          uuid.New(),
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
}

func extractIdFromRouteParameters(w http.ResponseWriter, r *http.Request) (id uuid.UUID, ok bool) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		libs.BadRequest(w, r, fmt.Sprintf("can't parse id '%s'", idStr))
		return uuid.UUID{}, false
	}

	return id, true
}

func redirectToAllFoo(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/foo", http.StatusMovedPermanently)
}
