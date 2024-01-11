//go:build db
// +build db

package adapter

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
	id IDGenerator
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func setup() {
	db = sqlx.MustOpen("mysql", os.Getenv("MYSQL_DSN"))
	id = NewXIDGenerator()
	db.MustExec("TRUNCATE TABLE todo")
}
