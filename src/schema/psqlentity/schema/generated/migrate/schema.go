// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "country", Type: field.TypeString},
		{Name: "province", Type: field.TypeString},
		{Name: "regency", Type: field.TypeString},
		{Name: "sub_district", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUUID, Unique: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "is_deleted", Type: field.TypeInt64, Default: 0},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUUID, Unique: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "is_deleted", Type: field.TypeInt64, Default: 0},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "hobbies", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUUID, Unique: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "is_deleted", Type: field.TypeInt64, Default: 0},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		RolesTable,
		UsersTable,
	}
)

func init() {
}
