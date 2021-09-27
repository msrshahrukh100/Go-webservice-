package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Product struct {
	Id   int    `json:"productId"`
	Name string `json:"product_name"`
}

var productList []Product

func init() {
	productJson := `[{
			"productId": 1,
			"product_name": "Soap"
		},
		{
			"productId": 2,
			"product_name": "Mobile"
		}
	]`

	err := json.Unmarshal([]byte(productJson), &productList)
	if err != nil {
		fmt.Println(err)
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split()
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(bodyBytes, &newProduct)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if newProduct.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return

	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product/", productHandler)
	http.ListenAndServe(":5000", nil)
}
