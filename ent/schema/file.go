package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid",uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("mime").Default("application/octet-stream").MaxLen(127),
		field.Uint("size"),
		field.String("bucket").NotEmpty(),
		field.String("directory").Default("/"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file_tickets",FileTicket.Type).Ref("files").Unique(),
	}
}

func (File) Indexes () []ent.Index {
	return []ent.Index{
		index.Fields("bucket","directory"),
		index.Fields("uuid").Unique(),
		index.Fields("uuid").Edges("file_tickets").Unique(),

	}
}