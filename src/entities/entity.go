package entities

type Entity interface {
    Create() (err error)
    Get(id int) (err error)
}