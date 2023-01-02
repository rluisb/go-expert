package main

import (
	"fmt"
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
	//db.Create(&Product{
	//	Name:  "Notebook",
	//	Price: 1000.00,
	//})

	//Insert or create batch
	//products := []Product{
	//	{Name: "Mouse", Price: 50.00},
	//	{Name: "Keyboard", Price: 100.00},
	//}
	//
	//db.Create(products)

	//Select First by id then by name
	//var product Product
	//db.First(&product, 4)
	//fmt.Println(product)
	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	//Select all
	//var products []Product
	//db.Limit(2).Offset(1).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}
	//}

	//select where
	//var products []Product
	//db.Where("price > ?", 90).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	//select where like
	//var products []Product
	//db.Where("name LIKE ?", "%book%").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}

	var products []Product
	db.Where("price > ?", 90).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
