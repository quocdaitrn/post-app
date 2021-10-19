package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var port string = ":9500"
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	router.HandleFunc("/posts", listPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}