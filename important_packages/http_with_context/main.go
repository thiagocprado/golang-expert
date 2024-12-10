package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// context é usado para controlar prazos, cancelamentos e metadados em operações que podem ser interrompidas,
	// como chamadas de API, goroutines ou processamento de dados

	// contexto vazio, sem regras
	ctx := context.Background()

	// cancela o contexto caso de o timeout
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond)
	// cancela o contexto
	defer cancel()

	// definimos no contexto que temos um tempo para completar a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
