package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Config recives the Host and Port to set the server.
type Config struct {
	Host string
	Port string
}

// App handles the routing and server
// It has a configuration.
type App struct {
	Router *mux.Router
	server *http.Server
	Config Config
}

// New recives r to set the router in App and a config
// to the server.
func New(c ...Config) *App {

	app := &App{
		Router: mux.NewRouter(),
		server: nil,
		Config: Config{
			Host: "127.0.0.1",
			Port: ":8080",
		},
	}

	if len(c) > 0 {
		app.Config = c[0]
	}

	return app
}

// RegisterRoutes routers send the app router
// to register routes passing the App router to h.
func (app *App) RegisterRoutes(h func(r *mux.Router)) {
	h(app.Router)
}

// Run sets the configuration to
// get set server ready and returns an error.
func (a *App) Run() error {
	a.server = &http.Server{
		Addr:    a.Config.Host + a.Config.Port,
		Handler: a.Router,
	}

	fmt.Println("Server ready!")
	fmt.Printf("Running on http://%s%s\n", a.Config.Host, a.Config.Port)
	return a.server.ListenAndServe()
}
