package dao

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/0xOuterCircle/indexer/src/entities"
    "github.com/0xOuterCircle/indexer/src/services/db"
    queries "github.com/0xOuterCircle/indexer/src/sql"
)

type Dao struct {
    Id          *int    `json:"id"`

    Name        string  `json:"name"`
    Description string  `json:"description"`
    Link        string  `json:"link"`
    Image       string  `json:"image"`

    Parent      *int    `json:"parent"`
    Registry    string  `json:"registry"`
    Governance  string  `json:"governance"`
    Creator     string  `json:"creator"`
}

func (d *Dao) Create() (err error)  {
    _, err = db.Instance().Execute(
        queries.Dao.Create,
        config.Global.EnvironmentId, d.Creator, d.Parent, d.Registry, d.Governance,
    )
    return err
}

func (d *Dao) Get(id int) (err error) {
    row := db.Instance().QueryRow(queries.Dao.Get, id)
    err = row.Scan(&d.Name, &d.Description, &d.Link, &d.Image, &d.Creator, &d.Parent, &d.Registry, &d.Governance)
    d.Id = &id
    return err
}

func (d *Dao) Aggregate(offset uint64, limit uint64, args ...interface{}) (res entities.EntityAggregate, err error) {
    res.Offset = offset
    res.Limit = limit
    res.Items = []entities.Entity{}

    params := append(args, config.Global.EnvironmentId, offset, limit)
    rows, err := db.Instance().QueryRows(queries.Dao.Aggregate, params...)
    defer rows.Close()
    if err != nil { return res, err }

    for rows.Next() {
        entity := &Dao{}
        err = rows.Scan(
            &res.Total,
            &entity.Name,
            &entity.Description,
            &entity.Link,
            &entity.Image,
            &entity.Creator,
            &entity.Parent,
            &entity.Registry,
            &entity.Governance,
        )
        res.Items = append(res.Items, entity)
    }

    // Handle requests with offset exceeds limit
    if len(res.Items) == 0 {
        params = append(args, config.Global.EnvironmentId)
        row := db.Instance().QueryRow(queries.Dao.Total, config.Global.EnvironmentId)
        err = row.Scan(&res.Total)
    }

    return res, err
}