package main

import (
	"context"
	"fmt"

	"github.com/abdullahalsazib/orders_api/application"
)

func main() {
	fmt.Println("Order-api")
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("faild to start: ", err)
	}
}
