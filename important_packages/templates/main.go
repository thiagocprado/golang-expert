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
	tmp := template.New("course-template")

	// temos que utilizar o . nas variáveis que queremos passar como parâmetro
	tmp, _ = tmp.Parse("Course: {{.Name}} - Duration: {{.Hours}} hours")

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
