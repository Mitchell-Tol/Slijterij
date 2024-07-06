package categorymodel

type CategoryEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	BarId string `json:"bar_id"`
}

type Category struct {
	Name string `json:"name"`
	BarId string `json:"bar_id"`
}

type UpdatedCategory struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type CategoryId struct {
	Id string `json:"id"`
}
