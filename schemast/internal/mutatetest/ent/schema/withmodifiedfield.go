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

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type WithModifiedField struct {
	ent.Schema
}

func (WithModifiedField) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Immutable().MaxLen(10),
	}
}

func (WithModifiedField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).Unique(),
	}
}
