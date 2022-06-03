// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithfieldone"
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithFieldOneDelete is the builder for deleting a MessageWithFieldOne entity.
type MessageWithFieldOneDelete struct {
	config
	hooks    []Hook
	mutation *MessageWithFieldOneMutation
}

// Where appends a list predicates to the MessageWithFieldOneDelete builder.
func (mwfod *MessageWithFieldOneDelete) Where(ps ...predicate.MessageWithFieldOne) *MessageWithFieldOneDelete {
	mwfod.mutation.Where(ps...)
	return mwfod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mwfod *MessageWithFieldOneDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mwfod.hooks) == 0 {
		affected, err = mwfod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithFieldOneMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mwfod.mutation = mutation
			affected, err = mwfod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mwfod.hooks) - 1; i >= 0; i-- {
			if mwfod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwfod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mwfod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfod *MessageWithFieldOneDelete) ExecX(ctx context.Context) int {
	n, err := mwfod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mwfod *MessageWithFieldOneDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: messagewithfieldone.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: messagewithfieldone.FieldID,
			},
		},
	}
	if ps := mwfod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mwfod.driver, _spec)
}

// MessageWithFieldOneDeleteOne is the builder for deleting a single MessageWithFieldOne entity.
type MessageWithFieldOneDeleteOne struct {
	mwfod *MessageWithFieldOneDelete
}

// Exec executes the deletion query.
func (mwfodo *MessageWithFieldOneDeleteOne) Exec(ctx context.Context) error {
	n, err := mwfodo.mwfod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{messagewithfieldone.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mwfodo *MessageWithFieldOneDeleteOne) ExecX(ctx context.Context) {
	mwfodo.mwfod.ExecX(ctx)
}
