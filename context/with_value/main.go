package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// context.WithValue adiciona um valor ao contexto usando o esquema chave-valor.
	// Pode ser útil para passar informações como tokens, IDs ou configurações entre funções.
	// ⚠️ Mas evite usar para dados que podem ser passados como argumento direto. Use com moderação.
	ctx = context.WithValue(ctx, "token", "senha123")

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// ctx.Value permite recuperar o valor associado a uma chave
	// É importante garantir que a chave seja única (de preferência, um tipo próprio)
	token := ctx.Value("token")

	// Aqui estamos apenas imprimindo o valor recuperado
	fmt.Println("Token recebido:", token)
}
