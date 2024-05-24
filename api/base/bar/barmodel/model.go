package barmodel

type Bar struct {
    Id string `json:"id"`
    Password string `json:"password"`
}

type BarEntity struct {
    Id string `json:"id"`
    Password string `json:"password"`
    Token string `json:"token"`
}

type TokenizedBar struct {
    Id string `json:"id"`
    Token string `json:"token"`
}

func MapEntityToTokenized(entity BarEntity) TokenizedBar {
    return TokenizedBar{entity.Id, entity.Token}
}
