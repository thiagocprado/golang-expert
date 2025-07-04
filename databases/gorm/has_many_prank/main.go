package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Uma categoria pode ter vários produtos
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
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// c := Category{Name: "Eletrônicos"}
	// db.Create(&c)

	c := Category{Name: "Cozinha"}
	db.Create(&c)

	// p := Product{Name: "Smartphone", Price: 999.99, CategoryID: 1}
	// db.Create(&p)

	p := Product{Name: "Panela", Price: 99.99, CategoryID: 2}
	db.Create(&p)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		println(c.Name)
		for _, p := range c.Products {
			println(" Produto ->", p.Name)
		}
	}
}
