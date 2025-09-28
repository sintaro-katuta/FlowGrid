package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// DBに接続するための構造体
type Server struct {
	DB *sql.DB
}

func DBConnect() (*sql.DB, error) {
	var err error

	// 環境変数を変数に格納する
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./flowgrid.db" // デフォルトパス
	}

	// SQLiteデータベースを開く
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("DB open Error:", err)
		return nil, err
	}

	// 外部キー制約を有効にする
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		fmt.Println("Error enabling foreign keys:", err)
		return nil, err
	}

	// 接続が有効であるか確認する
	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("pingErr:", pingErr)
		return nil, pingErr
	}

	fmt.Println("SQLite接続成功！！")
	return db, nil
}
