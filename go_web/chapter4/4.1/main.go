package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// 观察请求头信息
func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Accept-Encoding"]
	fmt.Fprintln(w, h)

}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Printf("接收到请求来的数据%v\n", string(body))
	fmt.Fprintln(w, string(body))
}

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// r.ParseMultipartForm(1024)

	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
	fmt.Fprintln(w, r.FormValue("post"))
	fmt.Fprintln(w, r.MultipartForm) // 获取文件内容
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./form.html")
	t.Execute(w, "")
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/body", body)
	http.HandleFunc("/process", process)
	http.HandleFunc("/index", index)
	server.ListenAndServe()
}
