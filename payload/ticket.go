package payload

import validation "github.com/go-ozzo/ozzo-validation"

type TicketResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"ticketType"`
	Status         string `json:"ticketStatus"`
	ReporterID     string `json:"reporterId"`
	AssigneeID     string `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}

type CreateTicketRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"ticketType"`
	Status         string `json:"ticketStatus"`
	ReporterID     string `json:"reporterId"`
	AssigneeID     string `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}

func (c *CreateTicketRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 0)))
}
