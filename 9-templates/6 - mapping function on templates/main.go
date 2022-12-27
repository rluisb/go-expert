package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	//Adicionamos funções para que possam ser usadas no template
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
}
