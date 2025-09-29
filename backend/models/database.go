package models

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

// DBに接続するための構造体
type Server struct {
	DB *sql.DB
}

// DBConnect は環境に応じて適切なデータベース接続を返します
func DBConnect() (*sql.DB, error) {
	// Cloudflare Workers環境かどうかを判定
	if isCloudflareEnvironment() {
		return connectToD1()
	}
	
	// 通常のSQLite接続
	return connectToSQLite()
}

// isCloudflareEnvironment はCloudflare Workers環境かどうかを判定します
func isCloudflareEnvironment() bool {
	// Cloudflare Workers環境では特定の環境変数が設定されている
	return os.Getenv("CF_PAGES") == "1" || 
	       os.Getenv("CLOUDFLARE_WORKERS") != "" ||
	       runtime.GOOS == "js" // WASM環境
}

// connectToD1 はCloudflare D1データベースに接続します
func connectToD1() (*sql.DB, error) {
	// Cloudflare D1はWrangler経由でバインドされるため、
	// 実際の接続はCloudflare Workersランタイムが処理します
	fmt.Println("Cloudflare D1接続（バインド経由）")
	
	// ダミーの接続を返す（実際の接続はCloudflare Workersランタイムが処理）
	// 本番環境では適切なD1ドライバーを使用する必要があります
	return nil, fmt.Errorf("D1接続はCloudflare Workers環境でのみ利用可能です")
}

// connectToSQLite はSQLiteデータベースに接続します
func connectToSQLite() (*sql.DB, error) {
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
