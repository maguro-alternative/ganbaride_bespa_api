package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

    // importされたタイミングでinit関数が実行される。
	_ "github.com/lib/pq"

	// windowsは必須
	"github.com/joho/godotenv"
)

// 外部からも参照するので大文字で始める。
var Db *sql.DB

func init() {
	var err error

	err = godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	//dsn := fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("PGUSER"), os.Getenv("PGPASSOWORD"), os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))
	//dsn := fmt.Sprintf(os.Getenv("PGURL"))

	fmt.Println(dsn)

    // データベースと接続。
	Db, err = sql.Open("postgres", dsn)

	// エラーハンドリング省略
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer Db.Close()

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
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}else{
		fmt.Println("Connection has been established!")
	}
}


func riderinsert(){
	sql := `
	CREATE TABLE IF NOT EXISTS rider(
		number varchar(20) PRIMARY KEY,
		name varchar(50) not null,
		rarity int not null,
		category varchar(50)[] not null,
		attribute varchar(2) not null,
		attack bigint not null,
		defence bigint not null,
		hitpoint bigint not null,
		finishingattack bigint not null,
		skillname varchar(50) not null,
		attackteki bigint not null,
		defenceteki bigint not null,
		hitpointteki bigint not null,
		finishingattackteki bigint not null
	)`

	_, err := Db.Exec(sql)
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}

	sql = `
	INSERT INTO rider(
		number,
		name,
		rarity,
		category,
		attribute,
		attack,
		defence,
		hitpoint,
		finishngattack,
		skillname,
		attackteki,
		defenceteki,
		hitpointteki,
		finishngattackteki
	) VALUES
		(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
		ON CONFLICT DO NOTHING
	`
	_, err = Db.Exec(
		sql,
		err,

	)
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
}