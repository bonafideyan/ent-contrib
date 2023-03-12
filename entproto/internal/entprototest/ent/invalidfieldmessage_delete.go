// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entproto/internal/entprototest/ent/invalidfieldmessage"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InvalidFieldMessageDelete is the builder for deleting a InvalidFieldMessage entity.
type InvalidFieldMessageDelete struct {
	config
	hooks    []Hook
	mutation *InvalidFieldMessageMutation
}

// Where appends a list predicates to the InvalidFieldMessageDelete builder.
func (ifmd *InvalidFieldMessageDelete) Where(ps ...predicate.InvalidFieldMessage) *InvalidFieldMessageDelete {
	ifmd.mutation.Where(ps...)
	return ifmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ifmd *InvalidFieldMessageDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, InvalidFieldMessageMutation](ctx, ifmd.sqlExec, ifmd.mutation, ifmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmd *InvalidFieldMessageDelete) ExecX(ctx context.Context) int {
	n, err := ifmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ifmd *InvalidFieldMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: invalidfieldmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: invalidfieldmessage.FieldID,
			},
		},
	}
	if ps := ifmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ifmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ifmd.mutation.done = true
	return affected, err
}

// InvalidFieldMessageDeleteOne is the builder for deleting a single InvalidFieldMessage entity.
type InvalidFieldMessageDeleteOne struct {
	ifmd *InvalidFieldMessageDelete
}

// Exec executes the deletion query.
func (ifmdo *InvalidFieldMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := ifmdo.ifmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invalidfieldmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ifmdo *InvalidFieldMessageDeleteOne) ExecX(ctx context.Context) {
	ifmdo.ifmd.ExecX(ctx)
}
