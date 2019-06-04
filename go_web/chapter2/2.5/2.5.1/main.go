package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	var files []string
	for _, file := range fn {
		files = append(files, fmt.Sprintf("./templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func index(w http.ResponseWriter, r *http.Request) {
	// threads, err := data.Threads()
	// if err == nil {
	// _, err = session(w, r)
	// 	if err != nil {
	// 		generateHTML(w, threads, "layout", "public.navbar", "index")
	// 	} else {
	// 		generateHTML(w, threads, "layout", "public.navbar", "index")
	// 	}
	// }
	generateHTML(w, "", "layout", "navbar", "content")
}

// 增加模版展示
func templateExample(w http.ResponseWriter, r *http.Request) {
	var files []string
	files = append(files, "./templates/example.html", "./templates/layout.html")

	t := template.Must(template.ParseFiles("./templates/layout.html"))
	t.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/templates", templateExample)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
