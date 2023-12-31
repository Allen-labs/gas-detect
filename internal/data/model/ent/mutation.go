// Code generated by ent, DO NOT EDIT.

package ent

import (
	"gas-detect/internal/data/model/ent/gasdetect"
	"gas-detect/internal/data/model/ent/predicate"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeGasDetect = "GasDetect"
)

// GasDetectMutation represents an operation that mutates the GasDetect nodes in the graph.
type GasDetectMutation struct {
	config
	op            Op
	typ           string
	id            *int32
	created_time  *time.Time
	is_deleted    *bool
	metrics       *[]map[string]string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*GasDetect, error)
	predicates    []predicate.GasDetect
}

var _ ent.Mutation = (*GasDetectMutation)(nil)

// gasdetectOption allows management of the mutation configuration using functional options.
type gasdetectOption func(*GasDetectMutation)

// newGasDetectMutation creates new mutation for the GasDetect entity.
func newGasDetectMutation(c config, op Op, opts ...gasdetectOption) *GasDetectMutation {
	m := &GasDetectMutation{
		config:        c,
		op:            op,
		typ:           TypeGasDetect,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withGasDetectID sets the ID field of the mutation.
func withGasDetectID(id int32) gasdetectOption {
	return func(m *GasDetectMutation) {
		var (
			err   error
			once  sync.Once
			value *GasDetect
		)
		m.oldValue = func(ctx context.Context) (*GasDetect, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().GasDetect.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withGasDetect sets the old GasDetect of the mutation.
func withGasDetect(node *GasDetect) gasdetectOption {
	return func(m *GasDetectMutation) {
		m.oldValue = func(context.Context) (*GasDetect, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m GasDetectMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m GasDetectMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of GasDetect entities.
func (m *GasDetectMutation) SetID(id int32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *GasDetectMutation) ID() (id int32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *GasDetectMutation) IDs(ctx context.Context) ([]int32, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int32{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().GasDetect.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedTime sets the "created_time" field.
func (m *GasDetectMutation) SetCreatedTime(t time.Time) {
	m.created_time = &t
}

// CreatedTime returns the value of the "created_time" field in the mutation.
func (m *GasDetectMutation) CreatedTime() (r time.Time, exists bool) {
	v := m.created_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedTime returns the old "created_time" field's value of the GasDetect entity.
// If the GasDetect object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GasDetectMutation) OldCreatedTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedTime: %w", err)
	}
	return oldValue.CreatedTime, nil
}

// ResetCreatedTime resets all changes to the "created_time" field.
func (m *GasDetectMutation) ResetCreatedTime() {
	m.created_time = nil
}

// SetIsDeleted sets the "is_deleted" field.
func (m *GasDetectMutation) SetIsDeleted(b bool) {
	m.is_deleted = &b
}

// IsDeleted returns the value of the "is_deleted" field in the mutation.
func (m *GasDetectMutation) IsDeleted() (r bool, exists bool) {
	v := m.is_deleted
	if v == nil {
		return
	}
	return *v, true
}

// OldIsDeleted returns the old "is_deleted" field's value of the GasDetect entity.
// If the GasDetect object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GasDetectMutation) OldIsDeleted(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIsDeleted is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIsDeleted requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIsDeleted: %w", err)
	}
	return oldValue.IsDeleted, nil
}

// ResetIsDeleted resets all changes to the "is_deleted" field.
func (m *GasDetectMutation) ResetIsDeleted() {
	m.is_deleted = nil
}

// SetMetrics sets the "metrics" field.
func (m *GasDetectMutation) SetMetrics(value []map[string]string) {
	m.metrics = &value
}

// Metrics returns the value of the "metrics" field in the mutation.
func (m *GasDetectMutation) Metrics() (r []map[string]string, exists bool) {
	v := m.metrics
	if v == nil {
		return
	}
	return *v, true
}

// OldMetrics returns the old "metrics" field's value of the GasDetect entity.
// If the GasDetect object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *GasDetectMutation) OldMetrics(ctx context.Context) (v []map[string]string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMetrics is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMetrics requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMetrics: %w", err)
	}
	return oldValue.Metrics, nil
}

// ClearMetrics clears the value of the "metrics" field.
func (m *GasDetectMutation) ClearMetrics() {
	m.metrics = nil
	m.clearedFields[gasdetect.FieldMetrics] = struct{}{}
}

// MetricsCleared returns if the "metrics" field was cleared in this mutation.
func (m *GasDetectMutation) MetricsCleared() bool {
	_, ok := m.clearedFields[gasdetect.FieldMetrics]
	return ok
}

// ResetMetrics resets all changes to the "metrics" field.
func (m *GasDetectMutation) ResetMetrics() {
	m.metrics = nil
	delete(m.clearedFields, gasdetect.FieldMetrics)
}

// Where appends a list predicates to the GasDetectMutation builder.
func (m *GasDetectMutation) Where(ps ...predicate.GasDetect) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *GasDetectMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (GasDetect).
func (m *GasDetectMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *GasDetectMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.created_time != nil {
		fields = append(fields, gasdetect.FieldCreatedTime)
	}
	if m.is_deleted != nil {
		fields = append(fields, gasdetect.FieldIsDeleted)
	}
	if m.metrics != nil {
		fields = append(fields, gasdetect.FieldMetrics)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *GasDetectMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case gasdetect.FieldCreatedTime:
		return m.CreatedTime()
	case gasdetect.FieldIsDeleted:
		return m.IsDeleted()
	case gasdetect.FieldMetrics:
		return m.Metrics()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *GasDetectMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case gasdetect.FieldCreatedTime:
		return m.OldCreatedTime(ctx)
	case gasdetect.FieldIsDeleted:
		return m.OldIsDeleted(ctx)
	case gasdetect.FieldMetrics:
		return m.OldMetrics(ctx)
	}
	return nil, fmt.Errorf("unknown GasDetect field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GasDetectMutation) SetField(name string, value ent.Value) error {
	switch name {
	case gasdetect.FieldCreatedTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedTime(v)
		return nil
	case gasdetect.FieldIsDeleted:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIsDeleted(v)
		return nil
	case gasdetect.FieldMetrics:
		v, ok := value.([]map[string]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMetrics(v)
		return nil
	}
	return fmt.Errorf("unknown GasDetect field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *GasDetectMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *GasDetectMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *GasDetectMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown GasDetect numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *GasDetectMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(gasdetect.FieldMetrics) {
		fields = append(fields, gasdetect.FieldMetrics)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *GasDetectMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *GasDetectMutation) ClearField(name string) error {
	switch name {
	case gasdetect.FieldMetrics:
		m.ClearMetrics()
		return nil
	}
	return fmt.Errorf("unknown GasDetect nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *GasDetectMutation) ResetField(name string) error {
	switch name {
	case gasdetect.FieldCreatedTime:
		m.ResetCreatedTime()
		return nil
	case gasdetect.FieldIsDeleted:
		m.ResetIsDeleted()
		return nil
	case gasdetect.FieldMetrics:
		m.ResetMetrics()
		return nil
	}
	return fmt.Errorf("unknown GasDetect field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *GasDetectMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *GasDetectMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *GasDetectMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *GasDetectMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *GasDetectMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *GasDetectMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *GasDetectMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown GasDetect unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *GasDetectMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown GasDetect edge %s", name)
}
