package routes

import (
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/go-chi/chi"
	"github.com/volatiletech/authboss/v3"
)

func RegisterAllRoutes(boss *authboss.Authboss, mux *chi.Mux, db data.IDatabase)  {
	configureMainPageRouter(mux)
	configureFooRouter(mux, db.GetFooStorage())
	configureBarRouter(mux, db.GetBarStorage())
	configureUserRouter(boss, mux, db.GetUserStorage())
}
