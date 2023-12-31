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
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/predicate"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/purchase"
)

// PurchaseUpdate is the builder for updating Purchase entities.
type PurchaseUpdate struct {
	config
	hooks    []Hook
	mutation *PurchaseMutation
}

// Where appends a list predicates to the PurchaseUpdate builder.
func (pu *PurchaseUpdate) Where(ps ...predicate.Purchase) *PurchaseUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetProductID sets the "product_id" field.
func (pu *PurchaseUpdate) SetProductID(u uuid.UUID) *PurchaseUpdate {
	pu.mutation.SetProductID(u)
	return pu
}

// SetSupplierID sets the "supplier_id" field.
func (pu *PurchaseUpdate) SetSupplierID(u uuid.UUID) *PurchaseUpdate {
	pu.mutation.SetSupplierID(u)
	return pu
}

// SetQuantity sets the "quantity" field.
func (pu *PurchaseUpdate) SetQuantity(i int64) *PurchaseUpdate {
	pu.mutation.ResetQuantity()
	pu.mutation.SetQuantity(i)
	return pu
}

// AddQuantity adds i to the "quantity" field.
func (pu *PurchaseUpdate) AddQuantity(i int64) *PurchaseUpdate {
	pu.mutation.AddQuantity(i)
	return pu
}

// SetDate sets the "date" field.
func (pu *PurchaseUpdate) SetDate(t time.Time) *PurchaseUpdate {
	pu.mutation.SetDate(t)
	return pu
}

// SetTotalAmount sets the "total_amount" field.
func (pu *PurchaseUpdate) SetTotalAmount(i int64) *PurchaseUpdate {
	pu.mutation.ResetTotalAmount()
	pu.mutation.SetTotalAmount(i)
	return pu
}

// AddTotalAmount adds i to the "total_amount" field.
func (pu *PurchaseUpdate) AddTotalAmount(i int64) *PurchaseUpdate {
	pu.mutation.AddTotalAmount(i)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PurchaseUpdate) SetCreatedAt(t time.Time) *PurchaseUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableCreatedAt(t *time.Time) *PurchaseUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetCreatedBy sets the "created_by" field.
