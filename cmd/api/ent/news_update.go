// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aramen-api/cmd/api/ent/news"
	"aramen-api/cmd/api/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsUpdate is the builder for updating News entities.
type NewsUpdate struct {
	config
	hooks    []Hook
	mutation *NewsMutation
}

// Where appends a list predicates to the NewsUpdate builder.
func (nu *NewsUpdate) Where(ps ...predicate.News) *NewsUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetName sets the "name" field.
func (nu *NewsUpdate) SetName(s string) *NewsUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetDescription sets the "description" field.
func (nu *NewsUpdate) SetDescription(s string) *NewsUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetTag sets the "tag" field.
func (nu *NewsUpdate) SetTag(s string) *NewsUpdate {
	nu.mutation.SetTag(s)
	return nu
}

// SetPlace sets the "place" field.
func (nu *NewsUpdate) SetPlace(s string) *NewsUpdate {
	nu.mutation.SetPlace(s)
	return nu
}

// SetDate sets the "date" field.
func (nu *NewsUpdate) SetDate(s string) *NewsUpdate {
	nu.mutation.SetDate(s)
	return nu
}

// Mutation returns the NewsMutation object of the builder.
func (nu *NewsUpdate) Mutation() *NewsMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NewsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nu.hooks) == 0 {
		affected, err = nu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nu.mutation = mutation
			affected, err = nu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nu.hooks) - 1; i >= 0; i-- {
			if nu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NewsUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NewsUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NewsUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NewsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   news.Table,
			Columns: news.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: news.FieldID,
			},
		},
	}
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldName,
		})
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDescription,
		})
	}
	if value, ok := nu.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldTag,
		})
	}
	if value, ok := nu.mutation.Place(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldPlace,
		})
	}
	if value, ok := nu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDate,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// NewsUpdateOne is the builder for updating a single News entity.
type NewsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NewsMutation
}

// SetName sets the "name" field.
func (nuo *NewsUpdateOne) SetName(s string) *NewsUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NewsUpdateOne) SetDescription(s string) *NewsUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetTag sets the "tag" field.
func (nuo *NewsUpdateOne) SetTag(s string) *NewsUpdateOne {
	nuo.mutation.SetTag(s)
	return nuo
}

// SetPlace sets the "place" field.
func (nuo *NewsUpdateOne) SetPlace(s string) *NewsUpdateOne {
	nuo.mutation.SetPlace(s)
	return nuo
}

// SetDate sets the "date" field.
func (nuo *NewsUpdateOne) SetDate(s string) *NewsUpdateOne {
	nuo.mutation.SetDate(s)
	return nuo
}

// Mutation returns the NewsMutation object of the builder.
func (nuo *NewsUpdateOne) Mutation() *NewsMutation {
	return nuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NewsUpdateOne) Select(field string, fields ...string) *NewsUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated News entity.
func (nuo *NewsUpdateOne) Save(ctx context.Context) (*News, error) {
	var (
		err  error
		node *News
	)
	if len(nuo.hooks) == 0 {
		node, err = nuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nuo.mutation = mutation
			node, err = nuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nuo.hooks) - 1; i >= 0; i-- {
			if nuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*News)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from NewsMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NewsUpdateOne) SaveX(ctx context.Context) *News {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NewsUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NewsUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NewsUpdateOne) sqlSave(ctx context.Context) (_node *News, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   news.Table,
			Columns: news.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: news.FieldID,
			},
		},
	}
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "News.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, news.FieldID)
		for _, f := range fields {
			if !news.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != news.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldName,
		})
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDescription,
		})
	}
	if value, ok := nuo.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldTag,
		})
	}
	if value, ok := nuo.mutation.Place(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldPlace,
		})
	}
	if value, ok := nuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDate,
		})
	}
	_node = &News{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{news.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}