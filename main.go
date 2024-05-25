package main

import (
    "fmt"
    "net/http"
    "slijterij/api/base"
    "slijterij/api/base/drinks"
    "slijterij/api/base/bar"
	"slijterij/api/base/device"
    "slijterij/db"
)

func main() {
    fmt.Println("Starting API Service")

    SetEnvVars()
    store := db.NewStore()

    mux := http.NewServeMux()
    mux.Handle("/", base.NewHandler())
    mux.Handle("/drinks", drinks.NewHandler(store))
    mux.Handle("/bar", bar.NewHandler(store, bar.REGULAR))
    mux.Handle("/bar/login", bar.NewHandler(store, bar.LOGIN))
    mux.Handle("/device", device.NewHandler(store))

    fmt.Println("Running...")
    http.ListenAndServe(":8080", mux)
}
