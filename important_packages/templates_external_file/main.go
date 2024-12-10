package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	// poderiamos passar diversos arquivos de uma vez para o parse files
	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	if err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Node JS", 20},
		{"Java", 80},
	}); err != nil {
		panic(err)
	}
}
