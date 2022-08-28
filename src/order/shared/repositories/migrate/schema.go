// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "order_details", Type: field.TypeJSON},
		{Name: "preferred_exchanges", Type: field.TypeJSON},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// OutboxesColumns holds the columns for the "outboxes" table.
	OutboxesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "topic", Type: field.TypeString},
		{Name: "key", Type: field.TypeString},
		{Name: "payload", Type: field.TypeBytes},
		{Name: "headers", Type: field.TypeJSON},
		{Name: "retry_count", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PENDING", "SUCCEEDED", "FAILED"}},
		{Name: "last_retry", Type: field.TypeTime, Nullable: true},
		{Name: "processing_errors", Type: field.TypeJSON, Nullable: true},
	}
	// OutboxesTable holds the schema information for the "outboxes" table.
	OutboxesTable = &schema.Table{
		Name:       "outboxes",
		Columns:    OutboxesColumns,
		PrimaryKey: []*schema.Column{OutboxesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "outbox_last_retry",
				Unique:  false,
				Columns: []*schema.Column{OutboxesColumns[8]},
			},
			{
				Name:    "outbox_status",
				Unique:  false,
				Columns: []*schema.Column{OutboxesColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
		OutboxesTable,
	}
)

func init() {
}
