package handler

type DeleteDealerConfigRequestParams struct {
	DealerId string `json:"dealer_id" binding:"required"`
}
