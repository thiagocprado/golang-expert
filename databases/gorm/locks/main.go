package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
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

	// db.Create(&Category{Name: "Eletrônicos"})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		panic(err)
	}
	c.Name = "Cozinha"
	tx.Debug().Save(&c)
	tx.Commit()
}

// lock otimista
// o lock otimista é usado para evitar conflitos de concorrência
// ele não bloqueia o registro no banco de dados, mas verifica se o registro foi
// modificado por outra transação antes de salvar as alterações

// lock pessimista
// o lock pessimista é usado para evitar conflitos de concorrência
// ele bloqueia o registro no banco de dados, impedindo que outras transações
// acessem o registro até que a transação atual seja concluída
