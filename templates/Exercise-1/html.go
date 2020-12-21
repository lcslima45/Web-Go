package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//recebe nome da Linha de Comando
	name := os.Args[1]
	nextName := os.Args[2]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1], " ", os.Args[2])
	str := fmt.Sprint(`
		<!DOCTYPE html>
		<html lang=pt-br"> 
		<head>
		<meta charset="UTF-8">
		<title> Hello World!</title> 
		</head> 
		<body>
		<h1 style="color:pink; background-color:yellow">` + name + ` ` + nextName + `</h1>
		</body> 
		</html>
		`)
	file, err := os.Create("index.html")
	if err != nil {
		log.Println("Error making file")
	}
	defer file.Close()
	io.Copy(file, strings.NewReader(str))
}
