package payload

import validation "github.com/go-ozzo/ozzo-validation"

type TicketResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	ReporterID     string `json:"reporterId"`
	AssigneeID     string `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}

type CreateTicketRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	ReporterID     string `json:"reporterId"`
	AssigneeID     string `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}

func (c *CreateTicketRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 0)))
}

type UpdateTicketRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	ReporterID     string `json:"reporterId"`
	AssigneeID     string `json:"assigneeId"`
	ParentTicketID int    `json:"parentTicketId"`
}

func (c *UpdateTicketRequest) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 0)))
}
