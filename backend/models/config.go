package models

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

// DBに接続するための構造体
type Server struct {
	DB *sql.DB
}

func DBConnect() (*sql.DB, error) {
	var err error

	// .envファイルを読み込む
	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error open .env file")
		return nil, err
	}

	// 環境変数を変数に格納する
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// 接続プロパティをキャプチャする
	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Net:    "tcp",
		Addr:   dbHost,
		DBName: dbName,
	}

	// データベースを開く
	db, err := sql.Open("mysql", cfg.FormatDSN())
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
