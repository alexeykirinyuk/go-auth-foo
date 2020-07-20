package routes

import (
	"github.com/alexeykirinyuk/tech-task-go/data"
	"github.com/alexeykirinyuk/tech-task-go/data/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

const fooTemplateBasePath = baseTemplatePath + "foo/"

func configureFooRouter(mux *chi.Mux, storage data.IFooStorage) {
	mux.MethodFunc("GET", "/foo", func(w http.ResponseWriter, r *http.Request) {
		items, err := storage.GetAll()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		Render(w, r, fooTemplateBasePath+ "view.tpl", items)
	})
	mux.MethodFunc("GET", "/foo/create", func(w http.ResponseWriter, r *http.Request) {
		Render(w, r, "views/foo/create.tpl", nil)
	})
	mux.MethodFunc("POST", "/foo/create", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			InternalServerError(w, r)
			return
		}

		foo := models.Foo{
			Id:          uuid.New(),
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		err = storage.Add(foo)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllFoo(w, r)
	})
	mux.MethodFunc("POST", "/foo/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
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

		RedirectToAllFoo(w, r)
	})
	mux.MethodFunc("GET", "/foo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
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

		Render(w, r, fooTemplateBasePath+ "update.tpl", item)
	})
	mux.MethodFunc("POST", "/foo/update/{id}", func(w http.ResponseWriter, r *http.Request) {
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

		foo := models.Foo{
			Id:          id,
			Title:       r.FormValue("title"),
			Description: r.FormValue("description"),
		}

		err = storage.Update(foo)
		if err != nil {
			InternalServerError(w, r)
			return
		}

		RedirectToAllFoo(w, r)
	})
}

func RedirectToAllFoo(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, "/foo", http.StatusMovedPermanently)
}