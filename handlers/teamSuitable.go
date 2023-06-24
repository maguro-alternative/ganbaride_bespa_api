package handlers

import (
    "fmt"
    "net/http"
)

func TeamSuitable(w http.ResponseWriter, r *http.Request){

	// HTTPメソッドをチェック（POSTのみ許可）
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed) // 405
        w.Write([]byte("POSTだけだよー"))
        return
    }
	
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello World")
}