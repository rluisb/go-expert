package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryID int
	Category   Category
	//SerialNumber SerialNumber
	gorm.Model
}

//type SerialNumber struct {
//	ID        int `gorm:"primaryKey"`
//	Number    string
//	ProductID int
//}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	db.AutoMigrate(&Product{}, &Category{})
	//Create Category
	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	// Create Product
	product := Product{
		Name:       "Mouse",
		Price:      1000.00,
		CategoryID: category.ID,
	}
	db.Create(&product)

	//db.Create(&SerialNumber{
	//	Number:    "123456",
	//	ProductID: product.ID,
	//})

	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}

	//var categories []Category
	//err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	//if err != nil {
	//	panic(err)
	//}
	//for _, category := range categories {
	//	fmt.Println(category.Name, ":")
	//	for _, product := range category.Products {
	//		println("- ", product.Name, category.Name)
	//	}
	//}
}
