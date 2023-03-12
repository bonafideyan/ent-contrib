// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/validmessage"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ValidMessageUpdate is the builder for updating ValidMessage entities.
type ValidMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ValidMessageMutation
}

// Where appends a list predicates to the ValidMessageUpdate builder.
func (vmu *ValidMessageUpdate) Where(ps ...predicate.ValidMessage) *ValidMessageUpdate {
	vmu.mutation.Where(ps...)
	return vmu
}

// SetName sets the "name" field.
func (vmu *ValidMessageUpdate) SetName(s string) *ValidMessageUpdate {
	vmu.mutation.SetName(s)
	return vmu
}

// SetTs sets the "ts" field.
func (vmu *ValidMessageUpdate) SetTs(t time.Time) *ValidMessageUpdate {
	vmu.mutation.SetTs(t)
	return vmu
}

// SetUUID sets the "uuid" field.
func (vmu *ValidMessageUpdate) SetUUID(u uuid.UUID) *ValidMessageUpdate {
	vmu.mutation.SetUUID(u)
	return vmu
}

// SetU8 sets the "u8" field.
func (vmu *ValidMessageUpdate) SetU8(u uint8) *ValidMessageUpdate {
	vmu.mutation.ResetU8()
	vmu.mutation.SetU8(u)
	return vmu
}

// AddU8 adds u to the "u8" field.
func (vmu *ValidMessageUpdate) AddU8(u int8) *ValidMessageUpdate {
	vmu.mutation.AddU8(u)
	return vmu
}

// SetOpti8 sets the "opti8" field.
func (vmu *ValidMessageUpdate) SetOpti8(i int8) *ValidMessageUpdate {
	vmu.mutation.ResetOpti8()
	vmu.mutation.SetOpti8(i)
	return vmu
}

// SetNillableOpti8 sets the "opti8" field if the given value is not nil.
func (vmu *ValidMessageUpdate) SetNillableOpti8(i *int8) *ValidMessageUpdate {
	if i != nil {
		vmu.SetOpti8(*i)
	}
	return vmu
}

// AddOpti8 adds i to the "opti8" field.
func (vmu *ValidMessageUpdate) AddOpti8(i int8) *ValidMessageUpdate {
	vmu.mutation.AddOpti8(i)
	return vmu
}

