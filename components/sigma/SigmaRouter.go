package sigma

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

const (
	sigmaTemplateBasePath = "components/sigma/templates/"
	baseRoute             = "/sigma"
)

func ConfigureRouter(mux *chi.Mux, boss *authboss.Authboss, dbProvider data.IDatabaseProvider) {
	mux.Group(func(r chi.Router) {
		libs.ConfigureAuthMiddleware(r, boss, auth.RoleAdmin)
		r.Route(baseRoute, func(r1 chi.Router) {
			configure(r1, dbProvider)
		})
	})
}

func configure(mux chi.Router, dbProvider data.IDatabaseProvider) {
	mux.MethodFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		sigmaStorage := newStorage(dbProvider)

		items, err := sigmaStorage.getAll()
		if err != nil {
			panic(err)
		}

		libs.Render(w, sigmaTemplateBasePath+"view.tpl", items)
	})
	mux.MethodFunc("GET", "/create", func(w http.ResponseWriter, r *http.Request) {
		libs.Render(w, sigmaTemplateBasePath+"create.tpl", nil)
	})
	mux.MethodFunc("POST", "/create", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		sigma := extractSigmaFromFormData(r)

		errs, ok := service.create(sigma)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllSigma(w, r)
	})
	mux.MethodFunc("POST", "/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
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

		redirectToAllSigma(w, r)
	})
	mux.MethodFunc("GET", "/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := extractIdFromRouteParameters(w, r)
		if !ok {
			return
		}

		sigma, errs, ok := service.get(id)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		libs.Render(w, sigmaTemplateBasePath+"update.tpl", sigma)
	})
	mux.MethodFunc("POST", "/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := extractIdFromRouteParameters(w, r)
		if !ok {
			return
		}

		sigma := extractSigmaFromFormData(r)
		sigma.Id = id

		errs, ok := service.update(sigma)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllSigma(w, r)
	})
}

func extractSigmaFromFormData(r *http.Request) sigma {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	return sigma{
		Id:   uuid.New(),
		Info: r.FormValue("info"),
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

func redirectToAllSigma(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/sigma", http.StatusMovedPermanently)
}
