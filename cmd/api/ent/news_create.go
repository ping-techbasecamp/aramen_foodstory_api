// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aramen-api/cmd/api/ent/news"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NewsCreate is the builder for creating a News entity.
type NewsCreate struct {
	config
	mutation *NewsMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (nc *NewsCreate) SetName(s string) *NewsCreate {
	nc.mutation.SetName(s)
	return nc
}

// SetDescription sets the "description" field.
func (nc *NewsCreate) SetDescription(s string) *NewsCreate {
	nc.mutation.SetDescription(s)
	return nc
}

// SetTag sets the "tag" field.
func (nc *NewsCreate) SetTag(s string) *NewsCreate {
	nc.mutation.SetTag(s)
	return nc
}

// SetPlace sets the "place" field.
func (nc *NewsCreate) SetPlace(s string) *NewsCreate {
	nc.mutation.SetPlace(s)
	return nc
}

// SetDate sets the "date" field.
func (nc *NewsCreate) SetDate(s string) *NewsCreate {
	nc.mutation.SetDate(s)
	return nc
}

// Mutation returns the NewsMutation object of the builder.
func (nc *NewsCreate) Mutation() *NewsMutation {
	return nc.mutation
}

// Save creates the News in the database.
func (nc *NewsCreate) Save(ctx context.Context) (*News, error) {
	var (
		err  error
		node *News
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NewsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			if node, err = nc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			if nc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, nc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (nc *NewsCreate) SaveX(ctx context.Context) *News {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NewsCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NewsCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NewsCreate) check() error {
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "News.name"`)}
	}
	if _, ok := nc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "News.description"`)}
	}
	if _, ok := nc.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required field "News.tag"`)}
	}
	if _, ok := nc.mutation.Place(); !ok {
		return &ValidationError{Name: "place", err: errors.New(`ent: missing required field "News.place"`)}
	}
	if _, ok := nc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "News.date"`)}
	}
	return nil
}

func (nc *NewsCreate) sqlSave(ctx context.Context) (*News, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NewsCreate) createSpec() (*News, *sqlgraph.CreateSpec) {
	var (
		_node = &News{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: news.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: news.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := nc.mutation.Tag(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldTag,
		})
		_node.Tag = value
	}
	if value, ok := nc.mutation.Place(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldPlace,
		})
		_node.Place = value
	}
	if value, ok := nc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: news.FieldDate,
		})
		_node.Date = value
	}
	return _node, _spec
}

// NewsCreateBulk is the builder for creating many News entities in bulk.
type NewsCreateBulk struct {
	config
	builders []*NewsCreate
}

// Save creates the News entities in the database.
func (ncb *NewsCreateBulk) Save(ctx context.Context) ([]*News, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*News, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NewsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NewsCreateBulk) SaveX(ctx context.Context) []*News {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NewsCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NewsCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}
