package api

import (
    "github.com/0xOuterCircle/indexer/src/services/api/http_controllers"
    "log"
    "net/http"
)

func Run()  {

    server := &http.Server{
        Addr:    ":8080",
        Handler: http_controllers.Routing(),
    }
    log.Println("API: running on :8080")

    err := server.ListenAndServe()
    if err != nil {
        log.Println("Error while running the router:", err)
    }

}