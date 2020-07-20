package routes

import (
	"github.com/go-chi/chi"
	"net/http"
)

const baseTemplatePath = "views/"

func configureMainPageRouter(mux *chi.Mux) {
	mux.MethodFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		Render(w, r, baseTemplatePath+"main.tpl", nil)
	})
}
