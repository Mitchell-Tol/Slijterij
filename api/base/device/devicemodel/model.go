package devicemodel

type Device struct {
    BarId string `json:"bar_id"`
    Name string `json:"name"`
}

type DeviceEntity struct {
    Id string `json:"id"`
    BarId string `json:"bar_id"`
    Name string `json:"name"`
}

type UpdatedDevice struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type DeviceId struct {
	Id string `json:"id"`
}
