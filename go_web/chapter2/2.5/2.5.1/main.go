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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
