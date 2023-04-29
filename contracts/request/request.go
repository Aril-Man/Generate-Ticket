package request

type RequestTicket struct {
	Name         string `json:"name"`
	Nik          string `json:"nik"`
	DateBought   string `json:"date_bought"`
	InvoiceCode  string `json:"invoice_code"`
	SeatRow      string `json:"seat_row"`
	Category     string `json:"category"`
	TicketNumber string `json:"ticket_number"`
}
