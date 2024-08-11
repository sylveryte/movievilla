package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"movievilla.sahani.dev/internal"
	"movievilla.sahani.dev/packages/healthcheck"
	"movievilla.sahani.dev/packages/users"
)

func main() {
	var cfg internal.Config
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "developmnent", "Environment developmnent, staging or production")
	flag.Parse()

	router := chi.NewRouter()
	// middlewares
	router.Use(middleware.Logger)

	apiRouter := chi.NewRouter()
	router.Mount("/v1/", apiRouter)

	app := &internal.Application{
		Config: cfg,
		Router: apiRouter,
	}

	// register routes
	users.Register(app)
	healthcheck.Register(app)

	fmt.Println("Starting server at", app.Config.Port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", app.Config.Port), router)
	if err != nil {
		log.Fatal("Server failed", err)
	}
}
