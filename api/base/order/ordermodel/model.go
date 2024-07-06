package ordermodel

type OrderEntity struct {
	Id string `json:"id"`
	DeviceId string `json:"device_id"`
	ProductId string `json:"product_id"`
	Timestamp string `json:"timestamp"`
	Amount int `json:"amount"`
	PricePerProduct float32 `json:"price_per_product"`
}

type Order struct {
	DeviceId string `json:"device_id"`
	ProductId string `json:"product_id"`
	Timestamp string `json:"timestamp"`
	Amount int `json:"amount"`
	PricePerProduct float32 `json:"price_per_product"`
}

type UpdatedOrder struct {
	Id string `json:"id"`
	DeviceId string `json:"device_id"`
	ProductId string `json:"product_id"`
	Amount int `json:"amount"`
	PricePerProduct float32 `json:"price_per_product"`
}

type OrderId struct {
	Id string `json:"id"`
}
