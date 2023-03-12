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

package billproduct

import (
	"entgo.io/contrib/entgql/internal/todogotype/ent/predicate"
	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldName, v))
}

// Sku applies equality check predicate on the "sku" field. It's identical to SkuEQ.
func Sku(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldSku, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldQuantity, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldContainsFold(FieldName, v))
}

// SkuEQ applies the EQ predicate on the "sku" field.
func SkuEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldSku, v))
}

// SkuNEQ applies the NEQ predicate on the "sku" field.
func SkuNEQ(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNEQ(FieldSku, v))
}

// SkuIn applies the In predicate on the "sku" field.
func SkuIn(vs ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldIn(FieldSku, vs...))
}

// SkuNotIn applies the NotIn predicate on the "sku" field.
func SkuNotIn(vs ...string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNotIn(FieldSku, vs...))
}

// SkuGT applies the GT predicate on the "sku" field.
func SkuGT(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGT(FieldSku, v))
}

// SkuGTE applies the GTE predicate on the "sku" field.
func SkuGTE(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGTE(FieldSku, v))
}

// SkuLT applies the LT predicate on the "sku" field.
func SkuLT(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLT(FieldSku, v))
}

// SkuLTE applies the LTE predicate on the "sku" field.
func SkuLTE(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLTE(FieldSku, v))
}

// SkuContains applies the Contains predicate on the "sku" field.
func SkuContains(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldContains(FieldSku, v))
}

// SkuHasPrefix applies the HasPrefix predicate on the "sku" field.
func SkuHasPrefix(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldHasPrefix(FieldSku, v))
}

// SkuHasSuffix applies the HasSuffix predicate on the "sku" field.
func SkuHasSuffix(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldHasSuffix(FieldSku, v))
}

// SkuEqualFold applies the EqualFold predicate on the "sku" field.
func SkuEqualFold(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEqualFold(FieldSku, v))
}

// SkuContainsFold applies the ContainsFold predicate on the "sku" field.
func SkuContainsFold(v string) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldContainsFold(FieldSku, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v uint64) predicate.BillProduct {
	return predicate.BillProduct(sql.FieldLTE(FieldQuantity, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BillProduct) predicate.BillProduct {
	return predicate.BillProduct(func(s *sql.Selector) {
		p(s.Not())
	})
}
