package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Item struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Qty   int     `json:"qty"`
}

type Checkout struct {
	SKU string `json:"sku"`
	Qty int    `json:"qty"`
}

var items []Item

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(items)
}

func CheckPromo(skus []string) float32 {
	var (
		total       float32
		macPromo    = []string{"43N23P", "234234"}
		googlePromo = []string{"120P90", "120P90", "120P90"}
		alexaPromo  = []string{"A304SD", "A304SD", "A304SD"}
	)

	for _, item := range items {
		if reflect.DeepEqual(macPromo, skus) {
			if item.SKU == macPromo[0] {
				total = item.Price
			}
		}

		if reflect.DeepEqual(googlePromo, skus) {
			if item.SKU == googlePromo[0] {
				total = item.Price * 2
			}
		}

		if reflect.DeepEqual(alexaPromo, skus) {
			if item.SKU == alexaPromo[0] {
				total = (item.Price * 3) - (item.Price*3)*10/100
			}
		}
	}
	return total
}

func CheckoutItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var checkoutItems []Checkout

	_ = json.NewDecoder(r.Body).Decode(&checkoutItems)

	var skus []string
	for i := 0; i < len(checkoutItems); i++ {
		skus = append(skus, checkoutItems[i].SKU)
	}

	json.NewEncoder(w).Encode(map[string]float32{
		"Total": CheckPromo(skus),
	})
}

func main() {
	// Mock data items
	items = append(items, Item{SKU: "120P90", Name: "Google Home", Price: 49.99, Qty: 10})
	items = append(items, Item{SKU: "43N23P", Name: "MacBook Pro", Price: 5399.99, Qty: 5})
	items = append(items, Item{SKU: "A304SD", Name: "Alexa Speaker", Price: 109.50, Qty: 10})
	items = append(items, Item{SKU: "234234", Name: "Raspberyy Pi B", Price: 30.00, Qty: 2})

	// Router
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/health", HealthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/items", GetAllItems).Methods(http.MethodGet)
	router.HandleFunc("/checkout", CheckoutItems).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
