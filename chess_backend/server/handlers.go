package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func (api APIServer) HandleGetAllPlayers(w http.ResponseWriter, r *http.Request) {
	rows, err := api.store.getAllAccounts()
	if err != nil {
		log.Fatal(err)
	}
	WriteJSON(w, http.StatusOK, rows)

}

func (api APIServer) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	createReq := new(CreateAccountReq)
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	user := NewPlayer(createReq.username, createReq.passwordhash, createReq.email)
	err := api.store.createAccount(user)
	if err != nil {
		log.Fatal(err)
		WriteJSON(w, http.StatusInternalServerError, err)
		return
	}
	WriteJSON(w, http.StatusOK, user)
}

func (api APIServer) HandleDeleteUser(w http.ResponseWriter, r *http.Request) error {

	username := chi.URLParam(r, "username")
	if err := api.store.deleteAccount(username); err != nil {
		return err
	}
	WriteJSON(w, http.StatusOK, "Deleted")
	return nil
}
