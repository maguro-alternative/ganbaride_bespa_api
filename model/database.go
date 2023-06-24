package model

import (
	"database/sql"
	"fmt"
	"os"

    // importされたタイミングでinit関数が実行される。
	_ "github.com/go-sql-driver/mysql"
)

// 外部からも参照するので大文字で始める。
var Db *sql.DB

func init() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

    // データベースと接続。
	Db, err = sql.Open("mysql", dsn)

	// エラーハンドリング省略

    // 疎通確認を行う。
	err = Db.Ping()

	// エラーハンドリング省略

    // database.goがimportされたらinit関数が走り、このSQLが実行される。
	sql := `CREATE TABLE IF NOT EXISTS todos(
			id varchar(26) not null,
			name varchar(100) not null,
			status varchar(100) not null
		)`

	_, err = Db.Exec(sql)

	// エラーハンドリング省略

	fmt.Println(err)

	fmt.Println("Connection has been established!")
}
