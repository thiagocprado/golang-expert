package main

import (
	"log"
	"net/http"
)

func main() {
	// servidores arquivo são utilizados para renderizar paginas html do lado do servidor
	fileServer := http.FileServer(http.Dir("./file_server/public"))

	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Blog page"))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
