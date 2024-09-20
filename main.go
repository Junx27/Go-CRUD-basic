package main

import (
	model "aplikasi/models"
	"aplikasi/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	model.ConnectDatabase()
	r := mux.NewRouter()

	routers.ProductRouter(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
