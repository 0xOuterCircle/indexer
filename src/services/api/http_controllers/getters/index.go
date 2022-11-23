package getters

import (
    "encoding/json"
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/0xOuterCircle/indexer/src/utils"
    "github.com/0xOuterCircle/indexer/src/utils/notifier"
    "net/http"
    "sync"
)

var once sync.Once
func Index(w http.ResponseWriter, r *http.Request) {
    defer utils.ControllerRecover(w, r) // Handles panic(err)

    once.Do(func() {  // The notify will be sent during healthchek procedure
        go notifier.Deployment()
    })

    resp, err := json.Marshal(config.Global)
    if err != nil { panic(err) }

    // TODO fix CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(resp)
    if err != nil { panic(err) }
}