package dao

import (
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/0xOuterCircle/indexer/src/services/db"
)

type Dao struct {
    Id          *int             `json:"id"`
    Parent      *int            `json:"parent"`
    Registry    string          `json:"registry"`
    Governance  string          `json:"governance"`
    Creator     string          `json:"creator"`
}

func (d *Dao) Create() (err error)  {
    _, err = db.Instance().Execute(`insert into dao
        (environment, creator, parent, registry, governance) values ($1, $2, $3, $4, $5)`,
        config.Global.EnvironmentId, d.Creator, d.Parent, d.Registry, d.Governance)
    return err
}

func (d *Dao) Get(id int) (err error) {
    row := db.Instance().QueryRow(`select creator, parent, registry, governance from dao where id = $1`, id)
    err = row.Scan(&d.Creator, &d.Parent, &d.Registry, &d.Governance)
    d.Id = &id
    return err
}