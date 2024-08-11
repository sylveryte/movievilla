package healthcheck

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"movievilla.sahani.dev/internal"
)

func Register(app *internal.Application) {
	r := chi.NewRouter()
	app.Router.Mount("/healthcheck", r)

  r.Get("/{id}", healthCheckHandler)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r,"id")
	fmt.Println("Id is ", id)

	fmt.Fprintf(w, "{kversion:\"%s\"}",internal.Version)
}
