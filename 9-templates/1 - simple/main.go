package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	cursoTemplate := template.New("CursoTemplate")
	//Em go, quando usamos template, utilizamos . antes do nome da variável do struct que iremos passar por parametro
	cursoTemplate, _ = cursoTemplate.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err :=  cursoTemplate.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}