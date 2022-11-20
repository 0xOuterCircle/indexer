package http_controllers

import (
    "encoding/json"
    "github.com/0xOuterCircle/indexer/src/config"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
    defer ControllerRecover(w, r) // Handles panic(err)

    resp, err := json.Marshal(config.Global)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err != nil { panic(err) }
    _, err = w.Write(resp)
    if err != nil { panic(err) }
}