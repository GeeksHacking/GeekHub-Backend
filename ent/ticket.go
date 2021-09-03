// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/geekshacking/geekhub-backend/ent/project"
	"github.com/geekshacking/geekhub-backend/ent/ticket"
	"github.com/geekshacking/geekhub-backend/ent/user"
)

// Ticket is the model entity for the Ticket schema.
type Ticket struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Type holds the value of the "type" field.
	Type ticket.Type `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status ticket.Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TicketQuery when eager-loading is set.
	Edges                 TicketEdges `json:"edges"`
	project_tickets       *int
	ticket_children       *int
	user_reported_tickets *int
	user_assigned_tickets *int
}

// TicketEdges holds the relations/edges for other nodes in the graph.
type TicketEdges struct {
	// Project holds the value of the project edge.
	Project *Project `json:"project,omitempty"`
	// Reporter holds the value of the reporter edge.
	Reporter *User `json:"reporter,omitempty"`
	// Assignee holds the value of the assignee edge.
	Assignee *User `json:"assignee,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Ticket `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Ticket `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// The edge project was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// ReporterOrErr returns the Reporter value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) ReporterOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Reporter == nil {
			// The edge reporter was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Reporter, nil
	}
	return nil, &NotLoadedError{edge: "reporter"}
}

// AssigneeOrErr returns the Assignee value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) AssigneeOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Assignee == nil {
			// The edge assignee was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Assignee, nil
	}
	return nil, &NotLoadedError{edge: "assignee"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TicketEdges) ParentOrErr() (*Ticket, error) {
	if e.loadedTypes[3] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ticket.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e TicketEdges) ChildrenOrErr() ([]*Ticket, error) {
	if e.loadedTypes[4] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ticket) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ticket.FieldID:
			values[i] = new(sql.NullInt64)
		case ticket.FieldName, ticket.FieldDescription, ticket.FieldType, ticket.FieldStatus:
			values[i] = new(sql.NullString)
		case ticket.ForeignKeys[0]: // project_tickets
			values[i] = new(sql.NullInt64)
		case ticket.ForeignKeys[1]: // ticket_children
			values[i] = new(sql.NullInt64)
		case ticket.ForeignKeys[2]: // user_reported_tickets
			values[i] = new(sql.NullInt64)
		case ticket.ForeignKeys[3]: // user_assigned_tickets
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ticket", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ticket fields.
func (t *Ticket) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ticket.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case ticket.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				t.Name = value.String
			}
		case ticket.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case ticket.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				t.Type = ticket.Type(value.String)
			}
		case ticket.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				t.Status = ticket.Status(value.String)
			}
		case ticket.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field project_tickets", value)
			} else if value.Valid {
				t.project_tickets = new(int)
				*t.project_tickets = int(value.Int64)
			}
		case ticket.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field ticket_children", value)
			} else if value.Valid {
				t.ticket_children = new(int)
				*t.ticket_children = int(value.Int64)
			}
		case ticket.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_reported_tickets", value)
			} else if value.Valid {
				t.user_reported_tickets = new(int)
				*t.user_reported_tickets = int(value.Int64)
			}
		case ticket.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_assigned_tickets", value)
			} else if value.Valid {
				t.user_assigned_tickets = new(int)
				*t.user_assigned_tickets = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryProject queries the "project" edge of the Ticket entity.
func (t *Ticket) QueryProject() *ProjectQuery {
	return (&TicketClient{config: t.config}).QueryProject(t)
}

// QueryReporter queries the "reporter" edge of the Ticket entity.
func (t *Ticket) QueryReporter() *UserQuery {
	return (&TicketClient{config: t.config}).QueryReporter(t)
}

// QueryAssignee queries the "assignee" edge of the Ticket entity.
func (t *Ticket) QueryAssignee() *UserQuery {
	return (&TicketClient{config: t.config}).QueryAssignee(t)
}

// QueryParent queries the "parent" edge of the Ticket entity.
func (t *Ticket) QueryParent() *TicketQuery {
	return (&TicketClient{config: t.config}).QueryParent(t)
}

// QueryChildren queries the "children" edge of the Ticket entity.
func (t *Ticket) QueryChildren() *TicketQuery {
	return (&TicketClient{config: t.config}).QueryChildren(t)
}

// Update returns a builder for updating this Ticket.
// Note that you need to call Ticket.Unwrap() before calling this method if this Ticket
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Ticket) Update() *TicketUpdateOne {
	return (&TicketClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Ticket entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Ticket) Unwrap() *Ticket {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ticket is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Ticket) String() string {
	var builder strings.Builder
	builder.WriteString("Ticket(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", description=")
	builder.WriteString(t.Description)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", t.Type))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", t.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Tickets is a parsable slice of Ticket.
type Tickets []*Ticket

func (t Tickets) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
