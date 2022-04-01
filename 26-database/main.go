package main

import (
	"log"
	"net/http"

	"database-golang/src/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// create router handlers
	controllers.CreateHandlers(router)

	log.Println("listening on port :5005")
	log.Fatal(http.ListenAndServe(":5005", router))
}
