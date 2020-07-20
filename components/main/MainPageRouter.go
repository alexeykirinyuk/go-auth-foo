package main

import (
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/go-chi/chi"
	"net/http"
)

const baseTemplatePath = "views/"

func configureMainPageRouter(mux *chi.Mux) {
	mux.MethodFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		libs.Render(w, r, baseTemplatePath+"main.tpl", nil)
	})
}
