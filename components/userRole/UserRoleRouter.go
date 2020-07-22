package userRole

import (
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/volatiletech/authboss/v3"
	"net/http"
)

const (
	userTemplateBasePath = "components/userRole/templates/"
	baseRoute            = "/user"
)

type updateViewModel struct {
	User  userRole
	Roles []roleViewModel
}

type roleViewModel struct {
	Name     string
	Selected bool
}

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
		userRoleStorage := newStorage(dbProvider)

		items, err := userRoleStorage.getAll()
		if err != nil {
			panic(err)
		}

		libs.Render(w, userTemplateBasePath+"view.tpl", items)
	})
	mux.MethodFunc("GET", "/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := getIdFromRouteParameter(w, r)
		if !ok {
			return
		}

		item, errs, ok := service.get(id)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		isAdmin := item.Role == RoleAdmin
		viewModel := &updateViewModel{
			User: item,
			Roles: []roleViewModel{
				{
					Name:     RoleAdmin,
					Selected: isAdmin,
				},
				{
					Name:     RoleMember,
					Selected: !isAdmin,
				},
			},
		}

		libs.Render(w, userTemplateBasePath+"update.tpl", viewModel)
	})
	mux.MethodFunc("POST", "/{id}", func(w http.ResponseWriter, r *http.Request) {
		service := newService(dbProvider)

		id, ok := getIdFromRouteParameter(w, r)
		if !ok {
			return
		}

		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		role := r.FormValue("role")

		errs, ok := service.updateRole(id, role)
		if !ok {
			libs.BadRequest(w, r, libs.ToResponse(errs))
			return
		}

		redirectToAllUser(w, r)
	})
}

func getIdFromRouteParameter(w http.ResponseWriter, r *http.Request) (id uuid.UUID, ok bool) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		libs.BadRequest(w, r, "Can't parse Id to UUID")
		return uuid.UUID{}, false
	}

	return id, true
}

func redirectToAllUser(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, baseRoute, http.StatusMovedPermanently)
}
