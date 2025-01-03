package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Wellcome on GO-Lang")
	// CHI to Create Router
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.Logger)
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Recoverer)

	router.Get("/hello", basicHandler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	CheckError(err)

}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World"))
}
