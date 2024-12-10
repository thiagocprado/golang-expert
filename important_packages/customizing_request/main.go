package main

import (
	"io"
	"net/http"
)

func main() {
	// o client é uma coisa
	c := http.Client{}

	// a requisição é outra
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	// aqui temos uma maior customização da request
	// temos uma customização também de headers
	req.Header.Set("Accept", "application/json")

	// pedimos ao client para fazer a request
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
