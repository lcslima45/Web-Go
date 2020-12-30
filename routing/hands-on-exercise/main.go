package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "My name is Lucas")
}
func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "This is the page of my dog")
}
func welcome(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Welcome my page")
}
func something(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Println("Error parsing file", err)
	}
	err = tpl.ExecuteTemplate(res, "template.html", "Luquinhas")
	if err != nil {
		log.Println("Error", err)
	}
}

func main() {
	http.Handle("/", http.HandlerFunc(welcome))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))
	http.Handle("/willkommen", http.HandlerFunc(something))
	http.ListenAndServe(":8080", nil)
}
