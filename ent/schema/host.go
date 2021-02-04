package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique().NotEmpty(),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admins",Admin.Type),
		edge.To("file_tickets",FileTicket.Type),
	}
}
func (Host) Indexes () []ent.Index {
	return []ent.Index{
		index.Fields("name","token").Unique(),
	}
}