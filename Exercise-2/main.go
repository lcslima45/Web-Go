package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}
func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error parsing files")
	}
	err = tpl.ExecuteTemplate(os.Stdout, "two.moe", nil)
	if err != nil {
		log.Fatalln("Error template executing")
	}
	err = tpl.ExecuteTemplate(os.Stdout, "Third.moe", nil)
	if err != nil {
		log.Fatalln("Error template executing")
	}
}
