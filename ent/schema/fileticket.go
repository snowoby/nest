package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// FileTicket holds the schema definition for the FileTicket entity.
type FileTicket struct {
	ent.Schema
}

// Fields of the FileTicket.
func (FileTicket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid",uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("original_filename").MaxLen(256).Default("file"),
		field.String("identifier").MaxLen(128),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the FileTicket.
func (FileTicket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files",File.Type),
		edge.From("hosts",Host.Type).Ref("file_tickets").Unique(),
	}
}
func (FileTicket) Indexes () []ent.Index {
	return []ent.Index{
		index.Fields("uuid").Unique(),
	}
}