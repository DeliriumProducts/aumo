package mysql

import (
	"strings"

	"upper.io/db.v3/lib/sqlbuilder"
)

const ErrDupEntry = 1062

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
