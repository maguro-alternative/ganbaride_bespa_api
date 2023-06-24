package main

import (
	"github.com/maguro-alternative/ganbaride_bespa_api/model"
	"github.com/maguro-alternative/ganbaride_bespa_api/handlers"

	"net/http"
	"fmt"
)

func migrate() {
	sql := `INSERT INTO todos(id, name, status) VALUES('00000000000000000000000000','買い物', '作業中'),('00000000000000000000000001','洗濯', '作業中'),('00000000000000000000000002','皿洗い', '完了');`

	_, err := model.Db.Exec(sql)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migration is success!")
}

func main() {
    // 省略
	port := "8080"
	http.HandleFunc("/fetch-todos", handlers.TeamSuitable)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}