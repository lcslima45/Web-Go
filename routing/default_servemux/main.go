package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "First Handle Func")
}
func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Second Handle Func")
}

func main() {
	http.HandleFunc("/hotdog", d)
	http.HandleFunc("/hotcat", c)
	http.ListenAndServe(":8080", nil)
}
