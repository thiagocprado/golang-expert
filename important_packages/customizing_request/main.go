package main

import (
	"io"
	"net/http"
)

func main() {
	// Criamos um client HTTP que será responsável por enviar a requisição.
	// O http.Client pode ser reutilizado para múltiplas requisições.
	c := http.Client{}

	// Criamos uma nova requisição HTTP do tipo GET para o site do Google.
	// O terceiro parâmetro é o corpo da requisição (usado em POST, PUT, etc.),
	// mas como estamos usando GET, passamos 'nil'.
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		// Se houver erro ao criar a requisição, encerramos o programa com panic.
		panic(err)
	}

	// Adicionamos um cabeçalho (header) à requisição.
	// Aqui estamos dizendo que queremos receber a resposta no formato JSON (se possível).
	req.Header.Set("Accept", "application/json")

	// Enviamos a requisição usando o client criado anteriormente.
	resp, err := c.Do(req)
	if err != nil {
		// Se der erro ao fazer a requisição (ex: conexão falhou), encerramos o programa.
		panic(err)
	}
	// Garantimos que o corpo da resposta será fechado após o uso, evitando vazamentos de recurso.
	defer resp.Body.Close()

	// Lemos todo o conteúdo do corpo da resposta.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Imprimimos o conteúdo da resposta como uma string no terminal.
	println(string(body))
}
