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
type Product struct {
  ProductID int
  Name string
  Options string
  Price float64
  Image string
}
func (store *dbStore) GetProducts() []Products {
  rows, err := store.db.Query("SELECT * FROM products ORDER BY id ASC")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  products := []Products{}
  for rows.Next() {
    product := Products{}

    _ = rows.Scan(&product.ProductID, &product.Name, &product.Total, &product.Image)
    products = append(products, product)
  }

  return products
}
func (store *dbStore) GetSection(section string) []Product {
  rows, err := store.db.Query("SELECT * FROM " + section + "Products ORDER BY id ASC")
  if err != nil {
    panic(err)
  }
  defer rows.Close()

  products := []Product{}
  for rows.Next() {
    product := Product{}

    _ = rows.Scan(&product.ProductID, &product.Name, &product.Options, &product.Price, &product.Image)
    products = append(products, product)
  }
  return products
}
func (store *dbStore) GetProduct(section string, pID string) Product {
  id, _ := strconv.Atoi(pID)
  row := store.db.QueryRow("SELECT * FROM " + section + "Products WHERE id=$1", id)
  product := Product{}
  _ = row.Scan(&product.ProductID, &product.Name, &product.Options, &product.Price, &product.Image)
  return product
}

var store dbStore
func InitStore(s dbStore) {
  store = s
}
