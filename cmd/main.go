package main

import (
	"helloworld-api/internal/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//localhost:8000/message?firstName=John&lastName=Clark
	router.HandleFunc("/message", routes.PrintMessage).Methods("GET")

	http.ListenAndServe(":8000", router)
}
