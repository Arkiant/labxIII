package routes

type SearchResponse struct {
	FulfillmentText string `json:"fulfillmentText"`

	OutputContexts []OutputContext `json:"outputContexts"`
}

type OutputContext struct {
	Name          string `json:"name"`
	LifespanCount int    `json:"lifespanCount"`
	Parameters    struct {
		HotelName string `json:"hotelName"`
		Price     string `json:"price"`
		OptionID  string `json:"optionID"`
	} `json:"parameters"`
}
