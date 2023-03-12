// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/contrib/entproto/internal/entprototest/ent/skipedgeexample"
	"entgo.io/contrib/entproto/internal/entprototest/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SkipEdgeExampleUpdate is the builder for updating SkipEdgeExample entities.
type SkipEdgeExampleUpdate struct {
	config
	hooks    []Hook
	mutation *SkipEdgeExampleMutation
}

// Where appends a list predicates to the SkipEdgeExampleUpdate builder.
func (seeu *SkipEdgeExampleUpdate) Where(ps ...predicate.SkipEdgeExample) *SkipEdgeExampleUpdate {
	seeu.mutation.Where(ps...)
	return seeu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (seeu *SkipEdgeExampleUpdate) SetUserID(id int) *SkipEdgeExampleUpdate {
	seeu.mutation.SetUserID(id)
	return seeu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (seeu *SkipEdgeExampleUpdate) SetNillableUserID(id *int) *SkipEdgeExampleUpdate {
	if id != nil {
		seeu = seeu.SetUserID(*id)
	}
	return seeu
}

// SetUser sets the "user" edge to the User entity.
func (seeu *SkipEdgeExampleUpdate) SetUser(u *User) *SkipEdgeExampleUpdate {
	return seeu.SetUserID(u.ID)
}

// Mutation returns the SkipEdgeExampleMutation object of the builder.
func (seeu *SkipEdgeExampleUpdate) Mutation() *SkipEdgeExampleMutation {
	return seeu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (seeu *SkipEdgeExampleUpdate) ClearUser() *SkipEdgeExampleUpdate {
	seeu.mutation.ClearUser()
	return seeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (seeu *SkipEdgeExampleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, SkipEdgeExampleMutation](ctx, seeu.sqlSave, seeu.mutation, seeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seeu *SkipEdgeExampleUpdate) SaveX(ctx context.Context) int {
	affected, err := seeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (seeu *SkipEdgeExampleUpdate) Exec(ctx context.Context) error {
	_, err := seeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seeu *SkipEdgeExampleUpdate) ExecX(ctx context.Context) {
	if err := seeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seeu *SkipEdgeExampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skipedgeexample.Table,
			Columns: skipedgeexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skipedgeexample.FieldID,
			},
		},
	}
	if ps := seeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if seeu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   skipedgeexample.UserTable,
			Columns: []string{skipedgeexample.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seeu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   skipedgeexample.UserTable,
			Columns: []string{skipedgeexample.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, seeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skipedgeexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	seeu.mutation.done = true
	return n, nil
}

// SkipEdgeExampleUpdateOne is the builder for updating a single SkipEdgeExample entity.
type SkipEdgeExampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SkipEdgeExampleMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (seeuo *SkipEdgeExampleUpdateOne) SetUserID(id int) *SkipEdgeExampleUpdateOne {
	seeuo.mutation.SetUserID(id)
	return seeuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (seeuo *SkipEdgeExampleUpdateOne) SetNillableUserID(id *int) *SkipEdgeExampleUpdateOne {
	if id != nil {
		seeuo = seeuo.SetUserID(*id)
	}
	return seeuo
}

// SetUser sets the "user" edge to the User entity.
func (seeuo *SkipEdgeExampleUpdateOne) SetUser(u *User) *SkipEdgeExampleUpdateOne {
	return seeuo.SetUserID(u.ID)
}

// Mutation returns the SkipEdgeExampleMutation object of the builder.
func (seeuo *SkipEdgeExampleUpdateOne) Mutation() *SkipEdgeExampleMutation {
	return seeuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (seeuo *SkipEdgeExampleUpdateOne) ClearUser() *SkipEdgeExampleUpdateOne {
	seeuo.mutation.ClearUser()
	return seeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (seeuo *SkipEdgeExampleUpdateOne) Select(field string, fields ...string) *SkipEdgeExampleUpdateOne {
	seeuo.fields = append([]string{field}, fields...)
	return seeuo
}

// Save executes the query and returns the updated SkipEdgeExample entity.
func (seeuo *SkipEdgeExampleUpdateOne) Save(ctx context.Context) (*SkipEdgeExample, error) {
	return withHooks[*SkipEdgeExample, SkipEdgeExampleMutation](ctx, seeuo.sqlSave, seeuo.mutation, seeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (seeuo *SkipEdgeExampleUpdateOne) SaveX(ctx context.Context) *SkipEdgeExample {
	node, err := seeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (seeuo *SkipEdgeExampleUpdateOne) Exec(ctx context.Context) error {
	_, err := seeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (seeuo *SkipEdgeExampleUpdateOne) ExecX(ctx context.Context) {
	if err := seeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (seeuo *SkipEdgeExampleUpdateOne) sqlSave(ctx context.Context) (_node *SkipEdgeExample, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   skipedgeexample.Table,
			Columns: skipedgeexample.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: skipedgeexample.FieldID,
			},
		},
	}
	id, ok := seeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SkipEdgeExample.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := seeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, skipedgeexample.FieldID)
		for _, f := range fields {
			if !skipedgeexample.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != skipedgeexample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := seeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if seeuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   skipedgeexample.UserTable,
			Columns: []string{skipedgeexample.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := seeuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   skipedgeexample.UserTable,
			Columns: []string{skipedgeexample.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SkipEdgeExample{config: seeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, seeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{skipedgeexample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	seeuo.mutation.done = true
	return _node, nil
}
