package http_controllers

import (
    "github.com/0xOuterCircle/indexer/src/services/api/http_controllers/getters"
    "github.com/0xOuterCircle/indexer/src/services/api/http_controllers/getters/aggregate"
    "github.com/gorilla/mux"
)

func Routing() *mux.Router {
    router := mux.NewRouter()

    // Getters
    router.HandleFunc("/", getters.Index).Methods("GET")
    router.HandleFunc("/dao/{id}", getters.Dao).Methods("GET")

    // Getters aggregate
    router.HandleFunc("/dao", aggregate.Dao).Methods("GET")

    return router
}