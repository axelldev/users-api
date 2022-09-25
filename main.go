package main

import (
	"log"
	"net/http"
	"os"

	"github.com/axelldev/users-api/app"
	"github.com/axelldev/users-api/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("cannot load .env file")
	}

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	app := app.New(app.Config{
		Host: HOST,
		Port: PORT,
	})

	RegisterRoutes(app)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func RegisterRoutes(a *app.App) {
	a.RegisterRoutes(func(r *mux.Router) {
		api := r.PathPrefix("/api").Subrouter()
		v1 := api.PathPrefix("/v1").Subrouter()

		// Users endpoints
		u := v1.PathPrefix("/users").Subrouter()
		u.HandleFunc("", handlers.GetUsers).Methods(http.MethodGet)
		u.HandleFunc("/{id}", handlers.GetUser)
	})
}
