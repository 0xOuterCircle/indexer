package http_controllers

import (
    "github.com/gorilla/mux"
)

func Routing() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", Index).Methods("GET")
    router.HandleFunc("/dao/{id}", DaoGet).Methods("GET")

    return router
}