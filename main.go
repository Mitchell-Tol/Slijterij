package main

import (
    "fmt"
    "net/http"
    "slijterij/api/base"
    "slijterij/api/base/drinks"
    "slijterij/api/base/bar"
	"slijterij/api/base/device"
	"slijterij/api/base/category"
	"slijterij/api/base/order"
    "slijterij/db"
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

    fmt.Println("Running...")
    http.ListenAndServe(":8080", mux)
}
