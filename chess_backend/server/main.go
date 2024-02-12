package server

import (
	"log"

	_ "github.com/lib/pq"
)

func StartServer() {
	store, err := NewPSQLDB()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	apiServer := NewAPIServer(":8080", store)
	apiServer.Run()
}