func (pu *PurchaseUpdate) SetCreatedBy(u uuid.UUID) *PurchaseUpdate {
	pu.mutation.SetCreatedBy(u)
	return pu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableCreatedBy(u *uuid.UUID) *PurchaseUpdate {
	if u != nil {
		pu.SetCreatedBy(*u)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PurchaseUpdate) SetUpdatedAt(t time.Time) *PurchaseUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableUpdatedAt(t *time.Time) *PurchaseUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PurchaseUpdate) ClearUpdatedAt() *PurchaseUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetUpdatedBy sets the "updated_by" field.
func (pu *PurchaseUpdate) SetUpdatedBy(u uuid.UUID) *PurchaseUpdate {
	pu.mutation.SetUpdatedBy(u)
	return pu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableUpdatedBy(u *uuid.UUID) *PurchaseUpdate {
	if u != nil {
		pu.SetUpdatedBy(*u)
	}
	return pu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (pu *PurchaseUpdate) ClearUpdatedBy() *PurchaseUpdate {
	pu.mutation.ClearUpdatedBy()
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PurchaseUpdate) SetDeletedAt(t time.Time) *PurchaseUpdate {
	pu.mutation.SetDeletedAt(t)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableDeletedAt(t *time.Time) *PurchaseUpdate {
	if t != nil {
		pu.SetDeletedAt(*t)
	}
	return pu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (pu *PurchaseUpdate) ClearDeletedAt() *PurchaseUpdate {
	pu.mutation.ClearDeletedAt()
	return pu
}

// SetDeletedBy sets the "deleted_by" field.
func (pu *PurchaseUpdate) SetDeletedBy(u uuid.UUID) *PurchaseUpdate {
	pu.mutation.SetDeletedBy(u)
	return pu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableDeletedBy(u *uuid.UUID) *PurchaseUpdate {
	if u != nil {
		pu.SetDeletedBy(*u)
	}
	return pu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (pu *PurchaseUpdate) ClearDeletedBy() *PurchaseUpdate {
	pu.mutation.ClearDeletedBy()
	return pu
}

// SetIsDeleted sets the "is_deleted" field.
func (pu *PurchaseUpdate) SetIsDeleted(i int64) *PurchaseUpdate {
	pu.mutation.ResetIsDeleted()
	pu.mutation.SetIsDeleted(i)
	return pu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pu *PurchaseUpdate) SetNillableIsDeleted(i *int64) *PurchaseUpdate {
	if i != nil {
		pu.SetIsDeleted(*i)
	}
	return pu
}

// AddIsDeleted adds i to the "is_deleted" field.
func (pu *PurchaseUpdate) AddIsDeleted(i int64) *PurchaseUpdate {
	pu.mutation.AddIsDeleted(i)
	return pu
}

// Mutation returns the PurchaseMutation object of the builder.
func (pu *PurchaseUpdate) Mutation() *PurchaseMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PurchaseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PurchaseUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PurchaseUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PurchaseUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PurchaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(purchase.Table, purchase.Columns, sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ProductID(); ok {
		_spec.SetField(purchase.FieldProductID, field.TypeUUID, value)
	}
	if value, ok := pu.mutation.SupplierID(); ok {
		_spec.SetField(purchase.FieldSupplierID, field.TypeUUID, value)
	}
	if value, ok := pu.mutation.Quantity(); ok {
		_spec.SetField(purchase.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedQuantity(); ok {
		_spec.AddField(purchase.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Date(); ok {
		_spec.SetField(purchase.FieldDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.TotalAmount(); ok {
		_spec.SetField(purchase.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedTotalAmount(); ok {
		_spec.AddField(purchase.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(purchase.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.CreatedBy(); ok {
		_spec.SetField(purchase.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(purchase.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(purchase.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.UpdatedBy(); ok {
		_spec.SetField(purchase.FieldUpdatedBy, field.TypeUUID, value)
	}
	if pu.mutation.UpdatedByCleared() {
		_spec.ClearField(purchase.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.SetField(purchase.FieldDeletedAt, field.TypeTime, value)
	}
	if pu.mutation.DeletedAtCleared() {
		_spec.ClearField(purchase.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.DeletedBy(); ok {
		_spec.SetField(purchase.FieldDeletedBy, field.TypeUUID, value)
	}
	if pu.mutation.DeletedByCleared() {
		_spec.ClearField(purchase.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := pu.mutation.IsDeleted(); ok {
		_spec.SetField(purchase.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedIsDeleted(); ok {
		_spec.AddField(purchase.FieldIsDeleted, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purchase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PurchaseUpdateOne is the builder for updating a single Purchase entity.
type PurchaseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PurchaseMutation
}

// SetProductID sets the "product_id" field.
func (puo *PurchaseUpdateOne) SetProductID(u uuid.UUID) *PurchaseUpdateOne {
	puo.mutation.SetProductID(u)
	return puo
}

// SetSupplierID sets the "supplier_id" field.
func (puo *PurchaseUpdateOne) SetSupplierID(u uuid.UUID) *PurchaseUpdateOne {
	puo.mutation.SetSupplierID(u)
	return puo
}

// SetQuantity sets the "quantity" field.
func (puo *PurchaseUpdateOne) SetQuantity(i int64) *PurchaseUpdateOne {
	puo.mutation.ResetQuantity()
	puo.mutation.SetQuantity(i)
	return puo
}

// AddQuantity adds i to the "quantity" field.
func (puo *PurchaseUpdateOne) AddQuantity(i int64) *PurchaseUpdateOne {
	puo.mutation.AddQuantity(i)
	return puo
}

// SetDate sets the "date" field.
func (puo *PurchaseUpdateOne) SetDate(t time.Time) *PurchaseUpdateOne {
	puo.mutation.SetDate(t)
	return puo
}

// SetTotalAmount sets the "total_amount" field.
func (puo *PurchaseUpdateOne) SetTotalAmount(i int64) *PurchaseUpdateOne {
	puo.mutation.ResetTotalAmount()
	puo.mutation.SetTotalAmount(i)
	return puo
}

// AddTotalAmount adds i to the "total_amount" field.
func (puo *PurchaseUpdateOne) AddTotalAmount(i int64) *PurchaseUpdateOne {
	puo.mutation.AddTotalAmount(i)
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PurchaseUpdateOne) SetCreatedAt(t time.Time) *PurchaseUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableCreatedAt(t *time.Time) *PurchaseUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetCreatedBy sets the "created_by" field.
func (puo *PurchaseUpdateOne) SetCreatedBy(u uuid.UUID) *PurchaseUpdateOne {
	puo.mutation.SetCreatedBy(u)
	return puo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableCreatedBy(u *uuid.UUID) *PurchaseUpdateOne {
	if u != nil {
		puo.SetCreatedBy(*u)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PurchaseUpdateOne) SetUpdatedAt(t time.Time) *PurchaseUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableUpdatedAt(t *time.Time) *PurchaseUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PurchaseUpdateOne) ClearUpdatedAt() *PurchaseUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetUpdatedBy sets the "updated_by" field.
func (puo *PurchaseUpdateOne) SetUpdatedBy(u uuid.UUID) *PurchaseUpdateOne {
	puo.mutation.SetUpdatedBy(u)
	return puo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableUpdatedBy(u *uuid.UUID) *PurchaseUpdateOne {
	if u != nil {
		puo.SetUpdatedBy(*u)
	}
	return puo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (puo *PurchaseUpdateOne) ClearUpdatedBy() *PurchaseUpdateOne {
	puo.mutation.ClearUpdatedBy()
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PurchaseUpdateOne) SetDeletedAt(t time.Time) *PurchaseUpdateOne {
	puo.mutation.SetDeletedAt(t)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableDeletedAt(t *time.Time) *PurchaseUpdateOne {
	if t != nil {
		puo.SetDeletedAt(*t)
	}
	return puo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (puo *PurchaseUpdateOne) ClearDeletedAt() *PurchaseUpdateOne {
	puo.mutation.ClearDeletedAt()
	return puo
}

// SetDeletedBy sets the "deleted_by" field.
func (puo *PurchaseUpdateOne) SetDeletedBy(u uuid.UUID) *PurchaseUpdateOne {
	puo.mutation.SetDeletedBy(u)
	return puo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableDeletedBy(u *uuid.UUID) *PurchaseUpdateOne {
	if u != nil {
		puo.SetDeletedBy(*u)
	}
	return puo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (puo *PurchaseUpdateOne) ClearDeletedBy() *PurchaseUpdateOne {
	puo.mutation.ClearDeletedBy()
	return puo
}

// SetIsDeleted sets the "is_deleted" field.
func (puo *PurchaseUpdateOne) SetIsDeleted(i int64) *PurchaseUpdateOne {
	puo.mutation.ResetIsDeleted()
	puo.mutation.SetIsDeleted(i)
	return puo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (puo *PurchaseUpdateOne) SetNillableIsDeleted(i *int64) *PurchaseUpdateOne {
	if i != nil {
		puo.SetIsDeleted(*i)
	}
	return puo
}

// AddIsDeleted adds i to the "is_deleted" field.
func (puo *PurchaseUpdateOne) AddIsDeleted(i int64) *PurchaseUpdateOne {
	puo.mutation.AddIsDeleted(i)
	return puo
}

// Mutation returns the PurchaseMutation object of the builder.
func (puo *PurchaseUpdateOne) Mutation() *PurchaseMutation {
	return puo.mutation
}

// Where appends a list predicates to the PurchaseUpdate builder.
func (puo *PurchaseUpdateOne) Where(ps ...predicate.Purchase) *PurchaseUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PurchaseUpdateOne) Select(field string, fields ...string) *PurchaseUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Purchase entity.
func (puo *PurchaseUpdateOne) Save(ctx context.Context) (*Purchase, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PurchaseUpdateOne) SaveX(ctx context.Context) *Purchase {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PurchaseUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PurchaseUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PurchaseUpdateOne) sqlSave(ctx context.Context) (_node *Purchase, err error) {
	_spec := sqlgraph.NewUpdateSpec(purchase.Table, purchase.Columns, sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Purchase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, purchase.FieldID)
		for _, f := range fields {
			if !purchase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != purchase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.ProductID(); ok {
		_spec.SetField(purchase.FieldProductID, field.TypeUUID, value)
	}
	if value, ok := puo.mutation.SupplierID(); ok {
		_spec.SetField(purchase.FieldSupplierID, field.TypeUUID, value)
	}
	if value, ok := puo.mutation.Quantity(); ok {
		_spec.SetField(purchase.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedQuantity(); ok {
		_spec.AddField(purchase.FieldQuantity, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Date(); ok {
		_spec.SetField(purchase.FieldDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.TotalAmount(); ok {
		_spec.SetField(purchase.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(purchase.FieldTotalAmount, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(purchase.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.CreatedBy(); ok {
		_spec.SetField(purchase.FieldCreatedBy, field.TypeUUID, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(purchase.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(purchase.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.UpdatedBy(); ok {
		_spec.SetField(purchase.FieldUpdatedBy, field.TypeUUID, value)
	}
	if puo.mutation.UpdatedByCleared() {
		_spec.ClearField(purchase.FieldUpdatedBy, field.TypeUUID)
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.SetField(purchase.FieldDeletedAt, field.TypeTime, value)
	}
	if puo.mutation.DeletedAtCleared() {
		_spec.ClearField(purchase.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.DeletedBy(); ok {
		_spec.SetField(purchase.FieldDeletedBy, field.TypeUUID, value)
	}
	if puo.mutation.DeletedByCleared() {
		_spec.ClearField(purchase.FieldDeletedBy, field.TypeUUID)
	}
	if value, ok := puo.mutation.IsDeleted(); ok {
		_spec.SetField(purchase.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedIsDeleted(); ok {
		_spec.AddField(purchase.FieldIsDeleted, field.TypeInt64, value)
	}
	_node = &Purchase{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{purchase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
