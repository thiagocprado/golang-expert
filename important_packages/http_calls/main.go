package main

import (
	"io"
	"net/http"
)

func main() {
	// estamos fazendo um requisição do tipo get, porém poderiamos ter outros tipos como post e put
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// fazemos a leitura do corpo da requisição
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	// fazemos a conversão por ser bytes
	println(string(res))

	// fechamos as operações com a requisição
	req.Body.Close()
}
