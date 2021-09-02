package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("repository"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("projects"),

		edge.To("tags", Tag.Type),
		edge.To("tickets", Ticket.Type),
		edge.To("languages", Language.Type),
		edge.To("owner", User.Type).Unique(),
	}
}
