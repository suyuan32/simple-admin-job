// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTaskGroup holds the string denoting the task_group field in the database.
	FieldTaskGroup = "task_group"
	// FieldCronExpression holds the string denoting the cron_expression field in the database.
	FieldCronExpression = "cron_expression"
	// FieldPattern holds the string denoting the pattern field in the database.
	FieldPattern = "pattern"
	// FieldPayload holds the string denoting the payload field in the database.
	FieldPayload = "payload"
	// FieldEntryID holds the string denoting the entry_id field in the database.
	FieldEntryID = "entry_id"
	// Table holds the table name of the task in the database.
	Table = "sys_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldTaskGroup,
	FieldCronExpression,
	FieldPattern,
	FieldPayload,
	FieldEntryID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
)