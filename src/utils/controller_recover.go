package utils

import (
    "log"
    "net/http"
)

func ControllerRecover(w http.ResponseWriter, r *http.Request) {
    // TODO implement concurrency: https://docs.sentry.io/platforms/go/concurrency/
    // TODO Sentry logging
    if recovered := recover(); recovered != nil {
        err := recovered.(error)
        if err.Error() == "no rows in result set" {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusNotFound)
        } else if err.Error() == "400" {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusBadRequest)
        } else {
            log.Println(err)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
        }
    }
}