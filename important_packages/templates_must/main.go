package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

func main() {
	course := Course{"Go", 40}

	// podemos simplificar o código de template como estamos fazendo abaixo
	t := template.Must(template.New("course-template").Parse("Course: {{.Name}} - Duration: {{.Hours}} hours"))

	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
