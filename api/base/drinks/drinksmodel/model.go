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
    LastChange float32 `json:"last_change"`
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
