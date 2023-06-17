// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aramen-api/cmd/api/ent/user"
	"aramen-api/cmd/api/ent/userotp"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserOtpCreate is the builder for creating a UserOtp entity.
type UserOtpCreate struct {
	config
	mutation *UserOtpMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uoc *UserOtpCreate) SetUserID(i int) *UserOtpCreate {
	uoc.mutation.SetUserID(i)
	return uoc
}

// SetReference sets the "reference" field.
func (uoc *UserOtpCreate) SetReference(s string) *UserOtpCreate {
	uoc.mutation.SetReference(s)
	return uoc
}

// SetHashedOtp sets the "hashed_otp" field.
func (uoc *UserOtpCreate) SetHashedOtp(s string) *UserOtpCreate {
	uoc.mutation.SetHashedOtp(s)
	return uoc
}

// SetUser sets the "user" edge to the User entity.
func (uoc *UserOtpCreate) SetUser(u *User) *UserOtpCreate {
	return uoc.SetUserID(u.ID)
}

// Mutation returns the UserOtpMutation object of the builder.
func (uoc *UserOtpCreate) Mutation() *UserOtpMutation {
	return uoc.mutation
}

// Save creates the UserOtp in the database.
func (uoc *UserOtpCreate) Save(ctx context.Context) (*UserOtp, error) {
	var (
		err  error
		node *UserOtp
	)
	if len(uoc.hooks) == 0 {
		if err = uoc.check(); err != nil {
			return nil, err
		}
		node, err = uoc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserOtpMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uoc.check(); err != nil {
				return nil, err
			}
			uoc.mutation = mutation
			if node, err = uoc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uoc.hooks) - 1; i >= 0; i-- {
			if uoc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uoc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uoc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserOtp)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserOtpMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uoc *UserOtpCreate) SaveX(ctx context.Context) *UserOtp {
	v, err := uoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uoc *UserOtpCreate) Exec(ctx context.Context) error {
	_, err := uoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uoc *UserOtpCreate) ExecX(ctx context.Context) {
	if err := uoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uoc *UserOtpCreate) check() error {
	if _, ok := uoc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserOtp.user_id"`)}
	}
	if _, ok := uoc.mutation.Reference(); !ok {
		return &ValidationError{Name: "reference", err: errors.New(`ent: missing required field "UserOtp.reference"`)}
	}
	if _, ok := uoc.mutation.HashedOtp(); !ok {
		return &ValidationError{Name: "hashed_otp", err: errors.New(`ent: missing required field "UserOtp.hashed_otp"`)}
	}
	if _, ok := uoc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserOtp.user"`)}
	}
	return nil
}

func (uoc *UserOtpCreate) sqlSave(ctx context.Context) (*UserOtp, error) {
	_node, _spec := uoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (uoc *UserOtpCreate) createSpec() (*UserOtp, *sqlgraph.CreateSpec) {
	var (
		_node = &UserOtp{config: uoc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userotp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userotp.FieldID,
			},
		}
	)
	if value, ok := uoc.mutation.Reference(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userotp.FieldReference,
		})
		_node.Reference = value
	}
	if value, ok := uoc.mutation.HashedOtp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userotp.FieldHashedOtp,
		})
		_node.HashedOtp = value
	}
	if nodes := uoc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userotp.UserTable,
			Columns: []string{userotp.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserOtpCreateBulk is the builder for creating many UserOtp entities in bulk.
type UserOtpCreateBulk struct {
	config
	builders []*UserOtpCreate
}

// Save creates the UserOtp entities in the database.
func (uocb *UserOtpCreateBulk) Save(ctx context.Context) ([]*UserOtp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uocb.builders))
	nodes := make([]*UserOtp, len(uocb.builders))
	mutators := make([]Mutator, len(uocb.builders))
	for i := range uocb.builders {
		func(i int, root context.Context) {
			builder := uocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserOtpMutation)
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
					_, err = mutators[i+1].Mutate(root, uocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uocb *UserOtpCreateBulk) SaveX(ctx context.Context) []*UserOtp {
	v, err := uocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uocb *UserOtpCreateBulk) Exec(ctx context.Context) error {
	_, err := uocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uocb *UserOtpCreateBulk) ExecX(ctx context.Context) {
	if err := uocb.Exec(ctx); err != nil {
		panic(err)
	}
}
