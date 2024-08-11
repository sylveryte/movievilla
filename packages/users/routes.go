package users

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"movievilla.sahani.dev/internal"
)

func Register(app *internal.Application) {
	r := chi.NewRouter()
	app.Router.Mount("/users", r)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("From user"))
	})
}
