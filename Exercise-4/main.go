package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Enterprise struct {
	Name      string
	Employers []string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	templates, err := ioutil.ReadDir("templates/")
	stored := []interface{}{prophets, values, prophetsMap, prophetsMap, enterpriseGo}
	if err != nil {
		log.Fatalln("Error reading file", err)
	}
	for index, temp := range templates {
		function := stored[index]
		err = tpl.ExecuteTemplate(os.Stdout, temp.Name(), function.(func())
		if err != nil {
			log.Println("Error executing", err)
		}
	}
}

func prophets() []string {
	return []string{
		"Jesus",
		"Maomé",
		"Saddhur Sudhar Sink",
		"Zoroastro",
		"Seiyia de Pegasus",
	}
}

func values() []int {
	return []int{
		1,
		2,
		3,
		55,
		17,
	}
}

func prophetsMap() map[string]string {

	a := map[string]string{
		"1": "Jesus",
		"2": "Maomé",
		"3": "Eu",
	}
	return a
}

func enterpriseGo() []Enterprise {
	enterprise := make([]Enterprise, 2)
	boing := Enterprise{
		"Boing",
		[]string{
			"Joao",
			"Paulo",
		},
	}
	gol := Enterprise{
		"GOL",
		[]string{
			"Barnabé",
			"Simão",
		},
	}
	enterprise = append(enterprise, boing)
	enterprise = append(enterprise, gol)
	return enterprise
}
