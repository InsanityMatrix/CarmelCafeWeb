package main
import (
  "database/sql"
)

type dbStore struct {
  db *sql.DB
}
type Products struct {
  ProductID int
  Name string
  Image string
  Total int
}
func (store *dbStore) GetProducts() []Products {
  rows, _ := store.db.Query("SELECT * FROM products ORDER BY id ASC")
  defer rows.Close()

  products := []Products{}
  for rows.Next() {
    product := Products{}

    _ = rows.Scan(&product.ProductID, &product.Name, &product.Total, &product.Image)
    products = append(products, product)
  }

  return products
}


var store dbStore
func InitStore(s dbStore) {
  store = s
}
