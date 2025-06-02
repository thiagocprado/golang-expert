package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// No Go, contexto (context) é usado para controlar prazos, cancelamentos e metadados em operações que podem ser interrompidas,
	// como chamadas de API, goroutines ou processamento de dados.

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	bookHotel(ctx)
}

// variaveis de contexto sempre são o primeiro parâmetro da função
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return

	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
