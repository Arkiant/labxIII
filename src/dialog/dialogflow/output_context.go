package dialogflow

type Parameters struct {
	HotelName string `json:"hotelName"`
	Price     string `json:"price"`
	OptionID  string `json:"optionID"`
}

type OutputContext struct {
	Name          string     `json:"name"`
	LifespanCount int        `json:"lifespanCount"`
	Parameters    Parameters `json:"parameters"`
}
