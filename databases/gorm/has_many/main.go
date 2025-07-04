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
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	c := Category{Name: "Cozinha"}
	db.Create(&c)

	p := Product{Name: "Panela", Price: 99.99, CategoryID: 1}
	db.Create(&p)

	s := SerialNumber{Number: "SN123456", ProductID: 1}
	db.Create(&s)

	var categories []Category
	// Serial Numbers são carregados junto com os Produtos
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		println(c.Name)
		for _, p := range c.Products {
			println(" Produto ->", p.Name)
			println("  Número de Série ->", p.SerialNumber.Number)
		}
	}
}
