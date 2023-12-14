//go:generate mockgen -source=adapter.go -destination=mock_adapter.go -package=adapter
package adapter

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Adapters struct {
	Pet *Pet
}

func New(db DBAdapter, id IDGenerator) Adapters {
	return Adapters{
		Pet: NewPet(db, id),
	}
}

type DBAdapter interface {
	DriverName() string
	MapperFunc(mf func(string) string)
	Rebind(query string) string
	Unsafe() *sqlx.DB
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	MustBegin() *sqlx.Tx
	Beginx() (*sqlx.Tx, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	MustExec(query string, args ...interface{}) sql.Result
	Preparex(query string) (*sqlx.Stmt, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
}
type IDGenerator interface {
	Generate() string
}
