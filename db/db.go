package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	// ローカルファイルを埋め込む
	_ "embed"

    // importされたタイミングでinit関数が実行される。
	_ "github.com/lib/pq"

	// windowsは必須
	"github.com/joho/godotenv"
)

// schema.sqlの内容を文字で読み取る。
var schema string

func NewDB(path string) (*sql.DB, error) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	//dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSOWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
	)
	//dsn := fmt.Sprintf(os.Getenv("PGURL"))

	fmt.Println(dsn)

    // データベースと接続。
	db, err := sql.Open("postgres", dsn)

	// エラーハンドリング省略
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()

    // 疎通確認を行う。
	err = db.Ping()
	// エラーハンドリング
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	_, err = db.Exec(schema)

	// エラーハンドリング
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}else{
		fmt.Println("Connection has been established!")
	}

	return db, nil
}
