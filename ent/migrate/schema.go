// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "host_admins", Type: field.TypeInt, Nullable: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "admins_hosts_admins",
				Columns: []*schema.Column{AdminsColumns[3]},

				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "admin_token_name",
				Unique:  true,
				Columns: []*schema.Column{AdminsColumns[1], AdminsColumns[2]},
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "mime", Type: field.TypeString, Size: 127, Default: "application/octet-stream"},
		{Name: "size", Type: field.TypeUint},
		{Name: "bucket", Type: field.TypeString},
		{Name: "directory", Type: field.TypeString, Default: "/"},
		{Name: "file_ticket_files", Type: field.TypeInt, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "files_file_tickets_files",
				Columns: []*schema.Column{FilesColumns[6]},

				RefColumns: []*schema.Column{FileTicketsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "file_bucket_directory",
				Unique:  false,
				Columns: []*schema.Column{FilesColumns[4], FilesColumns[5]},
			},
			{
				Name:    "file_uuid",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[1]},
			},
			{
				Name:    "file_uuid_file_ticket_files",
				Unique:  true,
				Columns: []*schema.Column{FilesColumns[1], FilesColumns[6]},
			},
		},
	}
	// FileTicketsColumns holds the columns for the "file_tickets" table.
	FileTicketsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "original_filename", Type: field.TypeString, Size: 256, Default: "file"},
		{Name: "identifier", Type: field.TypeString, Size: 128},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "host_file_tickets", Type: field.TypeInt, Nullable: true},
	}
	// FileTicketsTable holds the schema information for the "file_tickets" table.
	FileTicketsTable = &schema.Table{
		Name:       "file_tickets",
		Columns:    FileTicketsColumns,
		PrimaryKey: []*schema.Column{FileTicketsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "file_tickets_hosts_file_tickets",
				Columns: []*schema.Column{FileTicketsColumns[5]},

				RefColumns: []*schema.Column{HostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "fileticket_uuid",
				Unique:  true,
				Columns: []*schema.Column{FileTicketsColumns[1]},
			},
		},
	}
	// HostsColumns holds the columns for the "hosts" table.
	HostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "token", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// HostsTable holds the schema information for the "hosts" table.
	HostsTable = &schema.Table{
		Name:        "hosts",
		Columns:     HostsColumns,
		PrimaryKey:  []*schema.Column{HostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
		Indexes: []*schema.Index{
			{
				Name:    "host_name_token",
				Unique:  true,
				Columns: []*schema.Column{HostsColumns[2], HostsColumns[1]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		FilesTable,
		FileTicketsTable,
		HostsTable,
	}
)

func init() {
	AdminsTable.ForeignKeys[0].RefTable = HostsTable
	FilesTable.ForeignKeys[0].RefTable = FileTicketsTable
	FileTicketsTable.ForeignKeys[0].RefTable = HostsTable
}
