package main

import (
	"log"
	"net/http"
)

func main() {
	// Criamos um FileServer para servir arquivos estáticos (HTML, CSS, JS, imagens etc.)
	// a partir do diretório "./file_server/public".
	// Isso é útil para renderizar páginas HTML diretamente do servidor.
	fileServer := http.FileServer(http.Dir("./file_server/public"))

	// Criamos nosso próprio multiplexer (ServeMux), que nos dá mais controle sobre o roteamento das requisições.
	mux := http.NewServeMux()

	// O método Handle é usado para associar um handler que implementa a interface http.Handler.
	// Nesse caso, o fileServer é um handler válido, pois ele implementa ServeHTTP.
	mux.Handle("/", fileServer)

	// Já o HandleFunc é utilizado quando queremos associar uma função diretamente ao caminho.
	// Internamente, o HandleFunc transforma a função passada em um http.HandlerFunc,
	// que também implementa a interface http.Handler.
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog page"))
	})

	// Inicia o servidor na porta 8080 utilizando o mux para rotear as requisições.
	// log.Fatal é usado para registrar e encerrar o programa se ocorrer um erro.
	log.Fatal(http.ListenAndServe(":8080", mux))
}
