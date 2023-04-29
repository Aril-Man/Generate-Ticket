package response

type Response struct {
	Status bool `json:"status"`
	Code   int  `json:"code"`
	Data   Data `json:"data,omitempty"`
}

type Data struct {
	TicketUrl string `json:"ticket_url"`
}
