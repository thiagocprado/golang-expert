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
	// GORM já inclui o campo ID por padrão, mas podemos adicionar outros campos
	// como CreatedAt, UpdatedAt e DeletedAt se necessário
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Automaticamente cria a tabela se não existir
	db.AutoMigrate(&Product{})

	// Inserindo um produto
	// O GORM irá gerar o ID automaticamente
	// e o campo ID será incrementado automaticamente
	db.Create(&Product{Name: "Camiseta", Price: 49.90})

	products := []Product{
		{Name: "Calça", Price: 89.90},
		{Name: "Tênis", Price: 199.90},
		{Name: "Boné", Price: 29.90},
	}

	// Vários produtos podem ser inseridos de uma só vez
	db.Create(&products)

	var p Product
	// Busca o primeiro produto com ID 1
	db.First(&p, 1)
	fmt.Println(p)

	// Where
	db.First(&p, "name = ?", "Camiseta")
	fmt.Println(p)

	// Paginação
	db.Limit(2).Offset(2).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	db.Where("price > ?", 50.00).Find(&products)
	for _, p := range products {
		fmt.Println(p)
	}

	// Atualizando um produto
	db.First(&p, 1)
	p.Name = "New Mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	// Deletando um produto
	db.Delete(&p2, p2.ID)

	// Quando trabalhamos com DeleteAt o registro não é realmente removido do banco de dados,
	// mas marcado como deletado.

	// Para isso, precisamos adicionar o campo DeletedAt na struct
}
