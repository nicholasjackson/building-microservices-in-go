package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	memdb "github.com/hashicorp/go-memdb"
	nats "github.com/nats-io/go-nats"
)

type Product struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

var schema *memdb.DBSchema
var db *memdb.MemDB
var natsClient *nats.Conn

var natsServer = flag.String("nats", "", "NATS server URI")

func init() {
	flag.Parse()

	schema = &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"product": &memdb.TableSchema{
				Name: "product",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "SKU"},
					},
				},
			},
		},
	}

	err := schema.Validate()
	if err != nil {
		log.Fatal(err)
	}

	db, err = memdb.NewMemDB(schema)
	if err != nil {
		log.Fatal(err)
	}

	txn := db.Txn(true)

	if err := txn.Insert("product", Product{"Test1", "ABC232323"}); err != nil {
		log.Fatal(err)
	}

	if err := txn.Insert("product", Product{"Test2", "ABC883388"}); err != nil {
		log.Fatal(err)
	}

	txn.Commit()

	natsClient, err = nats.Connect("nats://" + *natsServer)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Starting Product Service")

	http.DefaultServeMux.HandleFunc("/product", productsHandler)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func productsHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		insertProduct(rw, r)
	} else {
		getProducts(rw, r)
	}
}

func insertProduct(rw http.ResponseWriter, r *http.Request) {
	log.Println("/insert handler called")

	product := &Product{}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, product)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	txn := db.Txn(true)
	if err := txn.Insert("product", product); err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	txn.Commit()

	natsClient.Publish("product.inserted", data)
}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	log.Println("/get handler called")

	txn := db.Txn(false)
	results, err := txn.Get("product", "id")
	if err != nil {
		log.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	products := make([]Product, 0)
	for {
		obj := results.Next()
		if obj == nil {
			break
		}

		fmt.Println(r)
		products = append(products, obj.(Product))
	}

	encoder := json.NewEncoder(rw)
	encoder.Encode(products)
}
