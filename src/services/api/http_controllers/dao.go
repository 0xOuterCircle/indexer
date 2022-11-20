package http_controllers

import (
    "encoding/json"
    "github.com/0xOuterCircle/indexer/src/entities/dao"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

func DaoGet(w http.ResponseWriter, r *http.Request) {
    defer ControllerRecover(w, r) // Handles panic(err)

    id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil { panic(err) }

    entity := dao.Dao{}
    err = entity.Get(id)
    if err != nil { panic(err) }
    resp, err := json.Marshal(entity)
    if err != nil { panic(err) }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err != nil { panic(err) }
    _, err = w.Write(resp)
    if err != nil { panic(err) }
}