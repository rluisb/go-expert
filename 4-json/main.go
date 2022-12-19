package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	// Para omitir propriedades podemos utilizar -
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}

	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// JSON utilizando os nomes das propriedades da struct
	// jsonPuro := []byte(`{"Numero":2,"Saldo":200}`)

	// JSON utilizando os nomes das tags das propriedades da struct
	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		println(err)
	}
	println(contaX.Saldo)
}
