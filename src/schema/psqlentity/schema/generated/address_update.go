// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/address"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/predicate"
)

// AddressUpdate is the builder for updating Address entities.
type AddressUpdate struct {
	config
	hooks    []Hook
	mutation *AddressMutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (au *AddressUpdate) Where(ps ...predicate.Address) *AddressUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCountry sets the "country" field.
func (au *AddressUpdate) SetCountry(s string) *AddressUpdate {
	au.mutation.SetCountry(s)
	return au
}

// SetProvince sets the "province" field.
func (au *AddressUpdate) SetProvince(s string) *AddressUpdate {
	au.mutation.SetProvince(s)
	return au
}

// SetRegency sets the "regency" field.
func (au *AddressUpdate) SetRegency(s string) *AddressUpdate {
	au.mutation.SetRegency(s)
	return au
}

// SetSubDistrict sets the "sub_district" field.
func (au *AddressUpdate) SetSubDistrict(s string) *AddressUpdate {
	au.mutation.SetSubDistrict(s)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AddressUpdate) SetCreatedAt(t time.Time) *AddressUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AddressUpdate) SetNillableCreatedAt(t *time.Time) *AddressUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetCreatedBy sets the "created_by" field.
func (au *AddressUpdate) SetCreatedBy(u uuid.UUID) *AddressUpdate {
	au.mutation.SetCreatedBy(u)
	return au
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (au *AddressUpdate) SetNillableCreatedBy(u *uuid.UUID) *AddressUpdate {
	if u != nil {
		au.SetCreatedBy(*u)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AddressUpdate) SetUpdatedAt(t time.Time) *AddressUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (au *AddressUpdate) SetNillableUpdatedAt(t *time.Time) *AddressUpdate {
	if t != nil {
		au.SetUpdatedAt(*t)
	}
	return au
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (au *AddressUpdate) ClearUpdatedAt() *AddressUpdate {
	au.mutation.ClearUpdatedAt()
	return au
}

// SetUpdatedBy sets the "updated_by" field.
func (au *AddressUpdate) SetUpdatedBy(u uuid.UUID) *AddressUpdate {
	au.mutation.SetUpdatedBy(u)
	return au
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (au *AddressUpdate) SetNillableUpdatedBy(u *uuid.UUID) *AddressUpdate {
	if u != nil {
		au.SetUpdatedBy(*u)
	}
	return au
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (au *AddressUpdate) ClearUpdatedBy() *AddressUpdate {
	au.mutation.ClearUpdatedBy()
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AddressUpdate) SetDeletedAt(t time.Time) *AddressUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AddressUpdate) SetNillableDeletedAt(t *time.Time) *AddressUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *AddressUpdate) ClearDeletedAt() *AddressUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// SetDeletedBy sets the "deleted_by" field.
func (au *AddressUpdate) SetDeletedBy(u uuid.UUID) *AddressUpdate {
	au.mutation.SetDeletedBy(u)
	return au
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (au *AddressUpdate) SetNillableDeletedBy(u *uuid.UUID) *AddressUpdate {
	if u != nil {
		au.SetDeletedBy(*u)
	}
	return au
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (au *AddressUpdate) ClearDeletedBy() *AddressUpdate {
	au.mutation.ClearDeletedBy()
	return au
}

// SetIsDeleted sets the "is_deleted" field.
func (au *AddressUpdate) SetIsDeleted(i int) *AddressUpdate {
	au.mutation.ResetIsDeleted()
	au.mutation.SetIsDeleted(i)
	return au
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (au *AddressUpdate) SetNillableIsDeleted(i *int) *AddressUpdate {
	if i != nil {
		au.SetIsDeleted(*i)
	}
	return au
}

// AddIsDeleted adds i to the "is_deleted" field.
func (au *AddressUpdate) AddIsDeleted(i int) *AddressUpdate {
	au.mutation.AddIsDeleted(i)
	return au
}

// Mutation returns the AddressMutation object of the builder.
func (au *AddressUpdate) Mutation() *AddressMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AddressUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AddressUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AddressUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AddressUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AddressUpdate) check() error {
	if v, ok := au.mutation.Country(); ok {
		if err := address.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`generated: validator failed for field "Address.country": %w`, err)}
		}
	}
	if v, ok := au.mutation.Province(); ok {
		if err := address.ProvinceValidator(v); err != nil {
			return &ValidationError{Name: "province", err: fmt.Errorf(`generated: validator failed for field "Address.province": %w`, err)}
		}
	}
	if v, ok := au.mutation.Regency(); ok {
		if err := address.RegencyValidator(v); err != nil {
			return &ValidationError{Name: "regency", err: fmt.Errorf(`generated: validator failed for field "Address.regency": %w`, err)}
		}
	}
	if v, ok := au.mutation.SubDistrict(); ok {
		if err := address.SubDistrictValidator(v); err != nil {
			return &ValidationError{Name: "sub_district", err: fmt.Errorf(`generated: validator failed for field "Address.sub_district": %w`, err)}
		}
	}
	if v, ok := au.mutation.IsDeleted(); ok {
		if err := address.IsDeletedValidator(v); err != nil {
			return &ValidationError{Name: "is_deleted", err: fmt.Errorf(`generated: validator failed for field "Address.is_deleted": %w`, err)}
		}
	}
	return nil
}

func (au *AddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
	}
	if value, ok := au.mutation.Province(); ok {
		_spec.SetField(address.FieldProvince, field.TypeString, value)
	}
	if value, ok := au.mutation.Regency(); ok {
		_spec.SetField(address.FieldRegency, field.TypeString, value)
	}
	if value, ok := au.mutation.SubDistrict(); ok {
		_spec.SetField(address.FieldSubDistrict, field.TypeString, value)
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(address.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.CreatedBy(); ok {
		_spec.SetField(address.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(address.FieldUpdatedAt, field.TypeTime, value)
	}
	if au.mutation.UpdatedAtCleared() {
		_spec.ClearField(address.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := au.mutation.UpdatedBy(); ok {
		_spec.SetField(address.FieldUpdatedBy, field.TypeUUID, value)
	}
	if au.mutation.UpdatedByCleared() {
		_spec.ClearField(address.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(address.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(address.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := au.mutation.DeletedBy(); ok {
		_spec.SetField(address.FieldDeletedBy, field.TypeUUID, value)
	}
	if au.mutation.DeletedByCleared() {
		_spec.ClearField(address.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := au.mutation.IsDeleted(); ok {
		_spec.SetField(address.FieldIsDeleted, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedIsDeleted(); ok {
		_spec.AddField(address.FieldIsDeleted, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AddressUpdateOne is the builder for updating a single Address entity.
type AddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AddressMutation
}

// SetCountry sets the "country" field.
func (auo *AddressUpdateOne) SetCountry(s string) *AddressUpdateOne {
	auo.mutation.SetCountry(s)
	return auo
}

// SetProvince sets the "province" field.
func (auo *AddressUpdateOne) SetProvince(s string) *AddressUpdateOne {
	auo.mutation.SetProvince(s)
	return auo
}

// SetRegency sets the "regency" field.
func (auo *AddressUpdateOne) SetRegency(s string) *AddressUpdateOne {
	auo.mutation.SetRegency(s)
	return auo
}

// SetSubDistrict sets the "sub_district" field.
func (auo *AddressUpdateOne) SetSubDistrict(s string) *AddressUpdateOne {
	auo.mutation.SetSubDistrict(s)
	return auo
}

// SetCreatedAt sets the "created_at" field.
func (auo *AddressUpdateOne) SetCreatedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCreatedAt(t *time.Time) *AddressUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetCreatedBy sets the "created_by" field.
func (auo *AddressUpdateOne) SetCreatedBy(u uuid.UUID) *AddressUpdateOne {
	auo.mutation.SetCreatedBy(u)
	return auo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *AddressUpdateOne {
	if u != nil {
		auo.SetCreatedBy(*u)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AddressUpdateOne) SetUpdatedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableUpdatedAt(t *time.Time) *AddressUpdateOne {
	if t != nil {
		auo.SetUpdatedAt(*t)
	}
	return auo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (auo *AddressUpdateOne) ClearUpdatedAt() *AddressUpdateOne {
	auo.mutation.ClearUpdatedAt()
	return auo
}

// SetUpdatedBy sets the "updated_by" field.
func (auo *AddressUpdateOne) SetUpdatedBy(u uuid.UUID) *AddressUpdateOne {
	auo.mutation.SetUpdatedBy(u)
	return auo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *AddressUpdateOne {
	if u != nil {
		auo.SetUpdatedBy(*u)
	}
	return auo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (auo *AddressUpdateOne) ClearUpdatedBy() *AddressUpdateOne {
	auo.mutation.ClearUpdatedBy()
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AddressUpdateOne) SetDeletedAt(t time.Time) *AddressUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableDeletedAt(t *time.Time) *AddressUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *AddressUpdateOne) ClearDeletedAt() *AddressUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// SetDeletedBy sets the "deleted_by" field.
func (auo *AddressUpdateOne) SetDeletedBy(u uuid.UUID) *AddressUpdateOne {
	auo.mutation.SetDeletedBy(u)
	return auo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableDeletedBy(u *uuid.UUID) *AddressUpdateOne {
	if u != nil {
		auo.SetDeletedBy(*u)
	}
	return auo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (auo *AddressUpdateOne) ClearDeletedBy() *AddressUpdateOne {
	auo.mutation.ClearDeletedBy()
	return auo
}

// SetIsDeleted sets the "is_deleted" field.
func (auo *AddressUpdateOne) SetIsDeleted(i int) *AddressUpdateOne {
	auo.mutation.ResetIsDeleted()
	auo.mutation.SetIsDeleted(i)
	return auo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (auo *AddressUpdateOne) SetNillableIsDeleted(i *int) *AddressUpdateOne {
	if i != nil {
		auo.SetIsDeleted(*i)
	}
	return auo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (auo *AddressUpdateOne) AddIsDeleted(i int) *AddressUpdateOne {
	auo.mutation.AddIsDeleted(i)
	return auo
}

// Mutation returns the AddressMutation object of the builder.
func (auo *AddressUpdateOne) Mutation() *AddressMutation {
	return auo.mutation
}

// Where appends a list predicates to the AddressUpdate builder.
func (auo *AddressUpdateOne) Where(ps ...predicate.Address) *AddressUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AddressUpdateOne) Select(field string, fields ...string) *AddressUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Address entity.
func (auo *AddressUpdateOne) Save(ctx context.Context) (*Address, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AddressUpdateOne) SaveX(ctx context.Context) *Address {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AddressUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AddressUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AddressUpdateOne) check() error {
	if v, ok := auo.mutation.Country(); ok {
		if err := address.CountryValidator(v); err != nil {
			return &ValidationError{Name: "country", err: fmt.Errorf(`generated: validator failed for field "Address.country": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Province(); ok {
		if err := address.ProvinceValidator(v); err != nil {
			return &ValidationError{Name: "province", err: fmt.Errorf(`generated: validator failed for field "Address.province": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Regency(); ok {
		if err := address.RegencyValidator(v); err != nil {
			return &ValidationError{Name: "regency", err: fmt.Errorf(`generated: validator failed for field "Address.regency": %w`, err)}
		}
	}
	if v, ok := auo.mutation.SubDistrict(); ok {
		if err := address.SubDistrictValidator(v); err != nil {
			return &ValidationError{Name: "sub_district", err: fmt.Errorf(`generated: validator failed for field "Address.sub_district": %w`, err)}
		}
	}
	if v, ok := auo.mutation.IsDeleted(); ok {
		if err := address.IsDeletedValidator(v); err != nil {
			return &ValidationError{Name: "is_deleted", err: fmt.Errorf(`generated: validator failed for field "Address.is_deleted": %w`, err)}
		}
	}
	return nil
}

func (auo *AddressUpdateOne) sqlSave(ctx context.Context) (_node *Address, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Address.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, address.FieldID)
		for _, f := range fields {
			if !address.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != address.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Country(); ok {
		_spec.SetField(address.FieldCountry, field.TypeString, value)
	}
	if value, ok := auo.mutation.Province(); ok {
		_spec.SetField(address.FieldProvince, field.TypeString, value)
	}
	if value, ok := auo.mutation.Regency(); ok {
		_spec.SetField(address.FieldRegency, field.TypeString, value)
	}
	if value, ok := auo.mutation.SubDistrict(); ok {
		_spec.SetField(address.FieldSubDistrict, field.TypeString, value)
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(address.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.CreatedBy(); ok {
		_spec.SetField(address.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(address.FieldUpdatedAt, field.TypeTime, value)
	}
	if auo.mutation.UpdatedAtCleared() {
		_spec.ClearField(address.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.UpdatedBy(); ok {
		_spec.SetField(address.FieldUpdatedBy, field.TypeUUID, value)
	}
	if auo.mutation.UpdatedByCleared() {
		_spec.ClearField(address.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(address.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(address.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := auo.mutation.DeletedBy(); ok {
		_spec.SetField(address.FieldDeletedBy, field.TypeUUID, value)
	}
	if auo.mutation.DeletedByCleared() {
		_spec.ClearField(address.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := auo.mutation.IsDeleted(); ok {
		_spec.SetField(address.FieldIsDeleted, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedIsDeleted(); ok {
		_spec.AddField(address.FieldIsDeleted, field.TypeInt, value)
	}
	_node = &Address{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{address.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}