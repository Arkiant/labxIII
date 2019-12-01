package dialogflow

type SearchResponse struct {
	FulfillmentText string          `json:"fulfillmentText"`
	OutputContexts  []OutputContext `json:"outputContexts"`
}
