package main

import (
	"fmt"
	"net/http"
	"slijterij/api/base"
	"slijterij/api/base/drinks"
)

func main() {
	fmt.Println("Starting API Service")

	mux := http.NewServeMux()
	mux.Handle("/", base.NewHandler())
	mux.Handle("/drinks/", drinks.NewHandler())

	fmt.Println("Running...")
	http.ListenAndServe(":8080", mux)
}
