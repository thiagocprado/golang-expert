package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		// podemos renderizar o html direto no nosso navegador também
		if err := t.Execute(w, Courses{
			{"Go", 40},
			{"Node JS", 20},
			{"Java", 80},
		}); err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
