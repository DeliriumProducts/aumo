package mysql

import (
	"strings"

	"upper.io/db.v3/lib/sqlbuilder"
)

// ErrDupEntry is an error code for duplicate entry
const ErrDupEntry = 1062

// ErrBadRef is an error code for bad reference on foreign key
const ErrBadRef = 1452

// ExecSchema takes a schema and splits it by `;` and executes it
func ExecSchema(db sqlbuilder.Database) error {
	l := strings.SplitAfter(Schema, ";")

	for _, sp := range l[:len(l)-1] {
		_, err := db.Exec(sp)
		if err != nil {
			return err
		}
	}

	return nil
}
