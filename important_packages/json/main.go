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
	// temos as chamadas tags para trabalhamos com chave e valor | podem ser chamadas de anotações também
	Number int `json:"n"`
	// podemos fazer validações também
	Balance int `json:"b validate=gt-0"`
	// podemos omitir valores também
	Digit int `json:"-"`
}

func main() {
	// struct > json //

	account := Account{Number: 1, Balance: 100}

	// marshal é serializar em json e o seu retorno é em bytes, quando usamos essa abordagem, guardamos o valor do json
	res, err := json.Marshal(account)
	if err != nil {
		println(err)
	}
	println(string(res))

	// o encoder faz a serialização e já entrega para um destino final (terminal, arquivo, web server...)
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		println(err)
	}

	// json > struct //

	pureJson := []byte(`{"Number": 2, "Balance": 200}`)
	var accountX Account

	// com o unmarshal, nós vamos pegar o json e passar para nossa struct
	// é importante observar que ele vai alterar o valor real da variável, por isso está passando o endereço de memória
	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		println(err)
	}
	fmt.Println(accountX)

	err = json.NewDecoder(strings.NewReader(string(pureJson))).Decode(&accountX)
	if err != nil {
		println(err)
	}

	// podemos ter um json com nome de campos distintos do da struct
	pureJsonV2 := []byte(`{"n": 2, "b": 200, "d": 1}`)
	var accountY AccountV2

	err = json.Unmarshal(pureJsonV2, &accountY)
	if err != nil {
		println(err)
	}
	fmt.Println(accountY)
}
