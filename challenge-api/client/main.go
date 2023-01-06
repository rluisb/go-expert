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
	ID         string    `json:"-"`
	Code       string    `json:"-"`
	CodeIn     string    `json:"-"`
	Name       string    `json:"-"`
	High       float64   `json:"-"`
	Low        float64   `json:"-"`
	VarBid     float64   `json:"-"`
	PctChange  float64   `json:"-"`
	Bid        float64   `json:"bid"`
	Ask        float64   `json:"-"`
	Timestamp  string    `json:"-"`
	CreateDate time.Time `json:"-"`
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
