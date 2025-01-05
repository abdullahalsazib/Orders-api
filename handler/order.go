package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order...!")
	w.Write([]byte("<h1>Create and order</h1>"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all order..!")
	w.Write([]byte("<h1>List all order</h1>"))
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get an order By Id..!")
	w.Write([]byte("<h1>Get an order by id</h1>"))
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order By Id")
	w.Write([]byte("<h1>Update and order by id</h1>"))
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by id")
	w.Write([]byte("<h1>Delete an order by id</h1>"))
}
