package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/abdullahalsazib/orders_api/application"
)

func main() {
	fmt.Println("Order-api")
	app := application.New()

	ctx, cansel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cansel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("faild to start: ", err)
	}
}
