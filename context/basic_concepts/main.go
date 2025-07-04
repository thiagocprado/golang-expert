package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Context é usado para controlar o tempo de execução de algo:
	// por exemplo, cancelar uma operação se ela demorar demais.
	// Muito útil em chamadas de API, goroutines, etc.

	ctx := context.Background()

	// Cria um contexto com limite de 3 segundos.
	// Depois disso, ele é cancelado automaticamente.
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel() // Garante que os recursos do contexto serão liberados

	bookHotel(ctx)
}

// A convenção em Go é passar o contexto como o primeiro parâmetro
func bookHotel(ctx context.Context) {
	// Funciona como um switch, mas é usado para esperar por múltiplos canais ao mesmo tempo
	select {
	case <-ctx.Done():
		// Isso acontece se o contexto for cancelado antes da operação terminar
		fmt.Println("Reserva de hotel cancelada. Tempo esgotado.")
		return
	case <-time.After(1 * time.Second):
		// Simula uma operação que demora 1 segundo (ex: reservar hotel)
		fmt.Println("Hotel reservado com sucesso.")
	}
}
