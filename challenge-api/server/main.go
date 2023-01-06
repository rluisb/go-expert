package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

type DollarQuotation struct {
	USDBRL struct {
		Code       string `json:"code"`
		CodeIn     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type DollarQuotationResponse struct {
	Code       string    `json:"code"`
	CodeIn     string    `json:"codeIn"`
	Name       string    `json:"name"`
	High       float64   `json:"high"`
	Low        float64   `json:"low"`
	VarBid     float64   `json:"varBid"`
	PctChange  float64   `json:"pctChange"`
	Bid        float64   `json:"bid"`
	Ask        float64   `json:"ask"`
	Timestamp  string    `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func main() {
	http.HandleFunc("/cotacao", GetDollarQuotationHandler)
	http.ListenAndServe(":8080", nil)
}

func GetDollarQuotationHandler(w http.ResponseWriter, _ *http.Request) {
	data, err := GetDollarQuotation()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&DollarQuotationResponse{
		Code:       data.USDBRL.Code,
		CodeIn:     data.USDBRL.CodeIn,
		Name:       data.USDBRL.Name,
		High:       ConvertStringToFloat(data.USDBRL.High, 64),
		Low:        ConvertStringToFloat(data.USDBRL.Low, 64),
		VarBid:     ConvertStringToFloat(data.USDBRL.VarBid, 64),
		PctChange:  ConvertStringToFloat(data.USDBRL.PctChange, 64),
		Bid:        ConvertStringToFloat(data.USDBRL.Bid, 64),
		Ask:        ConvertStringToFloat(data.USDBRL.Ask, 64),
		Timestamp:  data.USDBRL.Timestamp,
		CreateDate: ConvertStringToTime(data.USDBRL.CreateDate),
	})
}

func GetDollarQuotation() (*DollarQuotation, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data DollarQuotation
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func ConvertStringToTime(value string) time.Time {
	const format = "2006-01-02 15:04:05"
	date, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}
	return date
}

func ConvertStringToFloat(value string, bitSize int) float64 {
	converted, err := strconv.ParseFloat(value, bitSize)
	if err != nil {
		panic(err)
	}
	return converted
}
