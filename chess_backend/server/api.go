package server

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s APIServer) Run() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/players", s.HandleGetAllPlayers)
	r.Post("/account", s.HandleAddUser)
	r.Delete("/account/{username}", makeHttpFunc(s.HandleDeleteUser))
	log.Println("Listening on:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, r)
}
