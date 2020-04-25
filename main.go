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
