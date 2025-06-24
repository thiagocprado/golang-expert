package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Account struct {
	Number  int
	Balance int
}

type AccountV2 struct {
	// As "tags" permitem customizar o nome dos campos no JSON
	Number int `json:"n"`
	// Também é possível adicionar validações (ex: `gt-0` para maior que zero, se for interpretado por uma lib de validação)
	Balance int `json:"b validate=gt-0"`
	// O campo com `json:"-"` será ignorado na serialização/desserialização
	Digit int `json:"-"`
}

func main() {
	// ----------- struct > json -----------

	account := Account{Number: 1, Balance: 100}

	// Marshal serializa a struct para JSON e retorna um slice de bytes
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res)) // Convertemos os bytes para string para imprimir no terminal

	// Encoder serializa direto para um destino (ex: terminal, arquivo, resposta HTTP, etc)
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	// ----------- json > struct -----------

	// Um JSON cru que será desserializado
	pureJson := []byte(`{"Number": 2, "Balance": 200}`)
	var accountX Account

	// Unmarshal desserializa o JSON para dentro da struct, modificando o valor original via ponteiro
	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		println(err)
	}
	fmt.Println(accountX)

	// Decoder desserializa a partir de uma stream (ex: leitor de strings, arquivos, etc)
	err = json.NewDecoder(strings.NewReader(string(pureJson))).Decode(&accountX)
	if err != nil {
		println(err)
	}

	// ----------- json com nomes de campos diferentes -----------

	// Esse JSON usa nomes diferentes (n, b, d), por isso precisamos usar tags na struct
	pureJsonV2 := []byte(`{"n": 2, "b": 200, "d": 1}`)
	var accountY AccountV2

	err = json.Unmarshal(pureJsonV2, &accountY)
	if err != nil {
		println(err)
	}
	fmt.Println(accountY)
}
