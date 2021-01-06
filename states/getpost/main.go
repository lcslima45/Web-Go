package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", foo)
	http.HandleFunc("/post", post)
	http.HandleFunc("/get", get)
	http.Handle("/foo/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	fmt.Fprintln(w, "<b> Do my search: "+v+"</b>")
}

func post(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<form method="POST"> 
	<input type="text" name="q"> 
	<input type="submit"> 
	</form><br>`+v)
}
func get(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<form method="get"> 
	<input type="text" name="q"> 
	<input type="submit"> 
	</form><br>`+v)
}
