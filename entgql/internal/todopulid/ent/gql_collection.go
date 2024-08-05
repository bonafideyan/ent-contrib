// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entgql/internal/todopulid/ent/billproduct"
	"entgo.io/contrib/entgql/internal/todopulid/ent/category"
	"entgo.io/contrib/entgql/internal/todopulid/ent/friendship"
	"entgo.io/contrib/entgql/internal/todopulid/ent/group"
	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/contrib/entgql/internal/todopulid/ent/todo"
	"entgo.io/contrib/entgql/internal/todopulid/ent/user"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (bp *BillProductQuery) CollectFields(ctx context.Context, satisfies ...string) (*BillProductQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return bp, nil
	}
	if err := bp.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return bp, nil
}

func (bp *BillProductQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(billproduct.Columns))
		selectedFields = []string{billproduct.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[billproduct.FieldName]; !ok {
				selectedFields = append(selectedFields, billproduct.FieldName)
				fieldSeen[billproduct.FieldName] = struct{}{}
			}
		case "sku":
			if _, ok := fieldSeen[billproduct.FieldSku]; !ok {
				selectedFields = append(selectedFields, billproduct.FieldSku)
				fieldSeen[billproduct.FieldSku] = struct{}{}
			}
		case "quantity":
			if _, ok := fieldSeen[billproduct.FieldQuantity]; !ok {
				selectedFields = append(selectedFields, billproduct.FieldQuantity)
				fieldSeen[billproduct.FieldQuantity] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		bp.Select(selectedFields...)
	}
	return nil
}

type billproductPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BillProductPaginateOption
}

