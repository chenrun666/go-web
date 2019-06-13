package main

import (
	"net/http"
	"text/template"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func index(w http.ResponseWriter, r *http.Request) {
	args := struct {
		Flag  bool
		Name  string
		Hobby []string
	}{
		Flag:  true,
		Name:  "chenrun",
		Hobby: []string{"123", "456", "789"},
	}

	hobby := []string{
		"football",
		"basketball",
		"pingpang",
	}

	data := make(map[string]interface{})
	data["args"] = &args
	data["hobby"] = &hobby
	data["time"] = time.Now()
	// 创建funcMap
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("index.html").Funcs(funcMap)
	t = template.Must(t.ParseFiles("./index.html"))
	// t, _ = t.ParseFiles("./index.html")
	t.Execute(w, &data)
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
