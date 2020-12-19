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

type values struct {
	Value1 int
	Value2 int
}

func soma(a, b int) int {
	return a + b

}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {
	value := values{
		Value1: 1,
		Value2: 2,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", value)
	if err != nil {
		log.Println("Error executing", err)
	}

}
