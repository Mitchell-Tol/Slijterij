package drinksmodel

type DrinkEntity struct {
    Id string `json:"id"`
    Name string `json:"name"`
    BarId string `json:"bar_id"`
    StartPrice float32 `json:"start_price"`
    CurrentPrice float32 `json:"current_price"`
    RiseMultiplier float32 `json:"rise_multiplier"`
	Tag string `json:"tag"`
	CategoryId string `json:"category_id"`
	DropMultiplier float32 `json:"drop_multiplier"`
}

type Drink struct {
    Name string `json:"name"`
    BarId string `json:"bar_id"`
    StartPrice float32 `json:"start_price"`
    CurrentPrice float32 `json:"current_price"`
    RiseMultiplier float32 `json:"rise_multiplier"`
	Tag string `json:"tag"`
	CategoryId string `json:"category_id"`
	DropMultiplier float32 `json:"drop_multiplier"`
}

type DrinkId struct {
	Id string `json:"id"`
}

type DrinkByBar struct {
	BarId string `json:"bar_id"`
}
