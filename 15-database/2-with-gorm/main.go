package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//Insert or Create
	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.00,
	})

	//Insert or create batch
	products := []Product{
		{Name: "Mouse", Price: 50.00},
		{Name: "Keyboard", Price: 100.00},
	}

	db.Create(products)
}
