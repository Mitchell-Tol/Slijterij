package main

import (
    "fmt"
    "net/http"
    "slijterij/api/base"
    "slijterij/api/base/drinks"
    "slijterij/api/base/rooms"
)

func main() {
    fmt.Println("Starting API Service")

    mux := http.NewServeMux()
    mux.Handle("/", base.NewHandler())
    mux.Handle("/drinks", drinks.NewHandler())
    mux.Handle("/rooms", rooms.NewHandler(rooms.REGULAR))
    mux.Handle("/rooms/login", rooms.NewHandler(rooms.LOGIN))

    fmt.Println("Running...")
    http.ListenAndServe(":8080", mux)
}
