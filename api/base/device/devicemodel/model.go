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

type DeviceByBar struct {
	BarId string `json:"bar_id"`
}
