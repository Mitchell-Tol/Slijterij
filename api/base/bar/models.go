package bar

type Bar struct {
    Id string `json:"id"`
    Password string `json:"password"`
}

type TokenizedBar struct {
    Id string `json:"id"`
    Token string `json:"token"`
}
