package internal

import "github.com/go-chi/chi/v5"

// for app version
const Version = "1.0.0"

// config to hold all configurations
// port is port, env is wheteher developmnent staging or production
type Config struct {
	Port int
	Env  string
}

// to hold all app dependencies handlers, helpers and middlewares
type Application struct {
	Config    Config
	Router *chi.Mux
}
