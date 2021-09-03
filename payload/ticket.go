package payload

type TicketResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	TicketType     string `json:"ticketType"`
	TicketStatus   string `json:"ticketStatus"`
	ReporterID     int    `json:"reporterId"`
	AssigneeID     int    `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}
