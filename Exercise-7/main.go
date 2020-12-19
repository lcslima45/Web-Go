package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"soma": soma,
}

func soma(a, b int) int {
	return a + b

}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", nil)
	if err != nil {
		log.Println("Error executing", err)
	}

}
