package barmodel

type Bar struct {
	Name string `json:"name"`
    Password string `json:"password"`
	SuperAdmin int `json:"super_admin"`
}

type LoginBar struct {
	Name string `json:"name"`
    Password string `json:"password"`
}

type BarEntity struct {
    Id string `json:"id"`
	Name string `json:"name"`
    Password string `json:"password"`
    Token string `json:"token"`
	SuperAdmin int `json:"super_admin"`
}

type TokenizedBar struct {
    Id string `json:"id"`
	Name string `json:"name"`
    Token string `json:"token"`
	SuperAdmin int `json:"super_admin"`
}

type BarId struct {
	Id string `json:"id"`
}
