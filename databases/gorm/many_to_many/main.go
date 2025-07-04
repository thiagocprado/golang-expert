package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Many to Many precisa de uma tabela intermediária para relacionar as duas entidades
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	c := Category{Name: "Cozinha"}
	db.Create(&c)

	c2 := Category{Name: "Eletronicos"}
	db.Create(&c2)

	p := Product{Name: "Panela", Price: 99.99, Categories: []Category{c, c2}}
	db.Create(&p)

	var categories []Category
	// Serial Numbers são carregados junto com os Produtos
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
