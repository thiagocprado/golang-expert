package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// aqui trabalhamos com chave valor
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// através do valor da chave, conseguimos recuperar seu valor
	// mas é preciso ter cuidado, com os valores das chaves
	token := ctx.Value("token")
	fmt.Println(token)
}
