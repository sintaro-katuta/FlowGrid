//go:build tools
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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

	// マイグレーションSQLを読み込む
	migrationSQL, err := os.ReadFile("migrations/init.sql")
	if err != nil {
		log.Fatal("マイグレーションファイル読み込みエラー:", err)
	}

	// SQL文を分割して実行
	statements := strings.Split(string(migrationSQL), ";")

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		// CREATE TABLE文のみ実行
		if strings.HasPrefix(strings.ToUpper(stmt), "CREATE TABLE") {
			fmt.Println("実行中:", stmt[:50], "...")
			_, err := db.Exec(stmt)
			if err != nil {
				log.Printf("SQL実行エラー: %v\nSQL: %s\n", err, stmt)
			} else {
				fmt.Println("✓ テーブル作成成功")
			}
		}
	}

	fmt.Println("マイグレーション完了！")

	// テーブル一覧を表示
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name")
	if err != nil {
		log.Fatal("テーブル一覧取得エラー:", err)
	}
	defer rows.Close()

	fmt.Println("\n作成されたテーブル一覧:")
	count := 0
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- %s\n", tableName)
		count++
	}

	if count == 0 {
		fmt.Println("テーブルが作成されていません")
	} else {
		fmt.Printf("合計 %d 個のテーブルが作成されました\n", count)
	}
}
