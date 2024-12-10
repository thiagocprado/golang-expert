package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// o defer basicamente atrasa a finalização das operações, ou seja, essa linha será executada após rodar todos os outros processos
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))

	defer fmt.Println("Primeira linha")
	fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")
}
