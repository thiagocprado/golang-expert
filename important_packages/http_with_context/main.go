package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// O `context` é utilizado para controlar o ciclo de vida de operações assíncronas,
	// como cancelamento automático por timeout, cancelamento manual ou passagem de valores entre processos.

	// Criamos um contexto de base, sem nenhuma regra ou restrição.
	ctx := context.Background()

	// A partir do contexto de base, criamos um novo contexto com tempo limite (timeout).
	// Esse novo contexto será automaticamente cancelado após o tempo definido.
	// Também recebemos uma função `cancel()` que pode ser chamada manualmente para encerrar o contexto.
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond) // tempo muito curto para fins de demonstração

	// Garantimos que a função cancel será chamada no final da função main,
	// liberando os recursos associados ao contexto (mesmo que o timeout ocorra antes).
	defer cancel()

	// Criamos uma nova requisição HTTP, associando o contexto com timeout à requisição.
	// Isso significa que, se o tempo limite for atingido, a requisição será automaticamente cancelada.
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		// Se houver erro ao criar a requisição, interrompemos a execução.
		panic(err)
	}

	// Executamos a requisição usando o client padrão do pacote http.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Se a requisição falhar (por timeout ou outro erro), interrompemos a execução.
		panic(err)
	}
	// Garantimos o fechamento do corpo da resposta após sua leitura.
	defer resp.Body.Close()

	// Lemos o conteúdo completo do corpo da resposta.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Imprimimos o conteúdo da resposta no terminal.
	println(string(body))
}
