// Code generated by entc, DO NOT EDIT.

package fileticket

import (
	"nest/ent/predicate"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// OriginalFilename applies equality check predicate on the "original_filename" field. It's identical to OriginalFilenameEQ.
func OriginalFilename(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalFilename), v))
	})
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUUID), v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUUID), v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUUID), v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUUID), v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUUID), v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUUID), v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUUID), v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUUID), v))
	})
}

// OriginalFilenameEQ applies the EQ predicate on the "original_filename" field.
func OriginalFilenameEQ(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameNEQ applies the NEQ predicate on the "original_filename" field.
func OriginalFilenameNEQ(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameIn applies the In predicate on the "original_filename" field.
func OriginalFilenameIn(vs ...string) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOriginalFilename), v...))
	})
}

// OriginalFilenameNotIn applies the NotIn predicate on the "original_filename" field.
func OriginalFilenameNotIn(vs ...string) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOriginalFilename), v...))
	})
}

// OriginalFilenameGT applies the GT predicate on the "original_filename" field.
func OriginalFilenameGT(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameGTE applies the GTE predicate on the "original_filename" field.
func OriginalFilenameGTE(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameLT applies the LT predicate on the "original_filename" field.
func OriginalFilenameLT(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameLTE applies the LTE predicate on the "original_filename" field.
func OriginalFilenameLTE(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameContains applies the Contains predicate on the "original_filename" field.
func OriginalFilenameContains(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameHasPrefix applies the HasPrefix predicate on the "original_filename" field.
func OriginalFilenameHasPrefix(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameHasSuffix applies the HasSuffix predicate on the "original_filename" field.
func OriginalFilenameHasSuffix(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameEqualFold applies the EqualFold predicate on the "original_filename" field.
func OriginalFilenameEqualFold(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOriginalFilename), v))
	})
}

// OriginalFilenameContainsFold applies the ContainsFold predicate on the "original_filename" field.
func OriginalFilenameContainsFold(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOriginalFilename), v))
	})
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIdentifier), v))
	})
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIdentifier), v...))
	})
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIdentifier), v...))
	})
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIdentifier), v))
	})
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIdentifier), v))
	})
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIdentifier), v))
	})
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIdentifier), v))
	})
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIdentifier), v))
	})
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIdentifier), v))
	})
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIdentifier), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FileTicket {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FileTicket(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FilesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FilesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHosts applies the HasEdge predicate on the "hosts" edge.
func HasHosts() predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostsTable, HostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHostsWith applies the HasEdge predicate on the "hosts" edge with a given conditions (other predicates).
func HasHostsWith(preds ...predicate.Host) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HostsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, HostsTable, HostsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FileTicket) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FileTicket) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
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
func Not(p predicate.FileTicket) predicate.FileTicket {
	return predicate.FileTicket(func(s *sql.Selector) {
		p(s.Not())
	})
}
