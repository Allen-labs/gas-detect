// Code generated by ent, DO NOT EDIT.

package gasdetect

import (
	"gas-detect/internal/data/model/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int32) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedTime applies equality check predicate on the "created_time" field. It's identical to CreatedTimeEQ.
func CreatedTime(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedTime), v))
	})
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v bool) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// CreatedTimeEQ applies the EQ predicate on the "created_time" field.
func CreatedTimeEQ(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedTime), v))
	})
}

// CreatedTimeNEQ applies the NEQ predicate on the "created_time" field.
func CreatedTimeNEQ(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedTime), v))
	})
}

// CreatedTimeIn applies the In predicate on the "created_time" field.
func CreatedTimeIn(vs ...time.Time) predicate.GasDetect {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GasDetect(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedTime), v...))
	})
}

// CreatedTimeNotIn applies the NotIn predicate on the "created_time" field.
func CreatedTimeNotIn(vs ...time.Time) predicate.GasDetect {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GasDetect(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedTime), v...))
	})
}

// CreatedTimeGT applies the GT predicate on the "created_time" field.
func CreatedTimeGT(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedTime), v))
	})
}

// CreatedTimeGTE applies the GTE predicate on the "created_time" field.
func CreatedTimeGTE(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedTime), v))
	})
}

// CreatedTimeLT applies the LT predicate on the "created_time" field.
func CreatedTimeLT(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedTime), v))
	})
}

// CreatedTimeLTE applies the LTE predicate on the "created_time" field.
func CreatedTimeLTE(v time.Time) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedTime), v))
	})
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v bool) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsDeleted), v))
	})
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v bool) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsDeleted), v))
	})
}

// MetricsIsNil applies the IsNil predicate on the "metrics" field.
func MetricsIsNil() predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMetrics)))
	})
}

// MetricsNotNil applies the NotNil predicate on the "metrics" field.
func MetricsNotNil() predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMetrics)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GasDetect) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GasDetect) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
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
func Not(p predicate.GasDetect) predicate.GasDetect {
	return predicate.GasDetect(func(s *sql.Selector) {
		p(s.Not())
	})
}
