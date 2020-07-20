package main

import (
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	"github.com/alexeykirinyuk/tech-task-go/components/bar"
	"github.com/alexeykirinyuk/tech-task-go/components/foo"
	"github.com/alexeykirinyuk/tech-task-go/components/sigma"
	"github.com/alexeykirinyuk/tech-task-go/components/userRole"
	"github.com/alexeykirinyuk/tech-task-go/data/postgres"
	"github.com/go-chi/chi"
	"net/http"
)


func main() {
	dbProvider := postgres.NewProvider()

	boss, err := auth.ConfigureAuth(dbProvider)
	if err != nil {
		panic(err)
	}

	// TODO: add logging
	mux := chi.NewMux()
	auth.ConfigureMiddleware(mux, boss)

	userRole.ConfigureRouter(mux, boss, dbProvider)
	bar.ConfigureRouter(mux, boss, dbProvider)
	foo.ConfigureRouter(mux, boss, dbProvider)
	sigma.ConfigureRouter(mux, boss, dbProvider)

	err = http.ListenAndServe("localhost:5000", mux)
	if err != nil {
		panic(err)
	}
}


