package main

import (
	"fmt"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello Go!!!")

}

// 多路复用器
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("../../../../../../src/"))
	mux.Handle("/source/", http.StripPrefix("/source/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
