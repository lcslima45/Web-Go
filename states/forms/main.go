package main

import (
	"html/template"
	"log"
	"net/http"
)

type Form struct {
	FirstName string
	LastName  string
	Subscribe bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", form)
	http.ListenAndServe(":8080", nil)
}
func form(w http.ResponseWriter, r *http.Request) {
	FirstName := r.FormValue("FirstName")
	LastName := r.FormValue("LastName")
	Subscribe := r.FormValue("Subscribe") == "on"
	err := tpl.ExecuteTemplate(w, "index.html", Form{FirstName, LastName, Subscribe})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}
