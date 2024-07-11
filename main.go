package main

import (
	"fmt"
	"net/http"
	"slijterij/api/base"
	"slijterij/api/base/bar"
	"slijterij/api/base/category"
	"slijterij/api/base/device"
	"slijterij/api/base/drinks"
	"slijterij/api/base/order"
	"slijterij/db"
	"github.com/rs/cors"
)

func main() {
    fmt.Println("Starting API Service")

    store := db.NewStore()

	baseHandler := base.NewHandler()
	drinksHandler := drinks.NewHandler(store)
	categoryHandler := category.NewHandler(store)
	deviceHandler := device.NewHandler(store)
	orderHandler := order.NewHandler(store)
	barHandler := bar.NewHandler(store, bar.REGULAR)
	loginHandler := bar.NewHandler(store, bar.LOGIN)

    mux := http.NewServeMux()
    mux.Handle("/", baseHandler)
    mux.Handle("/drinks", drinksHandler)
    mux.Handle("/bar", barHandler)
    mux.Handle("/bar/login", loginHandler)
    mux.Handle("/device", deviceHandler)
	mux.Handle("/category", categoryHandler)
	mux.Handle("/order", orderHandler)

	corsSettings := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "http://172.235.160.66:*"},
		AllowCredentials: true,
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	})
	handler := corsSettings.Handler(mux)

    fmt.Println("Running...")
    http.ListenAndServe(":8080", handler)
}
