// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gas-detect/internal/data/model/ent/gasdetect"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GasDetectCreate is the builder for creating a GasDetect entity.
type GasDetectCreate struct {
	config
	mutation *GasDetectMutation
	hooks    []Hook
}

// SetCreatedTime sets the "created_time" field.
func (gdc *GasDetectCreate) SetCreatedTime(t time.Time) *GasDetectCreate {
	gdc.mutation.SetCreatedTime(t)
	return gdc
}

// SetNillableCreatedTime sets the "created_time" field if the given value is not nil.
func (gdc *GasDetectCreate) SetNillableCreatedTime(t *time.Time) *GasDetectCreate {
	if t != nil {
		gdc.SetCreatedTime(*t)
	}
	return gdc
}

// SetIsDeleted sets the "is_deleted" field.
func (gdc *GasDetectCreate) SetIsDeleted(b bool) *GasDetectCreate {
	gdc.mutation.SetIsDeleted(b)
	return gdc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (gdc *GasDetectCreate) SetNillableIsDeleted(b *bool) *GasDetectCreate {
	if b != nil {
		gdc.SetIsDeleted(*b)
	}
	return gdc
}

// SetMetrics sets the "metrics" field.
func (gdc *GasDetectCreate) SetMetrics(m []map[string]string) *GasDetectCreate {
	gdc.mutation.SetMetrics(m)
	return gdc
}

// SetID sets the "id" field.
func (gdc *GasDetectCreate) SetID(i int32) *GasDetectCreate {
	gdc.mutation.SetID(i)
	return gdc
}

// Mutation returns the GasDetectMutation object of the builder.
func (gdc *GasDetectCreate) Mutation() *GasDetectMutation {
	return gdc.mutation
}

// Save creates the GasDetect in the database.
func (gdc *GasDetectCreate) Save(ctx context.Context) (*GasDetect, error) {
	var (
		err  error
		node *GasDetect
	)
	gdc.defaults()
	if len(gdc.hooks) == 0 {
		if err = gdc.check(); err != nil {
			return nil, err
		}
		node, err = gdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GasDetectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gdc.check(); err != nil {
				return nil, err
			}
			gdc.mutation = mutation
			if node, err = gdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gdc.hooks) - 1; i >= 0; i-- {
			if gdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GasDetect)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GasDetectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gdc *GasDetectCreate) SaveX(ctx context.Context) *GasDetect {
	v, err := gdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gdc *GasDetectCreate) Exec(ctx context.Context) error {
	_, err := gdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gdc *GasDetectCreate) ExecX(ctx context.Context) {
	if err := gdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gdc *GasDetectCreate) defaults() {
	if _, ok := gdc.mutation.CreatedTime(); !ok {
		v := gasdetect.DefaultCreatedTime()
		gdc.mutation.SetCreatedTime(v)
	}
	if _, ok := gdc.mutation.IsDeleted(); !ok {
		v := gasdetect.DefaultIsDeleted
		gdc.mutation.SetIsDeleted(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gdc *GasDetectCreate) check() error {
	if _, ok := gdc.mutation.CreatedTime(); !ok {
		return &ValidationError{Name: "created_time", err: errors.New(`ent: missing required field "GasDetect.created_time"`)}
	}
	if _, ok := gdc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "GasDetect.is_deleted"`)}
	}
	return nil
}

func (gdc *GasDetectCreate) sqlSave(ctx context.Context) (*GasDetect, error) {
	_node, _spec := gdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (gdc *GasDetectCreate) createSpec() (*GasDetect, *sqlgraph.CreateSpec) {
	var (
		_node = &GasDetect{config: gdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: gasdetect.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: gasdetect.FieldID,
			},
		}
	)
	if id, ok := gdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gdc.mutation.CreatedTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: gasdetect.FieldCreatedTime,
		})
		_node.CreatedTime = value
	}
	if value, ok := gdc.mutation.IsDeleted(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: gasdetect.FieldIsDeleted,
		})
		_node.IsDeleted = value
	}
	if value, ok := gdc.mutation.Metrics(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: gasdetect.FieldMetrics,
		})
		_node.Metrics = value
	}
	return _node, _spec
}

// GasDetectCreateBulk is the builder for creating many GasDetect entities in bulk.
type GasDetectCreateBulk struct {
	config
	builders []*GasDetectCreate
}

// Save creates the GasDetect entities in the database.
func (gdcb *GasDetectCreateBulk) Save(ctx context.Context) ([]*GasDetect, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gdcb.builders))
	nodes := make([]*GasDetect, len(gdcb.builders))
	mutators := make([]Mutator, len(gdcb.builders))
	for i := range gdcb.builders {
		func(i int, root context.Context) {
			builder := gdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GasDetectMutation)
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
					_, err = mutators[i+1].Mutate(root, gdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
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
		if _, err := mutators[0].Mutate(ctx, gdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gdcb *GasDetectCreateBulk) SaveX(ctx context.Context) []*GasDetect {
	v, err := gdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gdcb *GasDetectCreateBulk) Exec(ctx context.Context) error {
	_, err := gdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gdcb *GasDetectCreateBulk) ExecX(ctx context.Context) {
	if err := gdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
