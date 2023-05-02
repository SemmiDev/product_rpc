package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"

	p "github.com/semmidev/product_rpc"
)

type ProductManager struct {
	products []p.Product
}

func (pm *ProductManager) AddProduct(p p.Product, reply *bool) error {
	pm.products = append(pm.products, p)
	*reply = true
	log.Println("Product Added: ", p)
	return nil
}

func (pm *ProductManager) GetProduct(id int, reply *p.Product) error {
	for _, p := range pm.products {
		if p.ID == id {
			*reply = p
			return nil
		}
	}
	return errors.New("Product not found")
}

func (pm *ProductManager) DeleteProduct(id int, reply *bool) error {
	for i, p := range pm.products {
		if p.ID == id {
			pm.products = append(pm.products[:i], pm.products[i+1:]...)
			*reply = true
			return nil
		}
	}
	return errors.New("Product not found")
}

func main() {
	ProductManager := &ProductManager{}

	server := rpc.NewServer()
	server.Register(ProductManager)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Println("Server is listening on port 1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			continue
		}
		go server.ServeConn(conn)
	}
}
