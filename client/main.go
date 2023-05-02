package main

import (
	"fmt"
	"log"
	"net/rpc"

	p "github.com/semmidev/product_rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Add a new product
	newProduct := &p.Product{
		ID:    1,
		Name:  "New Product",
		Price: 12.3,
	}

	var addResult bool
	err = client.Call("ProductManager.AddProduct", newProduct, &addResult)
	if err != nil {
		log.Fatal("add product error:", err)
	}

	if addResult {
		fmt.Println("Product added successfully!")
	} else {
		fmt.Println("Failed to add product")
	}
}
