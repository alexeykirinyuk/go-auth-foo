package routes

import (
"github.com/alexeykirinyuk/tech-task-go/data"
"github.com/alexeykirinyuk/tech-task-go/data/models"
"github.com/go-chi/chi"
"github.com/google/uuid"
"net/http"
)

const barTemplateBasePath = baseTemplatePath + "bar/"

func configureBarRouter(mux *chi.Mux, storage data.IBarStorage) {
	mux.MethodFunc("GET", "/bar", func(w http.ResponseWriter, r *http.Request) {
		items, err := storage.GetAll()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		Render(w, r, barTemplateBasePath + "view.tpl", items)
	})
	mux.MethodFunc("GET", "/bar/create", func(w http.ResponseWriter, r *http.Request) {
		Render(w, r, "views/bar/create.tpl", nil)
	})
	mux.MethodFunc("POST", "/bar/create", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		bar := models.Bar{
			Id:          uuid.New(),
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		err = storage.Add(bar)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllBar(w, r)
	})
	mux.MethodFunc("POST", "/bar/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		err = storage.Delete(id)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllBar(w, r)
	})
	mux.MethodFunc("GET", "/bar/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		item, err := storage.GetById(id)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		Render(w, r, barTemplateBasePath + "update.tpl", item)
	})
	mux.MethodFunc("POST", "/bar/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		err = r.ParseForm()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		bar := models.Bar{
			Id:          id,
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		err = storage.Update(bar)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllBar(w, r)
	})
}

func RedirectToAllBar(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "/bar", http.StatusMovedPermanently)
}