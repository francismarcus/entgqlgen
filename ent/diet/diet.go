// Code generated by entc, DO NOT EDIT.

package diet

import (
	"time"
)

const (
	// Label holds the string label denoting the diet type in the database.
	Label = "diet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGoalWeight holds the string denoting the goal_weight field in the database.
	FieldGoalWeight = "goal_weight"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"

	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"

	// Table holds the table name of the diet in the database.
	Table = "diets"
	// AuthorTable is the table the holds the author relation/edge.
	AuthorTable = "diets"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_diets"
)

// Columns holds all SQL columns for diet fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldGoalWeight,
	FieldLength,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Diet type.
var ForeignKeys = []string{
	"user_diets",
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
	// DefaultCreatedAt holds the default value on creation for the created_at field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the updated_at field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	UpdateDefaultUpdatedAt func() time.Time
)
