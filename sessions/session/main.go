package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//estrutura de tabela para o usuário
type user struct {
	UserName string
	First    string
	Last     string
}

//rotina para o template
var tpl *template.Template

//mapa representando um banco de dados do usuário
var dbUsers = map[string]user{}

//mapa representando um banco de dados de sessões
var dbSessions = map[string]string{}

func init() {
	//passa pasta de templates
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//função index
func index(w http.ResponseWriter, req *http.Request) {
	//observa cookie
	c, err := req.Cookie("session")
	if err != nil {
		//criar um id único
		id, _ := uuid.NewV4()
		//novo cookie
		c = &http.Cookie{
			Name:  "session",
			Value: id.String(),
		}
		//inscrever cookie na resposta
		http.SetCookie(w, c)
	}

	var u user

	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	//inserindo usuário do formulário em dbUsers

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(w, "index.html", u)

}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}
