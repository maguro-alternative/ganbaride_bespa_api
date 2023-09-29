package main

import (
	"github.com/maguro-alternative/ganbaride_bespa_api/db"
	"github.com/maguro-alternative/ganbaride_bespa_api/handlers/router"

	"net/http"
	"os"
	"log"
	"time"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatalln("main: failed to exit successfully, err =", err)
	}
}

func realMain() error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	// TODO: サーバーをlistenする
	// NOTE: ポート番号は上記のport変数を使用すること
	// NOTE: エラーが発生した場合はlog.Fatalでログを出力すること
	log.Fatal(http.ListenAndServe(port, mux))

	return nil
}