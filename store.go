package aumo

type Tx interface {
	Commit() error
	Rollback() error
}
