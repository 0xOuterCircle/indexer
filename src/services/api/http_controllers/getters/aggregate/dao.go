package aggregate

import (
    "encoding/json"
    "errors"
    "github.com/0xOuterCircle/indexer/src/entities/dao"
    "github.com/0xOuterCircle/indexer/src/utils"
    "net/http"
    "strconv"
)

func Dao(w http.ResponseWriter, r *http.Request) {
    defer utils.ControllerRecover(w, r) // Handles panic(err)

    offset, err := strconv.ParseUint(r.FormValue("offset"), 10, 64)
    limit, err := strconv.ParseUint(r.FormValue("limit"), 10, 64)
    if err != nil { panic(errors.New("400")) }

    entity := dao.Dao{}
    agg, err := entity.Aggregate(offset, limit)
    if err != nil { panic(err) }
    resp, err := json.Marshal(agg)
    if err != nil { panic(err) }

    // TODO fix CORS
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(resp)
    if err != nil { panic(err) }
}