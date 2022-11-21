package getters

import (
    "encoding/json"
    "github.com/0xOuterCircle/indexer/src/entities/dao"
    "github.com/0xOuterCircle/indexer/src/utils"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func Dao(w http.ResponseWriter, r *http.Request) {
    defer utils.ControllerRecover(w, r) // Handles panic(err)

    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil { panic(err) }

    entity := dao.Dao{}
    err = entity.Get(id)
    if err != nil { panic(err) }
    resp, err := json.Marshal(entity)
    if err != nil { panic(err) }

    // TODO fix CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(resp)
    if err != nil { panic(err) }
}