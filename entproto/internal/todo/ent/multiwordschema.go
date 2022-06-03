// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/contrib/entproto/internal/todo/ent/multiwordschema"
	"entgo.io/ent/dialect/sql"
)

// MultiWordSchema is the model entity for the MultiWordSchema schema.
type MultiWordSchema struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit multiwordschema.Unit `json:"unit,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MultiWordSchema) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case multiwordschema.FieldID:
			values[i] = new(sql.NullInt64)
		case multiwordschema.FieldUnit:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MultiWordSchema", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MultiWordSchema fields.
func (mws *MultiWordSchema) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case multiwordschema.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mws.ID = int(value.Int64)
		case multiwordschema.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				mws.Unit = multiwordschema.Unit(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MultiWordSchema.
// Note that you need to call MultiWordSchema.Unwrap() before calling this method if this MultiWordSchema
// was returned from a transaction, and the transaction was committed or rolled back.
func (mws *MultiWordSchema) Update() *MultiWordSchemaUpdateOne {
	return (&MultiWordSchemaClient{config: mws.config}).UpdateOne(mws)
}

// Unwrap unwraps the MultiWordSchema entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mws *MultiWordSchema) Unwrap() *MultiWordSchema {
	tx, ok := mws.config.driver.(*txDriver)
	if !ok {
		panic("ent: MultiWordSchema is not a transactional entity")
	}
	mws.config.driver = tx.drv
	return mws
}

// String implements the fmt.Stringer.
func (mws *MultiWordSchema) String() string {
	var builder strings.Builder
	builder.WriteString("MultiWordSchema(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mws.ID))
	builder.WriteString("unit=")
	builder.WriteString(fmt.Sprintf("%v", mws.Unit))
	builder.WriteByte(')')
	return builder.String()
}

// MultiWordSchemas is a parsable slice of MultiWordSchema.
type MultiWordSchemas []*MultiWordSchema

func (mws MultiWordSchemas) config(cfg config) {
	for _i := range mws {
		mws[_i].config = cfg
	}
}
