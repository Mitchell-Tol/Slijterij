package categorymodel

type CategoryEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	BarId string `json:"bar_id"`
	Color string `json:"color"`
}

type Category struct {
	Name string `json:"name"`
	BarId string `json:"bar_id"`
	Color string `json:"color"`
}

type UpdatedCategory struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color"`
}

type CategoryId struct {
	Id string `json:"id"`
}
