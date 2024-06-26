// Code generated by ent, DO NOT EDIT.

package ent

import (
	"RjsConfigService/ent/appconfig"
	"RjsConfigService/ent/predicate"
	"RjsConfigService/ent/user"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAppConfig = "AppConfig"
	TypeUser      = "User"
)

// AppConfigMutation represents an operation that mutates the AppConfig nodes in the graph.
type AppConfigMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	app_name      *string
	version       *int
	addversion    *int
	status        *string
	_config       *map[string]interface{}
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*AppConfig, error)
	predicates    []predicate.AppConfig
}

var _ ent.Mutation = (*AppConfigMutation)(nil)

// appconfigOption allows management of the mutation configuration using functional options.
type appconfigOption func(*AppConfigMutation)

// newAppConfigMutation creates new mutation for the AppConfig entity.
func newAppConfigMutation(c config, op Op, opts ...appconfigOption) *AppConfigMutation {
	m := &AppConfigMutation{
		config:        c,
		op:            op,
		typ:           TypeAppConfig,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAppConfigID sets the ID field of the mutation.
func withAppConfigID(id uuid.UUID) appconfigOption {
	return func(m *AppConfigMutation) {
		var (
			err   error
			once  sync.Once
			value *AppConfig
		)
		m.oldValue = func(ctx context.Context) (*AppConfig, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AppConfig.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAppConfig sets the old AppConfig of the mutation.
func withAppConfig(node *AppConfig) appconfigOption {
	return func(m *AppConfigMutation) {
		m.oldValue = func(context.Context) (*AppConfig, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AppConfigMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AppConfigMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of AppConfig entities.
func (m *AppConfigMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AppConfigMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AppConfigMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().AppConfig.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *AppConfigMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *AppConfigMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the AppConfig entity.
// If the AppConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AppConfigMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *AppConfigMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetAppName sets the "app_name" field.
func (m *AppConfigMutation) SetAppName(s string) {
	m.app_name = &s
}

// AppName returns the value of the "app_name" field in the mutation.
func (m *AppConfigMutation) AppName() (r string, exists bool) {
	v := m.app_name
	if v == nil {
		return
	}
	return *v, true
}

// OldAppName returns the old "app_name" field's value of the AppConfig entity.
// If the AppConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AppConfigMutation) OldAppName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAppName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAppName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAppName: %w", err)
	}
	return oldValue.AppName, nil
}

// ResetAppName resets all changes to the "app_name" field.
func (m *AppConfigMutation) ResetAppName() {
	m.app_name = nil
}

// SetVersion sets the "version" field.
func (m *AppConfigMutation) SetVersion(i int) {
	m.version = &i
	m.addversion = nil
}

// Version returns the value of the "version" field in the mutation.
func (m *AppConfigMutation) Version() (r int, exists bool) {
	v := m.version
	if v == nil {
		return
	}
	return *v, true
}

// OldVersion returns the old "version" field's value of the AppConfig entity.
// If the AppConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AppConfigMutation) OldVersion(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldVersion is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldVersion requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldVersion: %w", err)
	}
	return oldValue.Version, nil
}

// AddVersion adds i to the "version" field.
func (m *AppConfigMutation) AddVersion(i int) {
	if m.addversion != nil {
		*m.addversion += i
	} else {
		m.addversion = &i
	}
}

// AddedVersion returns the value that was added to the "version" field in this mutation.
func (m *AppConfigMutation) AddedVersion() (r int, exists bool) {
	v := m.addversion
	if v == nil {
		return
	}
	return *v, true
}

// ResetVersion resets all changes to the "version" field.
func (m *AppConfigMutation) ResetVersion() {
	m.version = nil
	m.addversion = nil
}

// SetStatus sets the "status" field.
func (m *AppConfigMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *AppConfigMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the AppConfig entity.
// If the AppConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AppConfigMutation) OldStatus(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *AppConfigMutation) ResetStatus() {
	m.status = nil
}

// SetConfig sets the "config" field.
func (m *AppConfigMutation) SetConfig(value map[string]interface{}) {
	m._config = &value
}

// Config returns the value of the "config" field in the mutation.
func (m *AppConfigMutation) Config() (r map[string]interface{}, exists bool) {
	v := m._config
	if v == nil {
		return
	}
	return *v, true
}

// OldConfig returns the old "config" field's value of the AppConfig entity.
// If the AppConfig object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AppConfigMutation) OldConfig(ctx context.Context) (v map[string]interface{}, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldConfig is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldConfig requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldConfig: %w", err)
	}
	return oldValue.Config, nil
}

// ResetConfig resets all changes to the "config" field.
func (m *AppConfigMutation) ResetConfig() {
	m._config = nil
}

// Where appends a list predicates to the AppConfigMutation builder.
func (m *AppConfigMutation) Where(ps ...predicate.AppConfig) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AppConfigMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AppConfigMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.AppConfig, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AppConfigMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AppConfigMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (AppConfig).
func (m *AppConfigMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AppConfigMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, appconfig.FieldCreatedAt)
	}
	if m.app_name != nil {
		fields = append(fields, appconfig.FieldAppName)
	}
	if m.version != nil {
		fields = append(fields, appconfig.FieldVersion)
	}
	if m.status != nil {
		fields = append(fields, appconfig.FieldStatus)
	}
	if m._config != nil {
		fields = append(fields, appconfig.FieldConfig)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AppConfigMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case appconfig.FieldCreatedAt:
		return m.CreatedAt()
	case appconfig.FieldAppName:
		return m.AppName()
	case appconfig.FieldVersion:
		return m.Version()
	case appconfig.FieldStatus:
		return m.Status()
	case appconfig.FieldConfig:
		return m.Config()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AppConfigMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case appconfig.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case appconfig.FieldAppName:
		return m.OldAppName(ctx)
	case appconfig.FieldVersion:
		return m.OldVersion(ctx)
	case appconfig.FieldStatus:
		return m.OldStatus(ctx)
	case appconfig.FieldConfig:
		return m.OldConfig(ctx)
	}
	return nil, fmt.Errorf("unknown AppConfig field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AppConfigMutation) SetField(name string, value ent.Value) error {
	switch name {
	case appconfig.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case appconfig.FieldAppName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAppName(v)
		return nil
	case appconfig.FieldVersion:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetVersion(v)
		return nil
	case appconfig.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case appconfig.FieldConfig:
		v, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetConfig(v)
		return nil
	}
	return fmt.Errorf("unknown AppConfig field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AppConfigMutation) AddedFields() []string {
	var fields []string
	if m.addversion != nil {
		fields = append(fields, appconfig.FieldVersion)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AppConfigMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case appconfig.FieldVersion:
		return m.AddedVersion()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AppConfigMutation) AddField(name string, value ent.Value) error {
	switch name {
	case appconfig.FieldVersion:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddVersion(v)
		return nil
	}
	return fmt.Errorf("unknown AppConfig numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AppConfigMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AppConfigMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AppConfigMutation) ClearField(name string) error {
	return fmt.Errorf("unknown AppConfig nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AppConfigMutation) ResetField(name string) error {
	switch name {
	case appconfig.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case appconfig.FieldAppName:
		m.ResetAppName()
		return nil
	case appconfig.FieldVersion:
		m.ResetVersion()
		return nil
	case appconfig.FieldStatus:
		m.ResetStatus()
		return nil
	case appconfig.FieldConfig:
		m.ResetConfig()
		return nil
	}
	return fmt.Errorf("unknown AppConfig field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AppConfigMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AppConfigMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AppConfigMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AppConfigMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AppConfigMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AppConfigMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AppConfigMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AppConfig unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AppConfigMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AppConfig edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	created_at    *time.Time
	role          *string
	app_name      *string
	status        *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetRole sets the "role" field.
func (m *UserMutation) SetRole(s string) {
	m.role = &s
}

// Role returns the value of the "role" field in the mutation.
func (m *UserMutation) Role() (r string, exists bool) {
	v := m.role
	if v == nil {
		return
	}
	return *v, true
}

// OldRole returns the old "role" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldRole(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRole is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRole requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRole: %w", err)
	}
	return oldValue.Role, nil
}

// ResetRole resets all changes to the "role" field.
func (m *UserMutation) ResetRole() {
	m.role = nil
}

// SetAppName sets the "app_name" field.
func (m *UserMutation) SetAppName(s string) {
	m.app_name = &s
}

// AppName returns the value of the "app_name" field in the mutation.
func (m *UserMutation) AppName() (r string, exists bool) {
	v := m.app_name
	if v == nil {
		return
	}
	return *v, true
}

// OldAppName returns the old "app_name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAppName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAppName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAppName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAppName: %w", err)
	}
	return oldValue.AppName, nil
}

// ResetAppName resets all changes to the "app_name" field.
func (m *UserMutation) ResetAppName() {
	m.app_name = nil
}

// SetStatus sets the "status" field.
func (m *UserMutation) SetStatus(s string) {
	m.status = &s
}

// Status returns the value of the "status" field in the mutation.
func (m *UserMutation) Status() (r string, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldStatus(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *UserMutation) ResetStatus() {
	m.status = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.role != nil {
		fields = append(fields, user.FieldRole)
	}
	if m.app_name != nil {
		fields = append(fields, user.FieldAppName)
	}
	if m.status != nil {
		fields = append(fields, user.FieldStatus)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldRole:
		return m.Role()
	case user.FieldAppName:
		return m.AppName()
	case user.FieldStatus:
		return m.Status()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldRole:
		return m.OldRole(ctx)
	case user.FieldAppName:
		return m.OldAppName(ctx)
	case user.FieldStatus:
		return m.OldStatus(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldRole:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRole(v)
		return nil
	case user.FieldAppName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAppName(v)
		return nil
	case user.FieldStatus:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldRole:
		m.ResetRole()
		return nil
	case user.FieldAppName:
		m.ResetAppName()
		return nil
	case user.FieldStatus:
		m.ResetStatus()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
