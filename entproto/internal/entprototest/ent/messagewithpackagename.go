// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithpackagename"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MessageWithPackageName is the model entity for the MessageWithPackageName schema.
type MessageWithPackageName struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageWithPackageName) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messagewithpackagename.FieldID:
			values[i] = new(sql.NullInt64)
		case messagewithpackagename.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageWithPackageName fields.
func (mwpn *MessageWithPackageName) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messagewithpackagename.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mwpn.ID = int(value.Int64)
		case messagewithpackagename.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mwpn.Name = value.String
			}
		default:
			mwpn.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MessageWithPackageName.
// This includes values selected through modifiers, order, etc.
func (mwpn *MessageWithPackageName) Value(name string) (ent.Value, error) {
	return mwpn.selectValues.Get(name)
}

// Update returns a builder for updating this MessageWithPackageName.
// Note that you need to call MessageWithPackageName.Unwrap() before calling this method if this MessageWithPackageName
// was returned from a transaction, and the transaction was committed or rolled back.
func (mwpn *MessageWithPackageName) Update() *MessageWithPackageNameUpdateOne {
	return NewMessageWithPackageNameClient(mwpn.config).UpdateOne(mwpn)
}

// Unwrap unwraps the MessageWithPackageName entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mwpn *MessageWithPackageName) Unwrap() *MessageWithPackageName {
	_tx, ok := mwpn.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageWithPackageName is not a transactional entity")
	}
	mwpn.config.driver = _tx.drv
	return mwpn
}

// String implements the fmt.Stringer.
func (mwpn *MessageWithPackageName) String() string {
	var builder strings.Builder
	builder.WriteString("MessageWithPackageName(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mwpn.ID))
	builder.WriteString("name=")
	builder.WriteString(mwpn.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MessageWithPackageNames is a parsable slice of MessageWithPackageName.
type MessageWithPackageNames []*MessageWithPackageName
