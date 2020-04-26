package main
import (
  "database/sql"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "os"
  "log"
  "fmt"
  _ "github.com/lib/pq"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/products", productListHandler)
  r.HandleFunc("/section", sectionHandler)
  r.HandleFunc("/product", productHandler)
  r.HandleFunc("/order", orderHandler)
  return r
}
//ENVIRONMENT VARIABLES: CCDB_URL
func main() {
  router := newRouter()
  port := ":25510"
  url := os.Getenv("CCDB_URL")
  db, err := sql.Open("postgres",url)

  if err != nil {
    log.Fatalf("Connection error: %s", err.Error())
    panic(err)
  }
  defer db.Close()
  err = db.Ping()
  if err != nil {
    log.Fatalf("Ping error: %s", err.Error())
    panic(err)
  }
  InitStore(dbStore{db: db})
  http.ListenAndServe(port, router)
}

func productListHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  products := store.GetProducts()
  data, _ := json.Marshal(products)
  fmt.Fprint(w, string(data))
}
func sectionHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Fprint(w, "false")
    return
  }
  w.Header().Set("Content-Type", "application/json")
  section := r.Form.Get("Section")
  products := store.GetSection(section)
  data, _ := json.Marshal(products)
  fmt.Fprint(w, string(data))
}
func productHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Fprint(w, "false")
    return
  }
  w.Header().Set("Content-Type", "application/json")
  section := r.Form.Get("Section")
  id := r.Form.Get("ProductID")
  product := store.GetProduct(section, id)
  data, _ := json.Marshal(product)
  fmt.Fprint(w, string(data))
}
func orderHandler(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Fprint(w, "false")
    return
  }
  //name := r.Form.Get("Name")
  //studID := r.Form.Get("StudentID")
  orderJSON := r.Form.Get("Order")
  //order will be json array
  var orders []Order
  json.Unmarshal([]byte(orderJSON),&orders)
  //We can now set up the Square Orders with the order info we have, ez dubs
  fmt.Fprint(w, orders)
}
