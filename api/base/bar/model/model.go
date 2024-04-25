package model

type Bar struct {
    Id string `json:"id"`
    Password string `json:"password"`
}

type BarEntity struct {
    Id string
    Password string
    Token string
}

type TokenizedBar struct {
    Id string `json:"id"`
    Token string `json:"token"`
}
