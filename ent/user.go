// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/geekshacking/geekhub-backend/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DisplayName holds the value of the "display_name" field.
	DisplayName string `json:"display_name,omitempty"`
	// Auth0ID holds the value of the "auth0_id" field.
	Auth0ID string `json:"auth0_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// ReportedTickets holds the value of the reported_tickets edge.
	ReportedTickets []*Ticket `json:"reported_tickets,omitempty"`
	// AssignedTickets holds the value of the assigned_tickets edge.
	AssignedTickets []*Ticket `json:"assigned_tickets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// ReportedTicketsOrErr returns the ReportedTickets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ReportedTicketsOrErr() ([]*Ticket, error) {
	if e.loadedTypes[1] {
		return e.ReportedTickets, nil
	}
	return nil, &NotLoadedError{edge: "reported_tickets"}
}

// AssignedTicketsOrErr returns the AssignedTickets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssignedTicketsOrErr() ([]*Ticket, error) {
	if e.loadedTypes[2] {
		return e.AssignedTickets, nil
	}
	return nil, &NotLoadedError{edge: "assigned_tickets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldDisplayName, user.FieldAuth0ID:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldDisplayName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field display_name", values[i])
			} else if value.Valid {
				u.DisplayName = value.String
			}
		case user.FieldAuth0ID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field auth0_id", values[i])
			} else if value.Valid {
				u.Auth0ID = value.String
			}
		}
	}
	return nil
}

// QueryProjects queries the "projects" edge of the User entity.
func (u *User) QueryProjects() *ProjectQuery {
	return (&UserClient{config: u.config}).QueryProjects(u)
}

// QueryReportedTickets queries the "reported_tickets" edge of the User entity.
func (u *User) QueryReportedTickets() *TicketQuery {
	return (&UserClient{config: u.config}).QueryReportedTickets(u)
}

// QueryAssignedTickets queries the "assigned_tickets" edge of the User entity.
func (u *User) QueryAssignedTickets() *TicketQuery {
	return (&UserClient{config: u.config}).QueryAssignedTickets(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", display_name=")
	builder.WriteString(u.DisplayName)
	builder.WriteString(", auth0_id=")
	builder.WriteString(u.Auth0ID)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
