package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// DBに接続するための構造体
type Server struct {
	DB *sql.DB
}

func DBConnect() (*sql.DB, error) {
	var err error

	// RailwayのDATABASE_URL環境変数を優先的に使用
	connStr := os.Getenv("DATABASE_URL")
	
	// DATABASE_URLが設定されていない場合は個別の環境変数を使用
	if connStr == "" {
		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASS")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")

		// 環境変数が設定されているか確認
		if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" {
			fmt.Println("Missing required environment variables: DATABASE_URL or DB_USER, DB_PASS, DB_HOST, DB_NAME")
			return nil, fmt.Errorf("missing required database environment variables")
		}

		// PostgreSQL接続文字列を作成
		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", dbUser, dbPass, dbHost, dbName)
	}

	// データベースを開く
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("DB open Error")
		return nil, err
	}

	// 接続が有効であるか確認する
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("pingErr")
		return nil, pingErr
	}

	fmt.Println("接続成功！！")
	return db, nil
}
