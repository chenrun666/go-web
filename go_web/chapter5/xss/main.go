package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	t := template.Must(template.ParseFiles("./form.html"))
	// t.Execute(w, r.FormValue("comment"))
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./form.html"))
	t.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", form)
	mux.HandleFunc("/process", process)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
