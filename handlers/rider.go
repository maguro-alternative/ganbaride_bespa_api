package handlers

import (
	"fmt"
	"net/http"

	"github.com/maguro-alternative/ganbaride_bespa_api/service"
)

type RiderHandler struct {
	svc *service.HeroCardService
}

func NewRiderHandler(svc *service.HeroCardService) *RiderHandler {
	return &RiderHandler{
		svc: svc,
	}
}

func (h *RiderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){

	// HTTPメソッドをチェック（POSTのみ許可）
    if r.Method == http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed) // 405
        w.Write([]byte("POSTだけだよー"))
        return
    }

    len := r.ContentLength
    body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))

	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello World")
}