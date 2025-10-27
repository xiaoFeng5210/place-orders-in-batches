package service

type Product struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
}

type OrderRequest struct {
	Dealer   string    `json:"dealer"`
	Products []Product `json:"products"`
	Amount   float64   `json:"amount"`
	PayType  string    `json:"pay_type"`
}

type OrderRespData struct {
	Id string `json:"_id"`
}

type OrderResponse struct {
	Code int           `json:"code"`
	Data OrderRespData `json:"data"`
}
