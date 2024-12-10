package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// criação de arquivo
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// aqui não temos certeza do que vamos escrever no arquivo, por isso usamos o byte
	size, err := f.Write([]byte("Dados tipo string, mas poderia ter outros valores!\n"))
	if err != nil {
		panic(err)
	}

	// aqui temos certeza que estamos escrevendo uma string
	size, err = f.WriteString("Dados tipo string com certeza!\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created successfully! Size: %d bytes\n", size)

	// encerra as operações com o arquivo
	f.Close()

	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	// o que fica salvo em arquivos são slices de bytes, por isso convertemos para string
	fmt.Println(string(file))

	// leitura de pouco em pouco abrindo o arquivo. Serve para arquivos com tamanho muito grande
	fileV2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// vamos trabalhar com buffer
	// que nada mais é do uma área temporária de armazenamento usada para gerenciar dados enquanto eles estão sendo transferidos de um lugar para outro
	reader := bufio.NewReader(fileV2)
	// aqui o buffer é um slice de bytes
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
