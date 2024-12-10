package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	// temos que ter um arquivo base, nesse caso, é o nosso content
	t := template.New("content.html")

	// podemos colocar um mapa de funções para serem parseados junto aos valores passados
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})

	// os três pontinhos é porque a função é variádica, ou seja, cada item do slice é um parâmetro
	t = template.Must(t.ParseFiles(templates...))
	if err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Node JS", 20},
		{"Java", 80},
	}); err != nil {
		panic(err)
	}
}
