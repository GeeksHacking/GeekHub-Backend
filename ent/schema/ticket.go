package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Ticket holds the schema definition for the Ticket entity.
type Ticket struct {
	ent.Schema
}

// Fields of the Ticket.
func (Ticket) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Enum("type").
			Values("epic", "story", "bug", "task"),
		field.Enum("status").
			Values("backlog", "development", "qa", "release"),
	}
}

// Edges of the Ticket.
func (Ticket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).Ref("tickets").Unique(),
		edge.From("reporter", User.Type).Ref("reported_tickets").Unique(),
		edge.From("assignee", User.Type).Ref("assigned_tickets").Unique(),
		edge.To("children", Ticket.Type).From("parent").Unique(),
	}
}
