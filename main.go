package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Wellcome on GO-Lang")

	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	// server running on http://localhost:3000

}

func basicHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, World"))
}
