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
    
    len := r.ContentLength
    body := make([]byte, len) // Content-Length と同じサイズの byte 配列を用意
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))

	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello World")
}