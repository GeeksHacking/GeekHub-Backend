// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/geekshacking/geekhub-backend/ent/language"
)

// Language is the model entity for the Language schema.
type Language struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LanguageQuery when eager-loading is set.
	Edges LanguageEdges `json:"edges"`
}

// LanguageEdges holds the relations/edges for other nodes in the graph.
type LanguageEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e LanguageEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Language) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case language.FieldID:
			values[i] = new(sql.NullInt64)
		case language.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Language", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Language fields.
func (l *Language) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case language.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		case language.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				l.Name = value.String
			}
		}
	}
	return nil
}

// QueryProjects queries the "projects" edge of the Language entity.
func (l *Language) QueryProjects() *ProjectQuery {
	return (&LanguageClient{config: l.config}).QueryProjects(l)
}

// Update returns a builder for updating this Language.
// Note that you need to call Language.Unwrap() before calling this method if this Language
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Language) Update() *LanguageUpdateOne {
	return (&LanguageClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Language entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Language) Unwrap() *Language {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Language is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Language) String() string {
	var builder strings.Builder
	builder.WriteString("Language(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", name=")
	builder.WriteString(l.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Languages is a parsable slice of Language.
type Languages []*Language

func (l Languages) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
