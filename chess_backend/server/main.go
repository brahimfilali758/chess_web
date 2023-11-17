package server

import (
	"fmt"
	"net/http"
)


func StartServer() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	fmt.Printf("Listening ..")
	err := http.ListenAndServe(":3333", nil)
	if (err != nil ){
		fmt.Printf("Error")
	}
}


