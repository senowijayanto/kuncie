package main

import (
	"fmt"
)

type Item struct {
	SKU   string  `json:"sku"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Qty   int     `json:"qty"`
}

var items []Item

func main() {
	fmt.Println("Checkout API")

	// Mock data items
	items = append(items, Item{SKU: "120P90", Name: "Google Home", Price: 49.99, Qty: 10})
	items = append(items, Item{SKU: "43N23P", Name: "MacBook Pro", Price: 5399.99, Qty: 5})
	items = append(items, Item{SKU: "A304SD", Name: "Alexa Speaker", Price: 109.50, Qty: 10})
	items = append(items, Item{SKU: "234234", Name: "Raspberyy Pi B", Price: 30.00, Qty: 2})
}
