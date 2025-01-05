package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRouters(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	err := a.rdb.Ping(ctx).Err()

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("Faild to close redis", err)
		}
	}()

	if err != nil {
		return fmt.Errorf("faild to connect to redis: %w", err)
	}
	fmt.Println("Starting server...")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("faild to start server %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancal := context.WithTimeout(context.Background(), time.Second*10)
		defer cancal()
		return server.Shutdown(timeout)
	}

	// ctx.Done()
	// err, open := <- ch

	return nil
}
