package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	ent.Schema
}

// Fields of the Admin.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
		field.String("name").Unique().NotEmpty(),
	}
}

// Edges of the Admin.
func (Admin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("host",Host.Type).Ref("admins").Unique(),
	}
}

func (Admin) Indexes () []ent.Index  {
	return []ent.Index{
		index.Fields("token","name").Unique(),
	}
}