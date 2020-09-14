package dbutils

import "database/sql"

const (
	// ChartTable table containing Charts info
	ChartTable = "charts"
	// RepositoryTable table containing repositories sync info
	RepositoryTable = "repos"
	// ChartFilesTable table containing files related to other charts
	ChartFilesTable = "files"
	// EnvvarPostgresTests enables tests that run against a local postgres
	EnvvarPostgresTests = "ENABLE_PG_INTEGRATION_TESTS"
)

type PostgresDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Begin() (*sql.Tx, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Close() error
	Exec(query string, args ...interface{}) (sql.Result, error)
}
