// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ent-atlas-test/ent/book"
	"ent-atlas-test/ent/predicate"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookUpdate is the builder for updating Book entities.
type BookUpdate struct {
	config
	hooks    []Hook
	mutation *BookMutation
}

// Where appends a list predicates to the BookUpdate builder.
func (bu *BookUpdate) Where(ps ...predicate.Book) *BookUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BookUpdate) SetTitle(s string) *BookUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BookUpdate) SetNillableTitle(s *string) *BookUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetBody sets the "body" field.
func (bu *BookUpdate) SetBody(s string) *BookUpdate {
	bu.mutation.SetBody(s)
	return bu
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (bu *BookUpdate) SetNillableBody(s *string) *BookUpdate {
	if s != nil {
		bu.SetBody(*s)
	}
	return bu
}

// SetPrice sets the "price" field.
func (bu *BookUpdate) SetPrice(i int) *BookUpdate {
	bu.mutation.ResetPrice()
	bu.mutation.SetPrice(i)
	return bu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (bu *BookUpdate) SetNillablePrice(i *int) *BookUpdate {
	if i != nil {
		bu.SetPrice(*i)
	}
	return bu
}

// AddPrice adds i to the "price" field.
func (bu *BookUpdate) AddPrice(i int) *BookUpdate {
	bu.mutation.AddPrice(i)
	return bu
}

// SetThoughts sets the "thoughts" field.
func (bu *BookUpdate) SetThoughts(s string) *BookUpdate {
	bu.mutation.SetThoughts(s)
	return bu
}

// SetNillableThoughts sets the "thoughts" field if the given value is not nil.
func (bu *BookUpdate) SetNillableThoughts(s *string) *BookUpdate {
	if s != nil {
		bu.SetThoughts(*s)
	}
	return bu
}

// Mutation returns the BookMutation object of the builder.
func (bu *BookUpdate) Mutation() *BookMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookUpdate) check() error {
	if v, ok := bu.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Body(); ok {
		if err := book.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Book.body": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Price(); ok {
		if err := book.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Book.price": %w`, err)}
		}
	}
	if v, ok := bu.mutation.Thoughts(); ok {
		if err := book.ThoughtsValidator(v); err != nil {
			return &ValidationError{Name: "thoughts", err: fmt.Errorf(`ent: validator failed for field "Book.thoughts": %w`, err)}
		}
	}
	return nil
}

func (bu *BookUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Body(); ok {
		_spec.SetField(book.FieldBody, field.TypeString, value)
	}
	if value, ok := bu.mutation.Price(); ok {
		_spec.SetField(book.FieldPrice, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedPrice(); ok {
		_spec.AddField(book.FieldPrice, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Thoughts(); ok {
		_spec.SetField(book.FieldThoughts, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookUpdateOne is the builder for updating a single Book entity.
type BookUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookMutation
}

// SetTitle sets the "title" field.
func (buo *BookUpdateOne) SetTitle(s string) *BookUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableTitle(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetBody sets the "body" field.
func (buo *BookUpdateOne) SetBody(s string) *BookUpdateOne {
	buo.mutation.SetBody(s)
	return buo
}

// SetNillableBody sets the "body" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableBody(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetBody(*s)
	}
	return buo
}

// SetPrice sets the "price" field.
func (buo *BookUpdateOne) SetPrice(i int) *BookUpdateOne {
	buo.mutation.ResetPrice()
	buo.mutation.SetPrice(i)
	return buo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillablePrice(i *int) *BookUpdateOne {
	if i != nil {
		buo.SetPrice(*i)
	}
	return buo
}

// AddPrice adds i to the "price" field.
func (buo *BookUpdateOne) AddPrice(i int) *BookUpdateOne {
	buo.mutation.AddPrice(i)
	return buo
}

// SetThoughts sets the "thoughts" field.
func (buo *BookUpdateOne) SetThoughts(s string) *BookUpdateOne {
	buo.mutation.SetThoughts(s)
	return buo
}

// SetNillableThoughts sets the "thoughts" field if the given value is not nil.
func (buo *BookUpdateOne) SetNillableThoughts(s *string) *BookUpdateOne {
	if s != nil {
		buo.SetThoughts(*s)
	}
	return buo
}

// Mutation returns the BookMutation object of the builder.
func (buo *BookUpdateOne) Mutation() *BookMutation {
	return buo.mutation
}

// Where appends a list predicates to the BookUpdate builder.
func (buo *BookUpdateOne) Where(ps ...predicate.Book) *BookUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookUpdateOne) Select(field string, fields ...string) *BookUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Book entity.
func (buo *BookUpdateOne) Save(ctx context.Context) (*Book, error) {
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookUpdateOne) SaveX(ctx context.Context) *Book {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookUpdateOne) check() error {
	if v, ok := buo.mutation.Title(); ok {
		if err := book.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Book.title": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Body(); ok {
		if err := book.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "Book.body": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Price(); ok {
		if err := book.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Book.price": %w`, err)}
		}
	}
	if v, ok := buo.mutation.Thoughts(); ok {
		if err := book.ThoughtsValidator(v); err != nil {
			return &ValidationError{Name: "thoughts", err: fmt.Errorf(`ent: validator failed for field "Book.thoughts": %w`, err)}
		}
	}
	return nil
}

func (buo *BookUpdateOne) sqlSave(ctx context.Context) (_node *Book, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(book.Table, book.Columns, sqlgraph.NewFieldSpec(book.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Book.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, book.FieldID)
		for _, f := range fields {
			if !book.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != book.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Body(); ok {
		_spec.SetField(book.FieldBody, field.TypeString, value)
	}
	if value, ok := buo.mutation.Price(); ok {
		_spec.SetField(book.FieldPrice, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedPrice(); ok {
		_spec.AddField(book.FieldPrice, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Thoughts(); ok {
		_spec.SetField(book.FieldThoughts, field.TypeString, value)
	}
	_node = &Book{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{book.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
