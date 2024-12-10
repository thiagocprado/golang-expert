package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

func main() {
	// através da rota é chamada uma função
	http.HandleFunc("/", SearchCepHandler)

	// multiplexer - basicamente é atracar funções a um componente, nesse caso estamos usando o padrão global do go
	http.ListenAndServe(":8080", nil)
}

// conseguimos pegar todos os dados de resposta e também de requsisição
func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Rota não encontrada!"))
		return
	}

	// manipular query string
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("CEP não informado"))
		return
	}

	res, err := SearchCep(cepParam)
	if err != nil {
		if err.Error() == "CEP não encontrado!" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write([]byte(err.Error()))
		return
	}

	// conseguimos manipular os headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// aqui vamos trabalhar com ponteiro para termos uma mudança global
func SearchCep(cep string) (*Address, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode == http.StatusBadRequest {
		return nil, errors.New("CEP não encontrado!")
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var address Address
	err = json.Unmarshal(res, &address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}