func newBillProductPaginateArgs(rv map[string]any) *billproductPaginateArgs {
	args := &billproductPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*BillProductWhereInput); ok {
		args.opts = append(args.opts, WithBillProductFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CategoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*CategoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CategoryQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(category.Columns))
		selectedFields = []string{category.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "todos":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: c.config}).Query()
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID pulid.ID `sql:"category_id"`
							Count  int      `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(category.TodosColumn), ids...))
						})
						if err := query.GroupBy(category.TodosColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[pulid.ID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Todos)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, todoImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(category.TodosColumn, limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			c.WithNamedTodos(alias, func(wq *TodoQuery) {
				*wq = *query
			})

		case "subCategories":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: c.config}).Query()
			)
			args := newCategoryPaginateArgs(fieldArgs(ctx, new(CategoryWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newCategoryPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					c.loadTotal = append(c.loadTotal, func(ctx context.Context, nodes []*Category) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID pulid.ID `sql:"category_id"`
							Count  int      `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(category.SubCategoriesTable)
							s.Join(joinT).On(s.C(category.FieldID), joinT.C(category.SubCategoriesPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(category.SubCategoriesPrimaryKey[0]), ids...))
							s.Select(joinT.C(category.SubCategoriesPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(category.SubCategoriesPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[pulid.ID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					c.loadTotal = append(c.loadTotal, func(_ context.Context, nodes []*Category) error {
						for i := range nodes {
							n := len(nodes[i].Edges.SubCategories)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, categoryImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(category.SubCategoriesPrimaryKey[0], limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			c.WithNamedSubCategories(alias, func(wq *CategoryQuery) {
				*wq = *query
			})
		case "text":
			if _, ok := fieldSeen[category.FieldText]; !ok {
				selectedFields = append(selectedFields, category.FieldText)
				fieldSeen[category.FieldText] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[category.FieldStatus]; !ok {
				selectedFields = append(selectedFields, category.FieldStatus)
				fieldSeen[category.FieldStatus] = struct{}{}
			}
		case "config":
			if _, ok := fieldSeen[category.FieldConfig]; !ok {
				selectedFields = append(selectedFields, category.FieldConfig)
				fieldSeen[category.FieldConfig] = struct{}{}
			}
		case "types":
			if _, ok := fieldSeen[category.FieldTypes]; !ok {
				selectedFields = append(selectedFields, category.FieldTypes)
				fieldSeen[category.FieldTypes] = struct{}{}
			}
		case "duration":
			if _, ok := fieldSeen[category.FieldDuration]; !ok {
				selectedFields = append(selectedFields, category.FieldDuration)
				fieldSeen[category.FieldDuration] = struct{}{}
			}
		case "count":
			if _, ok := fieldSeen[category.FieldCount]; !ok {
				selectedFields = append(selectedFields, category.FieldCount)
				fieldSeen[category.FieldCount] = struct{}{}
			}
		case "strings":
			if _, ok := fieldSeen[category.FieldStrings]; !ok {
				selectedFields = append(selectedFields, category.FieldStrings)
				fieldSeen[category.FieldStrings] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type categoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CategoryPaginateOption
}

func newCategoryPaginateArgs(rv map[string]any) *categoryPaginateArgs {
	args := &categoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*CategoryOrder:
			args.opts = append(args.opts, WithCategoryOrder(v))
		case []any:
			var orders []*CategoryOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &CategoryOrder{Field: &CategoryOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithCategoryOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*CategoryWhereInput); ok {
		args.opts = append(args.opts, WithCategoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FriendshipQuery) CollectFields(ctx context.Context, satisfies ...string) (*FriendshipQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FriendshipQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(friendship.Columns))
		selectedFields = []string{friendship.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			f.withUser = query
			if _, ok := fieldSeen[friendship.FieldUserID]; !ok {
				selectedFields = append(selectedFields, friendship.FieldUserID)
				fieldSeen[friendship.FieldUserID] = struct{}{}
			}

		case "friend":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			f.withFriend = query
			if _, ok := fieldSeen[friendship.FieldFriendID]; !ok {
				selectedFields = append(selectedFields, friendship.FieldFriendID)
				fieldSeen[friendship.FieldFriendID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[friendship.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, friendship.FieldCreatedAt)
				fieldSeen[friendship.FieldCreatedAt] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[friendship.FieldUserID]; !ok {
				selectedFields = append(selectedFields, friendship.FieldUserID)
				fieldSeen[friendship.FieldUserID] = struct{}{}
			}
		case "friendID":
			if _, ok := fieldSeen[friendship.FieldFriendID]; !ok {
				selectedFields = append(selectedFields, friendship.FieldFriendID)
				fieldSeen[friendship.FieldFriendID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		f.Select(selectedFields...)
	}
	return nil
}

type friendshipPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FriendshipPaginateOption
}

func newFriendshipPaginateArgs(rv map[string]any) *friendshipPaginateArgs {
	args := &friendshipPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*FriendshipWhereInput); ok {
		args.opts = append(args.opts, WithFriendshipFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GroupQuery) CollectFields(ctx context.Context, satisfies ...string) (*GroupQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GroupQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(group.Columns))
		selectedFields = []string{group.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "users":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: gr.config}).Query()
			)
			args := newUserPaginateArgs(fieldArgs(ctx, new(UserWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newUserPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					gr.loadTotal = append(gr.loadTotal, func(ctx context.Context, nodes []*Group) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID pulid.ID `sql:"group_id"`
							Count  int      `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(group.UsersTable)
							s.Join(joinT).On(s.C(user.FieldID), joinT.C(group.UsersPrimaryKey[0]))
							s.Where(sql.InValues(joinT.C(group.UsersPrimaryKey[1]), ids...))
							s.Select(joinT.C(group.UsersPrimaryKey[1]), sql.Count("*"))
							s.GroupBy(joinT.C(group.UsersPrimaryKey[1]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[pulid.ID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					gr.loadTotal = append(gr.loadTotal, func(_ context.Context, nodes []*Group) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Users)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(group.UsersPrimaryKey[1], limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			gr.WithNamedUsers(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[group.FieldName]; !ok {
				selectedFields = append(selectedFields, group.FieldName)
				fieldSeen[group.FieldName] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		gr.Select(selectedFields...)
	}
	return nil
}

type groupPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GroupPaginateOption
}

func newGroupPaginateArgs(rv map[string]any) *groupPaginateArgs {
	args := &groupPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*GroupWhereInput); ok {
		args.opts = append(args.opts, WithGroupFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TodoQuery) CollectFields(ctx context.Context, satisfies ...string) (*TodoQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TodoQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(todo.Columns))
		selectedFields = []string{todo.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, todoImplementors)...); err != nil {
				return err
			}
			t.withParent = query

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TodoClient{config: t.config}).Query()
			)
			args := newTodoPaginateArgs(fieldArgs(ctx, new(TodoWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newTodoPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					t.loadTotal = append(t.loadTotal, func(ctx context.Context, nodes []*Todo) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID pulid.ID `sql:"todo_children"`
							Count  int      `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(todo.ChildrenColumn), ids...))
						})
						if err := query.GroupBy(todo.ChildrenColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[pulid.ID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				} else {
					t.loadTotal = append(t.loadTotal, func(_ context.Context, nodes []*Todo) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Children)
							if nodes[i].Edges.totalCount[1] == nil {
								nodes[i].Edges.totalCount[1] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[1][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, todoImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(todo.ChildrenColumn, limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			t.WithNamedChildren(alias, func(wq *TodoQuery) {
				*wq = *query
			})

		case "category":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CategoryClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, categoryImplementors)...); err != nil {
				return err
			}
			t.withCategory = query
			if _, ok := fieldSeen[todo.FieldCategoryID]; !ok {
				selectedFields = append(selectedFields, todo.FieldCategoryID)
				fieldSeen[todo.FieldCategoryID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[todo.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, todo.FieldCreatedAt)
				fieldSeen[todo.FieldCreatedAt] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[todo.FieldStatus]; !ok {
				selectedFields = append(selectedFields, todo.FieldStatus)
				fieldSeen[todo.FieldStatus] = struct{}{}
			}
		case "priorityOrder":
			if _, ok := fieldSeen[todo.FieldPriority]; !ok {
				selectedFields = append(selectedFields, todo.FieldPriority)
				fieldSeen[todo.FieldPriority] = struct{}{}
			}
		case "text":
			if _, ok := fieldSeen[todo.FieldText]; !ok {
				selectedFields = append(selectedFields, todo.FieldText)
				fieldSeen[todo.FieldText] = struct{}{}
			}
		case "init":
			if _, ok := fieldSeen[todo.FieldInit]; !ok {
				selectedFields = append(selectedFields, todo.FieldInit)
				fieldSeen[todo.FieldInit] = struct{}{}
			}
		case "custom":
			if _, ok := fieldSeen[todo.FieldCustom]; !ok {
				selectedFields = append(selectedFields, todo.FieldCustom)
				fieldSeen[todo.FieldCustom] = struct{}{}
			}
		case "customp":
			if _, ok := fieldSeen[todo.FieldCustomp]; !ok {
				selectedFields = append(selectedFields, todo.FieldCustomp)
				fieldSeen[todo.FieldCustomp] = struct{}{}
			}
		case "categoryID":
			if _, ok := fieldSeen[todo.FieldCategoryID]; !ok {
				selectedFields = append(selectedFields, todo.FieldCategoryID)
				fieldSeen[todo.FieldCategoryID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type todoPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TodoPaginateOption
}

func newTodoPaginateArgs(rv map[string]any) *todoPaginateArgs {
	args := &todoPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*TodoOrder:
			args.opts = append(args.opts, WithTodoOrder(v))
		case []any:
			var orders []*TodoOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &TodoOrder{Field: &TodoOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithTodoOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*TodoWhereInput); ok {
		args.opts = append(args.opts, WithTodoFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "groups":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&GroupClient{config: u.config}).Query()
			)
			args := newGroupPaginateArgs(fieldArgs(ctx, new(GroupWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newGroupPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					u.loadTotal = append(u.loadTotal, func(ctx context.Context, nodes []*User) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID pulid.ID `sql:"user_id"`
							Count  int      `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							joinT := sql.Table(user.GroupsTable)
							s.Join(joinT).On(s.C(group.FieldID), joinT.C(user.GroupsPrimaryKey[1]))
							s.Where(sql.InValues(joinT.C(user.GroupsPrimaryKey[0]), ids...))
							s.Select(joinT.C(user.GroupsPrimaryKey[0]), sql.Count("*"))
							s.GroupBy(joinT.C(user.GroupsPrimaryKey[0]))
						})
						if err := query.Select().Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[pulid.ID]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					u.loadTotal = append(u.loadTotal, func(_ context.Context, nodes []*User) error {
						for i := range nodes {
							n := len(nodes[i].Edges.Groups)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, groupImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(user.GroupsPrimaryKey[0], limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			u.WithNamedGroups(alias, func(wq *GroupQuery) {
				*wq = *query
			})

		case "friends":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			u.WithNamedFriends(alias, func(wq *UserQuery) {
				*wq = *query
			})

		case "friendships":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FriendshipClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, friendshipImplementors)...); err != nil {
				return err
			}
			u.WithNamedFriendships(alias, func(wq *FriendshipQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "username":
			if _, ok := fieldSeen[user.FieldUsername]; !ok {
				selectedFields = append(selectedFields, user.FieldUsername)
				fieldSeen[user.FieldUsername] = struct{}{}
			}
		case "requiredMetadata":
			if _, ok := fieldSeen[user.FieldRequiredMetadata]; !ok {
				selectedFields = append(selectedFields, user.FieldRequiredMetadata)
				fieldSeen[user.FieldRequiredMetadata] = struct{}{}
			}
		case "metadata":
			if _, ok := fieldSeen[user.FieldMetadata]; !ok {
				selectedFields = append(selectedFields, user.FieldMetadata)
				fieldSeen[user.FieldMetadata] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
