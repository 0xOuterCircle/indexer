package http_controllers

import (
    "encoding/json"
    "log"
    "net/http"
)

func ControllerRecover(w http.ResponseWriter, r *http.Request) {
    // TODO implement concurrency: https://docs.sentry.io/platforms/go/concurrency/
    // TODO Sentry logging
    if recovered := recover(); recovered != nil {
        err := recovered.(error)
        if err.Error() == "no rows in result set" {
            resp, err := json.Marshal(struct {}{})
            if err != nil { ControllerRecover(w, r) }
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusNotFound)
            _, err = w.Write(resp)
        } else {
            log.Println(err)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusInternalServerError)
        }
    }
}