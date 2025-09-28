package db

import (
	"database/sql"
)

// Databaseインターフェース - ローカルSQLiteとCloudflare D1の両方に対応
type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

// SQLDatabase - 標準のsql.DBをラップ
type SQLDatabase struct {
	DB *sql.DB
}

func (s *SQLDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SQLDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SQLDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.DB.QueryRow(query, args...)
}

func (s *SQLDatabase) Prepare(query string) (*sql.Stmt, error) {
	return s.DB.Prepare(query)
}

// D1Database - Cloudflare D1用のアダプター（将来的な実装用）
type D1Database struct {
	// Cloudflare D1の実装
}

func (d *D1Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	// D1のExec実装
	return nil, nil
}

func (d *D1Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// D1のQuery実装
	return nil, nil
}

func (d *D1Database) QueryRow(query string, args ...interface{}) *sql.Row {
	// D1のQueryRow実装
	return nil
}

func (d *D1Database) Prepare(query string) (*sql.Stmt, error) {
	// D1のPrepare実装
	return nil, nil
}
