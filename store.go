package aumo

import "upper.io/db.v3/lib/sqlbuilder"

// Tx represents an SQL transaction
// NOTE: this is an implementation detail (sql)
// How can this be solved?
type Tx interface {
	sqlbuilder.Database
}
