package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type DollarQuotationResponse struct {
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

type DollarQuotation struct {
	ID         string    `json:"-"`
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
	ctx := context.Background()
	createDbFile()
	db, err := sql.Open("sqlite3", "./dollarQuotation.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTable(db)

	http.HandleFunc("/cotacao", GetDollarQuotationHandler(db, ctx))
	http.ListenAndServe(":8080", nil)
}

func GetDollarQuotationHandler(db *sql.DB, ctx context.Context) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request iniciada")
		defer log.Println("Request finalizada")

		log.Println("Iniciando busca dos dados...")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data, err := GetDollarQuotation(ctx, db)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(data)
	}
}

func GetDollarQuotation(ctx context.Context, db *sql.DB) (*DollarQuotation, error) {
	body, err := FetchDollarQuotation(ctx)
	if err != nil {
		return nil, err
	}

	var data DollarQuotationResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	dollarQuotation := MapQuotationToResponse(&data)

	err = insertDollarQuotation(ctx, db, dollarQuotation)
	if err != nil {
		return nil, err
	}

	return dollarQuotation, nil
}

func FetchDollarQuotation(ctx context.Context) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(200)*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func createTable(db *sql.DB) {
	createQuotationTableScript := `CREATE TABLE IF NOT EXISTS quotation (
		"id" VARCHAR(500) NOT NULL PRIMARY KEY,		
		"code" VARCHAR(3),
		"code_in" VARCHAR(3),
		"name" VARCHAR(100),
		"high" REAL,
		"low" REAL,
		"var_bid" REAL,
		"pct_change" REAL,
		"bid" REAL,
		"ask" REAL,
		"timestamp" TEXT,
		"create_date" TEXT
	  );`

	log.Println("Criando tabela")
	statement, err := db.Prepare(createQuotationTableScript)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Tabela criada")
}

func insertDollarQuotation(ctx context.Context, db *sql.DB, dollarQuotation *DollarQuotation) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	stmt, err := db.Prepare("INSERT INTO quotation(id, code, code_in, name, high, low, var_bid, pct_change, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, dollarQuotation.ID, dollarQuotation.Code, dollarQuotation.CodeIn, dollarQuotation.Name, dollarQuotation.High, dollarQuotation.Low, dollarQuotation.VarBid, dollarQuotation.PctChange, dollarQuotation.Bid, dollarQuotation.Ask, dollarQuotation.Timestamp, dollarQuotation.CreateDate)
	if err != nil {
		return err
	}
	return nil
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

func MapQuotationToResponse(dollarQuotation *DollarQuotationResponse) *DollarQuotation {
	return &DollarQuotation{
		ID:         uuid.New().String(),
		Code:       dollarQuotation.USDBRL.Code,
		CodeIn:     dollarQuotation.USDBRL.CodeIn,
		Name:       dollarQuotation.USDBRL.Name,
		High:       ConvertStringToFloat(dollarQuotation.USDBRL.High, 64),
		Low:        ConvertStringToFloat(dollarQuotation.USDBRL.Low, 64),
		VarBid:     ConvertStringToFloat(dollarQuotation.USDBRL.VarBid, 64),
		PctChange:  ConvertStringToFloat(dollarQuotation.USDBRL.PctChange, 64),
		Bid:        ConvertStringToFloat(dollarQuotation.USDBRL.Bid, 64),
		Ask:        ConvertStringToFloat(dollarQuotation.USDBRL.Ask, 64),
		Timestamp:  dollarQuotation.USDBRL.Timestamp,
		CreateDate: ConvertStringToTime(dollarQuotation.USDBRL.CreateDate),
	}
}

func createDbFile() {
	_, err := os.Stat("./dollarQuotation.db")

	if errors.Is(err, os.ErrNotExist) {
		log.Println("file does not exist")
		file, err := os.Create("./dollarQuotation.db")

		defer file.Close()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("file created")
	} else {
		log.Println("file exists")
	}
}