// ClearOpti8 clears the value of the "opti8" field.
func (vmu *ValidMessageUpdate) ClearOpti8() *ValidMessageUpdate {
	vmu.mutation.ClearOpti8()
	return vmu
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmu *ValidMessageUpdate) Mutation() *ValidMessageMutation {
	return vmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vmu *ValidMessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ValidMessageMutation](ctx, vmu.sqlSave, vmu.mutation, vmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vmu *ValidMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := vmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vmu *ValidMessageUpdate) Exec(ctx context.Context) error {
	_, err := vmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmu *ValidMessageUpdate) ExecX(ctx context.Context) {
	if err := vmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vmu *ValidMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validmessage.Table,
			Columns: validmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		},
	}
	if ps := vmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmu.mutation.Name(); ok {
		_spec.SetField(validmessage.FieldName, field.TypeString, value)
	}
	if value, ok := vmu.mutation.Ts(); ok {
		_spec.SetField(validmessage.FieldTs, field.TypeTime, value)
	}
	if value, ok := vmu.mutation.UUID(); ok {
		_spec.SetField(validmessage.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := vmu.mutation.U8(); ok {
		_spec.SetField(validmessage.FieldU8, field.TypeUint8, value)
	}
	if value, ok := vmu.mutation.AddedU8(); ok {
		_spec.AddField(validmessage.FieldU8, field.TypeUint8, value)
	}
	if value, ok := vmu.mutation.Opti8(); ok {
		_spec.SetField(validmessage.FieldOpti8, field.TypeInt8, value)
	}
	if value, ok := vmu.mutation.AddedOpti8(); ok {
		_spec.AddField(validmessage.FieldOpti8, field.TypeInt8, value)
	}
	if vmu.mutation.Opti8Cleared() {
		_spec.ClearField(validmessage.FieldOpti8, field.TypeInt8)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{validmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vmu.mutation.done = true
	return n, nil
}

// ValidMessageUpdateOne is the builder for updating a single ValidMessage entity.
type ValidMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ValidMessageMutation
}

// SetName sets the "name" field.
func (vmuo *ValidMessageUpdateOne) SetName(s string) *ValidMessageUpdateOne {
	vmuo.mutation.SetName(s)
	return vmuo
}

// SetTs sets the "ts" field.
func (vmuo *ValidMessageUpdateOne) SetTs(t time.Time) *ValidMessageUpdateOne {
	vmuo.mutation.SetTs(t)
	return vmuo
}

// SetUUID sets the "uuid" field.
func (vmuo *ValidMessageUpdateOne) SetUUID(u uuid.UUID) *ValidMessageUpdateOne {
	vmuo.mutation.SetUUID(u)
	return vmuo
}

// SetU8 sets the "u8" field.
func (vmuo *ValidMessageUpdateOne) SetU8(u uint8) *ValidMessageUpdateOne {
	vmuo.mutation.ResetU8()
	vmuo.mutation.SetU8(u)
	return vmuo
}

// AddU8 adds u to the "u8" field.
func (vmuo *ValidMessageUpdateOne) AddU8(u int8) *ValidMessageUpdateOne {
	vmuo.mutation.AddU8(u)
	return vmuo
}

// SetOpti8 sets the "opti8" field.
func (vmuo *ValidMessageUpdateOne) SetOpti8(i int8) *ValidMessageUpdateOne {
	vmuo.mutation.ResetOpti8()
	vmuo.mutation.SetOpti8(i)
	return vmuo
}

// SetNillableOpti8 sets the "opti8" field if the given value is not nil.
func (vmuo *ValidMessageUpdateOne) SetNillableOpti8(i *int8) *ValidMessageUpdateOne {
	if i != nil {
		vmuo.SetOpti8(*i)
	}
	return vmuo
}

// AddOpti8 adds i to the "opti8" field.
func (vmuo *ValidMessageUpdateOne) AddOpti8(i int8) *ValidMessageUpdateOne {
	vmuo.mutation.AddOpti8(i)
	return vmuo
}

// ClearOpti8 clears the value of the "opti8" field.
func (vmuo *ValidMessageUpdateOne) ClearOpti8() *ValidMessageUpdateOne {
	vmuo.mutation.ClearOpti8()
	return vmuo
}

// Mutation returns the ValidMessageMutation object of the builder.
func (vmuo *ValidMessageUpdateOne) Mutation() *ValidMessageMutation {
	return vmuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vmuo *ValidMessageUpdateOne) Select(field string, fields ...string) *ValidMessageUpdateOne {
	vmuo.fields = append([]string{field}, fields...)
	return vmuo
}

// Save executes the query and returns the updated ValidMessage entity.
func (vmuo *ValidMessageUpdateOne) Save(ctx context.Context) (*ValidMessage, error) {
	return withHooks[*ValidMessage, ValidMessageMutation](ctx, vmuo.sqlSave, vmuo.mutation, vmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vmuo *ValidMessageUpdateOne) SaveX(ctx context.Context) *ValidMessage {
	node, err := vmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vmuo *ValidMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := vmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vmuo *ValidMessageUpdateOne) ExecX(ctx context.Context) {
	if err := vmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vmuo *ValidMessageUpdateOne) sqlSave(ctx context.Context) (_node *ValidMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   validmessage.Table,
			Columns: validmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: validmessage.FieldID,
			},
		},
	}
	id, ok := vmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ValidMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, validmessage.FieldID)
		for _, f := range fields {
			if !validmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != validmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vmuo.mutation.Name(); ok {
		_spec.SetField(validmessage.FieldName, field.TypeString, value)
	}
	if value, ok := vmuo.mutation.Ts(); ok {
		_spec.SetField(validmessage.FieldTs, field.TypeTime, value)
	}
	if value, ok := vmuo.mutation.UUID(); ok {
		_spec.SetField(validmessage.FieldUUID, field.TypeUUID, value)
	}
	if value, ok := vmuo.mutation.U8(); ok {
		_spec.SetField(validmessage.FieldU8, field.TypeUint8, value)
	}
	if value, ok := vmuo.mutation.AddedU8(); ok {
		_spec.AddField(validmessage.FieldU8, field.TypeUint8, value)
	}
	if value, ok := vmuo.mutation.Opti8(); ok {
		_spec.SetField(validmessage.FieldOpti8, field.TypeInt8, value)
	}
	if value, ok := vmuo.mutation.AddedOpti8(); ok {
		_spec.AddField(validmessage.FieldOpti8, field.TypeInt8, value)
	}
	if vmuo.mutation.Opti8Cleared() {
		_spec.ClearField(validmessage.FieldOpti8, field.TypeInt8)
	}
	_node = &ValidMessage{config: vmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{validmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vmuo.mutation.done = true
	return _node, nil
}
