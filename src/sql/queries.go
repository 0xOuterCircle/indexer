package queries

type entityQuerySet struct {
    Create      string
    Get         string
    Aggregate   string
    Total       string
}

var Dao = entityQuerySet{
    Create: `insert into dao (environment, creator, parent, registry, governance) values ($1, $2, $3, $4, $5)`,
    Get: `select name, description, link, image, creator, parent, registry, governance from dao where id = $1`,
    Aggregate: `select count(*) over() as total, name, description, link, image, creator, parent, registry, governance from dao where environment = $1 offset $2 limit $3`,
    Total: `select count(*) as total from dao where environment = $1`,
}
