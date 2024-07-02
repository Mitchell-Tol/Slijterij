package categorymodel

type CategoryByBar struct {
	BarId string `json:"bar_id"`
}

type CategoryEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	BarId string `json:"BarId"`
}
