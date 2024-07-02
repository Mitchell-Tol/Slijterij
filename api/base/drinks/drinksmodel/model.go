package drinksmodel

type DrinkEntity struct {
    Id string `json:"id"`
    Name string `json:"name"`
    BarId string `json:"bar_id"`
    StartPrice float32 `json:"start_price"`
    CurrentPrice float32 `json:"current_price"`
    Multiplier float32 `json:"multiplier"`
	Tag string `json:"tag"`
	CategoryId string `json:"category_id"`
}

type DrinkId struct {
	Id string `json:"id"`
}
