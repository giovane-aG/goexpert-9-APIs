package main

import (
	"fmt"

	"github.com/giovane-aG/goexpert/9-APIs/configs"
	"github.com/giovane-aG/goexpert/9-APIs/internal/entity"
	"github.com/giovane-aG/goexpert/9-APIs/internal/infra/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb(config *configs.Conf) *gorm.DB {
	var dsn string = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	config := configs.LoadConfig("./")
	db := initDb(config)

	productModel := database.NewProduct(db)
	newProduct, err := entity.NewProduct("GALAX RTX 3070OC", 4000.00)
	if err != nil {
		panic(err)
	}

	err = newProduct.Validate()
	if err != nil {
		panic(err)
	}

	productModel.Create(newProduct)

	products, err := productModel.FindAll(0, 0, "")
	if err != nil {
		panic(err)
	}

	for _, v := range products {
		fmt.Println(v)
	}
}
