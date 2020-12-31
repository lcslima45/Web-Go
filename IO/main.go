package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/go.jpg", dogPic)
	http.HandleFunc("/hans.jpg", hansPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<img src="/go.jpg">
	<p> <h1>Doutor <em style="color:Pink"> Hans Chucrute </em> </h1> </p>
	<img src="hans.jpg">
	`)

}
func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("go.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	io.Copy(w, f)

}

func hansPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("hans.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	io.Copy(w, f)

}
