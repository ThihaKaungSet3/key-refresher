// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/token-refresher/ent/apikey"
	"github.com/token-refresher/ent/predicate"
)

// ApiKeyUpdate is the builder for updating ApiKey entities.
type ApiKeyUpdate struct {
	config
	hooks    []Hook
	mutation *ApiKeyMutation
}

// Where appends a list predicates to the ApiKeyUpdate builder.
func (aku *ApiKeyUpdate) Where(ps ...predicate.ApiKey) *ApiKeyUpdate {
	aku.mutation.Where(ps...)
	return aku
}

// SetRemainingCredit sets the "remaining_credit" field.
func (aku *ApiKeyUpdate) SetRemainingCredit(i int) *ApiKeyUpdate {
	aku.mutation.ResetRemainingCredit()
	aku.mutation.SetRemainingCredit(i)
	return aku
}

// AddRemainingCredit adds i to the "remaining_credit" field.
func (aku *ApiKeyUpdate) AddRemainingCredit(i int) *ApiKeyUpdate {
	aku.mutation.AddRemainingCredit(i)
	return aku
}

// SetStatus sets the "status" field.
func (aku *ApiKeyUpdate) SetStatus(a apikey.Status) *ApiKeyUpdate {
	aku.mutation.SetStatus(a)
	return aku
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (aku *ApiKeyUpdate) SetNillableStatus(a *apikey.Status) *ApiKeyUpdate {
	if a != nil {
		aku.SetStatus(*a)
	}
	return aku
}

// SetKey sets the "key" field.
func (aku *ApiKeyUpdate) SetKey(s string) *ApiKeyUpdate {
	aku.mutation.SetKey(s)
	return aku
}

// Mutation returns the ApiKeyMutation object of the builder.
func (aku *ApiKeyUpdate) Mutation() *ApiKeyMutation {
	return aku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aku *ApiKeyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aku.hooks) == 0 {
		if err = aku.check(); err != nil {
			return 0, err
		}
		affected, err = aku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApiKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aku.check(); err != nil {
				return 0, err
			}
			aku.mutation = mutation
			affected, err = aku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aku.hooks) - 1; i >= 0; i-- {
			if aku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aku *ApiKeyUpdate) SaveX(ctx context.Context) int {
	affected, err := aku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aku *ApiKeyUpdate) Exec(ctx context.Context) error {
	_, err := aku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aku *ApiKeyUpdate) ExecX(ctx context.Context) {
	if err := aku.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aku *ApiKeyUpdate) check() error {
	if v, ok := aku.mutation.Status(); ok {
		if err := apikey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ApiKey.status": %w`, err)}
		}
	}
	return nil
}

func (aku *ApiKeyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apikey.Table,
			Columns: apikey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: apikey.FieldID,
			},
		},
	}
	if ps := aku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aku.mutation.RemainingCredit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apikey.FieldRemainingCredit,
		})
	}
	if value, ok := aku.mutation.AddedRemainingCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apikey.FieldRemainingCredit,
		})
	}
	if value, ok := aku.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: apikey.FieldStatus,
		})
	}
	if value, ok := aku.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apikey.FieldKey,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ApiKeyUpdateOne is the builder for updating a single ApiKey entity.
type ApiKeyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApiKeyMutation
}

// SetRemainingCredit sets the "remaining_credit" field.
func (akuo *ApiKeyUpdateOne) SetRemainingCredit(i int) *ApiKeyUpdateOne {
	akuo.mutation.ResetRemainingCredit()
	akuo.mutation.SetRemainingCredit(i)
	return akuo
}

// AddRemainingCredit adds i to the "remaining_credit" field.
func (akuo *ApiKeyUpdateOne) AddRemainingCredit(i int) *ApiKeyUpdateOne {
	akuo.mutation.AddRemainingCredit(i)
	return akuo
}

// SetStatus sets the "status" field.
func (akuo *ApiKeyUpdateOne) SetStatus(a apikey.Status) *ApiKeyUpdateOne {
	akuo.mutation.SetStatus(a)
	return akuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (akuo *ApiKeyUpdateOne) SetNillableStatus(a *apikey.Status) *ApiKeyUpdateOne {
	if a != nil {
		akuo.SetStatus(*a)
	}
	return akuo
}

// SetKey sets the "key" field.
func (akuo *ApiKeyUpdateOne) SetKey(s string) *ApiKeyUpdateOne {
	akuo.mutation.SetKey(s)
	return akuo
}

// Mutation returns the ApiKeyMutation object of the builder.
func (akuo *ApiKeyUpdateOne) Mutation() *ApiKeyMutation {
	return akuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (akuo *ApiKeyUpdateOne) Select(field string, fields ...string) *ApiKeyUpdateOne {
	akuo.fields = append([]string{field}, fields...)
	return akuo
}

// Save executes the query and returns the updated ApiKey entity.
func (akuo *ApiKeyUpdateOne) Save(ctx context.Context) (*ApiKey, error) {
	var (
		err  error
		node *ApiKey
	)
	if len(akuo.hooks) == 0 {
		if err = akuo.check(); err != nil {
			return nil, err
		}
		node, err = akuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ApiKeyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = akuo.check(); err != nil {
				return nil, err
			}
			akuo.mutation = mutation
			node, err = akuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(akuo.hooks) - 1; i >= 0; i-- {
			if akuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = akuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, akuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ApiKey)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ApiKeyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) SaveX(ctx context.Context) *ApiKey {
	node, err := akuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (akuo *ApiKeyUpdateOne) Exec(ctx context.Context) error {
	_, err := akuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akuo *ApiKeyUpdateOne) ExecX(ctx context.Context) {
	if err := akuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (akuo *ApiKeyUpdateOne) check() error {
	if v, ok := akuo.mutation.Status(); ok {
		if err := apikey.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "ApiKey.status": %w`, err)}
		}
	}
	return nil
}

func (akuo *ApiKeyUpdateOne) sqlSave(ctx context.Context) (_node *ApiKey, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   apikey.Table,
			Columns: apikey.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: apikey.FieldID,
			},
		},
	}
	id, ok := akuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ApiKey.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := akuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apikey.FieldID)
		for _, f := range fields {
			if !apikey.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apikey.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := akuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := akuo.mutation.RemainingCredit(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apikey.FieldRemainingCredit,
		})
	}
	if value, ok := akuo.mutation.AddedRemainingCredit(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: apikey.FieldRemainingCredit,
		})
	}
	if value, ok := akuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: apikey.FieldStatus,
		})
	}
	if value, ok := akuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: apikey.FieldKey,
		})
	}
	_node = &ApiKey{config: akuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, akuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apikey.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
