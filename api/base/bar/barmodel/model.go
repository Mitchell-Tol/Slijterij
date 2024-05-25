package barmodel

type Bar struct {
	Name string `json:"name"`
    Password string `json:"password"`
}

type BarEntity struct {
    Id string `json:"id"`
	Name string `json:"name"`
    Password string `json:"password"`
    Token string `json:"token"`
}

type TokenizedBar struct {
    Id string `json:"id"`
	Name string `json:"name"`
    Token string `json:"token"`
}

func MapEntityToTokenized(entity BarEntity) TokenizedBar {
    return TokenizedBar{
		Id: entity.Id, 
		Name: entity.Name, 
		Token: entity.Token,
	}
}
