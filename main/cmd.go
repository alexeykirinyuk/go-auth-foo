package main

import (
	"fmt"
	"github.com/alexeykirinyuk/tech-task-go/components/auth"
	config2 "github.com/alexeykirinyuk/tech-task-go/config"
	"github.com/alexeykirinyuk/tech-task-go/data/postgres"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	config, err := config2.GetConfig()
	if err != nil {
		panic(err)
	}

	dbProvider := postgres.NewProvider(config.ConnectionString)
	err = runAllMigrations(dbProvider)
	if err != nil {
		panic(err)
	}

	boss, err := auth.ConfigureAuth(dbProvider, config.Port)
	if err != nil {
		panic(err)
	}

	mux := chi.NewMux()
	auth.ConfigureMiddleware(mux, boss)

	configureAllRoutes(mux, boss, dbProvider)

	url := fmt.Sprintf("localhost:%d", config.Port)
	err = http.ListenAndServe(url, mux)
	if err != nil {
		panic(err)
	}
}
