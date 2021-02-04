// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"nest/ent/file"
	"nest/ent/fileticket"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/google/uuid"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
	// Mime holds the value of the "mime" field.
	Mime string `json:"mime,omitempty"`
	// Size holds the value of the "size" field.
	Size uint `json:"size,omitempty"`
	// Bucket holds the value of the "bucket" field.
	Bucket string `json:"bucket,omitempty"`
	// Directory holds the value of the "directory" field.
	Directory string `json:"directory,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges             FileEdges `json:"edges"`
	file_ticket_files *int
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// FileTickets holds the value of the file_tickets edge.
	FileTickets *FileTicket
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// FileTicketsOrErr returns the FileTickets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) FileTicketsOrErr() (*FileTicket, error) {
	if e.loadedTypes[0] {
		if e.FileTickets == nil {
			// The edge file_tickets was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: fileticket.Label}
		}
		return e.FileTickets, nil
	}
	return nil, &NotLoadedError{edge: "file_tickets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*File) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldID, file.FieldSize:
			values[i] = &sql.NullInt64{}
		case file.FieldMime, file.FieldBucket, file.FieldDirectory:
			values[i] = &sql.NullString{}
		case file.FieldUUID:
			values[i] = &uuid.UUID{}
		case file.ForeignKeys[0]: // file_ticket_files
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case file.FieldUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value != nil {
				f.UUID = *value
			}
		case file.FieldMime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime", values[i])
			} else if value.Valid {
				f.Mime = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = uint(value.Int64)
			}
		case file.FieldBucket:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bucket", values[i])
			} else if value.Valid {
				f.Bucket = value.String
			}
		case file.FieldDirectory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field directory", values[i])
			} else if value.Valid {
				f.Directory = value.String
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field file_ticket_files", value)
			} else if value.Valid {
				f.file_ticket_files = new(int)
				*f.file_ticket_files = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryFileTickets queries the "file_tickets" edge of the File entity.
func (f *File) QueryFileTickets() *FileTicketQuery {
	return (&FileClient{config: f.config}).QueryFileTickets(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", f.UUID))
	builder.WriteString(", mime=")
	builder.WriteString(f.Mime)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", bucket=")
	builder.WriteString(f.Bucket)
	builder.WriteString(", directory=")
	builder.WriteString(f.Directory)
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
