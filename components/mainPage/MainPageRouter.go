package mainPage

import (
	"github.com/alexeykirinyuk/tech-task-go/libs"
	"github.com/go-chi/chi"
	"net/http"
)

const baseTemplatePath = "components/mainPage/"

func ConfigureRouter(mux *chi.Mux) {
	mux.MethodFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		libs.Render(w, baseTemplatePath+"main.tpl", nil)
	})
}
