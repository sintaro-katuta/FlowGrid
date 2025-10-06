//go:build tools
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	// 環境変数を変数に格納する
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// PostgreSQL接続文字列を作成
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=require", dbUser, dbPass, dbHost, dbName)
	fmt.Println("接続文字列:", connStr)

	// データベースを開く
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB open Error:", err)
	}
	defer db.Close()

	// 接続が有効であるか確認する
	err = db.Ping()
	if err != nil {
		log.Fatal("pingErr:", err)
	}

	fmt.Println("接続成功！！")

	// テーブル一覧を表示
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		log.Fatal("テーブル一覧取得エラー:", err)
	}
	defer rows.Close()

	fmt.Println("テーブル一覧:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		fmt.Println("-", tableName)
	}
}
