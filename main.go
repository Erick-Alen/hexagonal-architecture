package main

import (
	"database/sql"

	db2 "github.com/Erick-Alen/hexagonal-architecture/adapters/db"
	"github.com/Erick-Alen/hexagonal-architecture/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite:")
	productAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productAdapter)
	product, _ := productService.Create("Product Test", 25)
	productService.Enable(product)
}
 