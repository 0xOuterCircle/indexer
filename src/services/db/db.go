package db

// Vulnerable to Race condition,
// looks like we don't need a fix cuz we have no setters

// TODO SentryRecover

import (
    "context"
    "github.com/0xOuterCircle/indexer/src/config"
    "github.com/jackc/pgconn"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
    "log"
    "sync"
)

var once sync.Once

type db struct {
    pool       *pgxpool.Pool
    connString string
}

var dbInstance *db = nil

type DbInterface interface {
    Connect()
    QueryRow(query string, args ...interface{}) pgx.Row
    QueryRows(query string, args ...interface{}) (rows pgx.Rows, err error)
    Execute(query string, args ...interface{}) (tag pgconn.CommandTag, err error)
}

func Instance() DbInterface {
    once.Do(func() {
        dbInstance = &db{connString: config.Global.DbConnStr}
    })
    return dbInstance
}

func (db *db) Connect() {
    log.Println("DB: connecting to Indexer Database")
    connection, err := pgxpool.Connect(
        context.Background(),
        db.connString,
    )
    if err != nil {
        panic(err)
    } else {
        db.pool = connection
        log.Println("DB: connected to Indexer Database")
    }
}

func (db *db) QueryRows(query string, args ...interface{}) (rows pgx.Rows, err error) {
    // ToDo research reflect and implement writing to structs here
//    defer SentryRecover()
    rows, err = db.pool.Query(context.Background(), query, args...)
    return
}

func (db *db) QueryRow(query string, args ...interface{}) pgx.Row {
//    defer SentryRecover()
    return db.pool.QueryRow(context.Background(), query, args...)
}

func (db *db) Execute(query string, args ...interface{}) (tag pgconn.CommandTag, err error) {
//    defer SentryRecover()
    tag, err = db.pool.Exec(context.Background(), query, args...)
    return
}