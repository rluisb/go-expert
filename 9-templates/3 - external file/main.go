package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursoTemplateMust := template.Must(template.New("template.html").ParseFiles("template.html"))
	err :=  cursoTemplateMust.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 20},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}

}