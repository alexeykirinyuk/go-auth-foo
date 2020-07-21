package main

import (
	"github.com/alexeykirinyuk/tech-task-go/components/bar"
	"github.com/alexeykirinyuk/tech-task-go/components/foo"
	"github.com/alexeykirinyuk/tech-task-go/components/mainPage"
	"github.com/alexeykirinyuk/tech-task-go/components/sigma"
	"github.com/alexeykirinyuk/tech-task-go/components/userRole"
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/go-chi/chi"
	"github.com/volatiletech/authboss/v3"
)

func configureAllRoutes(mux *chi.Mux, boss *authboss.Authboss, dbProvider data.IDatabaseProvider) {
	mainPage.ConfigureRouter(mux)
	userRole.ConfigureRouter(mux, boss, dbProvider)
	bar.ConfigureRouter(mux, boss, dbProvider)
	foo.ConfigureRouter(mux, boss, dbProvider)
	sigma.ConfigureRouter(mux, boss, dbProvider)
}
