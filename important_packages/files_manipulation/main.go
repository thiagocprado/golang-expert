package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criação de um novo arquivo chamado "test.txt".
	// Se o arquivo já existir, ele será truncado (conteúdo apagado).
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	// Escrita no arquivo usando um slice de bytes.
	// Ideal quando os dados podem não ser apenas strings (ex: binários).
	size, err := f.Write([]byte("Dados tipo string, mas poderia ter outros valores!\n"))
	if err != nil {
		panic(err)
	}

	// Escrita no arquivo diretamente como string.
	// Útil quando se tem certeza de que os dados são texto.
	size, err = f.WriteString("Dados tipo string com certeza!\n")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes\n", size)

	// Fecha o arquivo após as operações de escrita.
	f.Close()

	// Leitura do arquivo completo de uma vez só.
	// Pode ser ineficiente para arquivos grandes.
	file, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	// Como os dados do arquivo são lidos como slice de bytes ([]byte),
	// convertemos para string para exibir no terminal.
	fmt.Println(string(file))

	// Abertura do arquivo para leitura em partes (streaming).
	fileV2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	// Criação de um leitor com buffer.
	// Buffers são úteis para ler dados em blocos pequenos, melhorando a performance
	// especialmente com arquivos grandes.
	reader := bufio.NewReader(fileV2)

	// Buffer de 10 bytes para leitura parcial.
	buffer := make([]byte, 10)

	for {
		// Lê até 10 bytes por vez e retorna a quantidade real lida (n).
		n, err := reader.Read(buffer)
		if err != nil {
			// Fim do arquivo ou erro na leitura.
			break
		}

		// Mostra os dados lidos convertendo para string.
		fmt.Println(string(buffer[:n]))
	}

	// Remove o arquivo do sistema após a leitura.
	err = os.Remove("test.txt")
	if err != nil {
		panic(err)
	}
}
