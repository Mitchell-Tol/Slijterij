package rooms

type Room struct {
    Id string `json:"id"`
    Password string `json:"password"`
}

type TokenizedRoom struct {
    Id string `json:"id"`
    Token string `json:"token"`
}
