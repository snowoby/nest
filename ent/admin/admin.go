// Code generated by entc, DO NOT EDIT.

package admin

const (
	// Label holds the string label denoting the admin type in the database.
	Label = "admin"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"

	// Table holds the table name of the admin in the database.
	Table = "admins"
	// HostTable is the table the holds the host relation/edge.
	HostTable = "admins"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "host_admins"
)

// Columns holds all SQL columns for admin fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Admin type.
var ForeignKeys = []string{
	"host_admins",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
