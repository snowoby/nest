// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"nest/ent/fileticket"
	"nest/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// FileTicketDelete is the builder for deleting a FileTicket entity.
type FileTicketDelete struct {
	config
	hooks    []Hook
	mutation *FileTicketMutation
}

// Where adds a new predicate to the FileTicketDelete builder.
func (ftd *FileTicketDelete) Where(ps ...predicate.FileTicket) *FileTicketDelete {
	ftd.mutation.predicates = append(ftd.mutation.predicates, ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FileTicketDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftd.hooks) == 0 {
		affected, err = ftd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTicketMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftd.mutation = mutation
			affected, err = ftd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftd.hooks) - 1; i >= 0; i-- {
			mut = ftd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTicketDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTicketDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fileticket.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileticket.FieldID,
			},
		},
	}
	if ps := ftd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ftd.driver, _spec)
}

// FileTicketDeleteOne is the builder for deleting a single FileTicket entity.
type FileTicketDeleteOne struct {
	ftd *FileTicketDelete
}

// Exec executes the deletion query.
func (ftdo *FileTicketDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fileticket.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTicketDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
