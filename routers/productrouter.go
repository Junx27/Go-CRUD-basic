package routers

import (
	"aplikasi/controllers/productcontroller"

	"github.com/gorilla/mux"
)

func ProductRouter(r *mux.Router) {
	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product/{id}", productcontroller.Delete).Methods("DELETE")
}
