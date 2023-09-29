package router

import (
	//"github.com/maguro-alternative/ganbaride_bespa_api/service"
	"github.com/maguro-alternative/ganbaride_bespa_api/handlers"
	//"github.com/maguro-alternative/ganbaride_bespa_api/handlers/router"

	"net/http"
	"database/sql"
)


func NewRouter(todoDB *sql.DB) *http.ServeMux {
	//var riderService = service.NewHeroCardService(todoDB)
	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.TeamSuitable)
	return mux
}