// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/predicate"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// SetHobbies sets the "hobbies" field.
func (uu *UserUpdate) SetHobbies(s []string) *UserUpdate {
	uu.mutation.SetHobbies(s)
	return uu
}

// AppendHobbies appends s to the "hobbies" field.
func (uu *UserUpdate) AppendHobbies(s []string) *UserUpdate {
	uu.mutation.AppendHobbies(s)
	return uu
}

// ClearHobbies clears the value of the "hobbies" field.
func (uu *UserUpdate) ClearHobbies() *UserUpdate {
	uu.mutation.ClearHobbies()
	return uu
}

// SetCreatedAt sets the "created_at" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// SetCreatedBy sets the "created_by" field.
func (uu *UserUpdate) SetCreatedBy(u uuid.UUID) *UserUpdate {
	uu.mutation.SetCreatedBy(u)
	return uu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedBy(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetCreatedBy(*u)
	}
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetUpdatedAt(*t)
	}
	return uu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uu *UserUpdate) ClearUpdatedAt() *UserUpdate {
	uu.mutation.ClearUpdatedAt()
	return uu
}

// SetUpdatedBy sets the "updated_by" field.
func (uu *UserUpdate) SetUpdatedBy(u uuid.UUID) *UserUpdate {
	uu.mutation.SetUpdatedBy(u)
	return uu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdatedBy(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetUpdatedBy(*u)
	}
	return uu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uu *UserUpdate) ClearUpdatedBy() *UserUpdate {
	uu.mutation.ClearUpdatedBy()
	return uu
}

// SetDeletedAt sets the "deleted_at" field.
func (uu *UserUpdate) SetDeletedAt(t time.Time) *UserUpdate {
	uu.mutation.SetDeletedAt(t)
	return uu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetDeletedAt(*t)
	}
	return uu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uu *UserUpdate) ClearDeletedAt() *UserUpdate {
	uu.mutation.ClearDeletedAt()
	return uu
}

// SetDeletedBy sets the "deleted_by" field.
func (uu *UserUpdate) SetDeletedBy(u uuid.UUID) *UserUpdate {
	uu.mutation.SetDeletedBy(u)
	return uu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDeletedBy(u *uuid.UUID) *UserUpdate {
	if u != nil {
		uu.SetDeletedBy(*u)
	}
	return uu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uu *UserUpdate) ClearDeletedBy() *UserUpdate {
	uu.mutation.ClearDeletedBy()
	return uu
}

// SetIsDeleted sets the "is_deleted" field.
func (uu *UserUpdate) SetIsDeleted(i int64) *UserUpdate {
	uu.mutation.ResetIsDeleted()
	uu.mutation.SetIsDeleted(i)
	return uu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsDeleted(i *int64) *UserUpdate {
	if i != nil {
		uu.SetIsDeleted(*i)
	}
	return uu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (uu *UserUpdate) AddIsDeleted(i int64) *UserUpdate {
	uu.mutation.AddIsDeleted(i)
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`generated: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.Hobbies(); ok {
		_spec.SetField(user.FieldHobbies, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedHobbies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldHobbies, value)
		})
	}
	if uu.mutation.HobbiesCleared() {
		_spec.ClearField(user.FieldHobbies, field.TypeJSON)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uu.mutation.UpdatedAtCleared() {
		_spec.ClearField(user.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeUUID, value)
	}
	if uu.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := uu.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uu.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uu.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeUUID, value)
	}
	if uu.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := uu.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedIsDeleted(); ok {
		_spec.AddField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// SetHobbies sets the "hobbies" field.
func (uuo *UserUpdateOne) SetHobbies(s []string) *UserUpdateOne {
	uuo.mutation.SetHobbies(s)
	return uuo
}

// AppendHobbies appends s to the "hobbies" field.
func (uuo *UserUpdateOne) AppendHobbies(s []string) *UserUpdateOne {
	uuo.mutation.AppendHobbies(s)
	return uuo
}

// ClearHobbies clears the value of the "hobbies" field.
func (uuo *UserUpdateOne) ClearHobbies() *UserUpdateOne {
	uuo.mutation.ClearHobbies()
	return uuo
}

// SetCreatedAt sets the "created_at" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// SetCreatedBy sets the "created_by" field.
func (uuo *UserUpdateOne) SetCreatedBy(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetCreatedBy(u)
	return uuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetCreatedBy(*u)
	}
	return uuo
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetUpdatedAt(*t)
	}
	return uuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (uuo *UserUpdateOne) ClearUpdatedAt() *UserUpdateOne {
	uuo.mutation.ClearUpdatedAt()
	return uuo
}

// SetUpdatedBy sets the "updated_by" field.
func (uuo *UserUpdateOne) SetUpdatedBy(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetUpdatedBy(u)
	return uuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetUpdatedBy(*u)
	}
	return uuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (uuo *UserUpdateOne) ClearUpdatedBy() *UserUpdateOne {
	uuo.mutation.ClearUpdatedBy()
	return uuo
}

// SetDeletedAt sets the "deleted_at" field.
func (uuo *UserUpdateOne) SetDeletedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetDeletedAt(t)
	return uuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetDeletedAt(*t)
	}
	return uuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (uuo *UserUpdateOne) ClearDeletedAt() *UserUpdateOne {
	uuo.mutation.ClearDeletedAt()
	return uuo
}

// SetDeletedBy sets the "deleted_by" field.
func (uuo *UserUpdateOne) SetDeletedBy(u uuid.UUID) *UserUpdateOne {
	uuo.mutation.SetDeletedBy(u)
	return uuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDeletedBy(u *uuid.UUID) *UserUpdateOne {
	if u != nil {
		uuo.SetDeletedBy(*u)
	}
	return uuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (uuo *UserUpdateOne) ClearDeletedBy() *UserUpdateOne {
	uuo.mutation.ClearDeletedBy()
	return uuo
}

// SetIsDeleted sets the "is_deleted" field.
func (uuo *UserUpdateOne) SetIsDeleted(i int64) *UserUpdateOne {
	uuo.mutation.ResetIsDeleted()
	uuo.mutation.SetIsDeleted(i)
	return uuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsDeleted(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetIsDeleted(*i)
	}
	return uuo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (uuo *UserUpdateOne) AddIsDeleted(i int64) *UserUpdateOne {
	uuo.mutation.AddIsDeleted(i)
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Name(); ok {
		if err := user.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "User.name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`generated: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.Hobbies(); ok {
		_spec.SetField(user.FieldHobbies, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedHobbies(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldHobbies, value)
		})
	}
	if uuo.mutation.HobbiesCleared() {
		_spec.ClearField(user.FieldHobbies, field.TypeJSON)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if uuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(user.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeUUID, value)
	}
	if uuo.mutation.UpdatedByCleared() {
		_spec.ClearField(user.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := uuo.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeTime, value)
	}
	if uuo.mutation.DeletedAtCleared() {
		_spec.ClearField(user.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := uuo.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeUUID, value)
	}
	if uuo.mutation.DeletedByCleared() {
		_spec.ClearField(user.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := uuo.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedIsDeleted(); ok {
		_spec.AddField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
