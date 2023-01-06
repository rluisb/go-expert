package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type DollarQuotation struct {
	Bid float64 `json:"bid"`
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(300)*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var data DollarQuotation
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	UpsertDataToFile(data)
}

func UpsertDataToFile(data DollarQuotation) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	templateForFileMust := template.Must(template.New("templateForFile").Parse("Dolar: {{.Bid}}"))
	err = templateForFileMust.Execute(f, data)
	if err != nil {
		panic(err)
	}

	log.Printf("Arquivo criado com sucesso!")
	f.Close()
}
