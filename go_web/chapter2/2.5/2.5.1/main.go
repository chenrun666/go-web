package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// Must 捕捉模板的异常
	t, _ := template.ParseFiles("./templates/tmpl.html")
	// Excute 只会执行第一个模版
	t.Execute(w, "Hello Go !!!")
}

func multiTemplate(w http.ResponseWriter, r *http.Request) {
	var files []string
	files = append(files, "./templates/layout.html", "./templates/tmpl.html")

	t, _ := template.ParseFiles(files...)
	// ExecuteTemplate 执行的时候指定模版的时候只需要执行名字。不需要加路径
	t.ExecuteTemplate(w, "tmpl.html", "Hello Go !!!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/tmpl", multiTemplate)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
