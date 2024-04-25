package main

import (
    "fmt"
    "net/http"
    "slijterij/api/base"
    "slijterij/api/base/drinks"
    "slijterij/api/base/bar"
)

func main() {
    fmt.Println("Starting API Service")

    mux := http.NewServeMux()
    mux.Handle("/", base.NewHandler())
    mux.Handle("/drinks", drinks.NewHandler())
    mux.Handle("/bar", bar.NewHandler(bar.REGULAR))
    mux.Handle("/bar/login", bar.NewHandler(bar.LOGIN))

    fmt.Println("Running...")
    http.ListenAndServe(":8080", mux)
}
