// Code generated by ent, DO NOT EDIT.

package apikey

import (
	"entgo.io/ent/dialect/sql"
	"github.com/token-refresher/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// RemainingCredit applies equality check predicate on the "remaining_credit" field. It's identical to RemainingCreditEQ.
func RemainingCredit(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemainingCredit), v))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// RemainingCreditEQ applies the EQ predicate on the "remaining_credit" field.
func RemainingCreditEQ(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemainingCredit), v))
	})
}

// RemainingCreditNEQ applies the NEQ predicate on the "remaining_credit" field.
func RemainingCreditNEQ(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemainingCredit), v))
	})
}

// RemainingCreditIn applies the In predicate on the "remaining_credit" field.
func RemainingCreditIn(vs ...int) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRemainingCredit), v...))
	})
}

// RemainingCreditNotIn applies the NotIn predicate on the "remaining_credit" field.
func RemainingCreditNotIn(vs ...int) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRemainingCredit), v...))
	})
}

// RemainingCreditGT applies the GT predicate on the "remaining_credit" field.
func RemainingCreditGT(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemainingCredit), v))
	})
}

// RemainingCreditGTE applies the GTE predicate on the "remaining_credit" field.
func RemainingCreditGTE(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemainingCredit), v))
	})
}

// RemainingCreditLT applies the LT predicate on the "remaining_credit" field.
func RemainingCreditLT(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemainingCredit), v))
	})
}

// RemainingCreditLTE applies the LTE predicate on the "remaining_credit" field.
func RemainingCreditLTE(v int) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemainingCredit), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.ApiKey {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
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
func Not(p predicate.ApiKey) predicate.ApiKey {
	return predicate.ApiKey(func(s *sql.Selector) {
		p(s.Not())
	})
}