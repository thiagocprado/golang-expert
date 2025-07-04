package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	// Produto pertence a uma categoria
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

	// Automaticamente cria a tabela se não existir
	db.AutoMigrate(&Product{}, &Category{})

	// Cria uma categoria
	c := Category{Name: "Eletrônicos"}
	db.Create(&c)

	// Cria um produto associado à categoria
	p := Product{Name: "Smartphone", Price: 999.99, CategoryID: c.ID}
	db.Create(&p)

	p = Product{Name: "Notebook", Price: 1999.99, CategoryID: c.ID}
	db.Create(&p)

	var products []Product
	// Busca todos os produtos e suas categorias associadas
	db.Preload("Category").Preload("SerialNumber").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name)
	}
}
