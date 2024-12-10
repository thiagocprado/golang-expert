package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	// buffer pode ser entendido como uma área temporária de memória usada para armazenar dados antes que eles sejam processados ou enviados para outro destino
	jsonVar := bytes.NewBuffer([]byte(`{"name":"thiago"}`))

	// o método post precisa de parâmetros a mais além da url
	resp, err := c.Post("http://google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.CopyBuffer(os.Stdout, resp.Body, nil)
}

// Buffers são especialmente úteis quando você está lidando com fluxos de dados ou I/O (entrada/saída),
// permitindo manipulação eficiente e reduzindo o número de operações diretas com os dispositivos ou streams subjacentes.

// Armazenamento Dinâmico: O tamanho do buffer cresce conforme necessário, dependendo da quantidade de dados escritos.

// Leitura e Escrita: Você pode escrever dados no buffer e depois lê-los, o que facilita o manuseio de dados temporários.

// Interfaces Compatíveis: Ele implementa as interfaces io.Reader e io.Writer, permitindo que seja usado em qualquer lugar onde essas interfaces sejam esperadas.

// **Quando Usar um Buffer?**

//Para manipular strings ou sequências de bytes dinamicamente.
//Quando você precisa concatenar dados de forma eficiente (evitando recriar strings repetidamente).
//Para implementar operações de leitura/escrita intermediárias entre streams de dados, como arquivos, conexões de rede, ou pipelines.

// Se você precisa de eficiência e manipulação temporária de dados em Go, o **buffer** é uma excelente ferramenta.
