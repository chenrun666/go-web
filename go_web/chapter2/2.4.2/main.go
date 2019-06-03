package main

import "net/http"

// 服务静态文件
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
