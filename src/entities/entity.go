package entities

type EntityAggregate struct {
    Total  uint64      `json:"total"`
    Offset uint64      `json:"offset"`
    Limit  uint64      `json:"limit"`
    Items  []Entity    `json:"items"`
}

type Entity interface{
    Create() (err error)
    Get(id int) (err error)
    Aggregate(offset uint64, limit uint64, args ...interface{}) (res EntityAggregate, err error)
}
