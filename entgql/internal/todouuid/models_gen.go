// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package todo

import (
	"github.com/facebookincubator/ent-contrib/entgql/internal/todouuid/ent/schema/uuidgql"
	"github.com/facebookincubator/ent-contrib/entgql/internal/todouuid/ent/todo"
)

type TodoInput struct {
	Status   todo.Status   `json:"status"`
	Priority *int          `json:"priority"`
	Text     string        `json:"text"`
	Parent   *uuidgql.UUID `json:"parent"`
}
