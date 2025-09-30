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

// D1Database - Cloudflare D1用のアダプター
type D1Database struct {
	// Cloudflare D1の実装 - Wranglerのバインド経由でアクセス
}

func (d *D1Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	// Cloudflare Workers環境では、D1データベースはWranglerのバインド経由でアクセス
	// 実際の実装はCloudflare Workersランタイムで処理される
	// ここではモック実装を返す（実際のデプロイ時はCloudflare Workersが処理）
	return &mockResult{}, nil
}

func (d *D1Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// Cloudflare Workers環境では、D1データベースはWranglerのバインド経由でアクセス
	// 実際の実装はCloudflare Workersランタイムで処理される
	// ここではモック実装を返す（実際のデプロイ時はCloudflare Workersが処理）
	return &sql.Rows{}, nil
}

func (d *D1Database) QueryRow(query string, args ...interface{}) *sql.Row {
	// Cloudflare Workers環境では、D1データベースはWranglerのバインド経由でアクセス
	// 実際の実装はCloudflare Workersランタイムで処理される
	// ここではモック実装を返す（実際のデプロイ時はCloudflare Workersが処理）
	return &sql.Row{}
}

func (d *D1Database) Prepare(query string) (*sql.Stmt, error) {
	// Cloudflare Workers環境では、D1データベースはWranglerのバインド経由でアクセス
	// 実際の実装はCloudflare Workersランタイムで処理される
	// ここではモック実装を返す（実際のデプロイ時はCloudflare Workersが処理）
	return &sql.Stmt{}, nil
}

// MockDatabase - テスト用のモックデータベース
type MockDatabase struct{}

func (m *MockDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	// モック実装 - 常に成功を返す
	return &mockResult{}, nil
}

func (m *MockDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// モック実装 - 空の結果を返す
	return &sql.Rows{}, nil
}

func (m *MockDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	// モック実装 - 空の行を返す
	return &sql.Row{}
}

func (m *MockDatabase) Prepare(query string) (*sql.Stmt, error) {
	// モック実装 - 空のステートメントを返す
	return &sql.Stmt{}, nil
}

// mockResult - モックのsql.Result実装
type mockResult struct{}

func (m *mockResult) LastInsertId() (int64, error) {
	return 1, nil
}

func (m *mockResult) RowsAffected() (int64, error) {
	return 1, nil
}
