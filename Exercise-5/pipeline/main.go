package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func double(x int) int {
	return 2 * x
}
func power(x int) int {
	return int(math.Pow(float64(x), 15.2))
}
func root(x int) int {
	return int(math.Sqrt(float64(x)))
}

var fm = template.FuncMap{
	"double": double,
	"power":  power,
	"root":   root,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln("Error executing", err)
	}
}
