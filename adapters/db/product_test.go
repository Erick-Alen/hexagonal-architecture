package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/Erick-Alen/hexagonal-architecture/adapters/db"
	"github.com/Erick-Alen/hexagonal-architecture/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal((err.Error()))
	}
	stmt.Exec()
	// _, _ = db.Exec("create table products (id text, name text, price int, status text)")
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values (
		"abc",
		"Product Test",
		0,
		"disabled"
	);`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
	// if err != nil {
	// 	t.Errorf("Error: %s", err.Error())
	// }
	// if product.GetID() != "abc" {
	// 	t.Errorf("Expected product ID is abc, but got %s", product.GetID())
	// }
}
func TestProductdb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 0.0

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}

// func createProduct(db *sql.DB, id string, name string, price float64, status string) {
// 	stmt, err := db.Prepare("insert into products(id, name, price, status) values(?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	stmt.Exec(id, name, price, status)
// }
