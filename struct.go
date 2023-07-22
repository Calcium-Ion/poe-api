package poe_api

type SubscriptionsMutation struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	SubscriptionName string      `json:"subscriptionName"`
	Query            interface{} `json:"query"`
	QueryHash        string      `json:"queryHash"`
}

type Payload struct {
	QueryName  string                 `json:"queryName"`
	Variables  map[string]interface{} `json:"variables"`
	Extensions map[string]interface{} `json:"extensions"`
}
