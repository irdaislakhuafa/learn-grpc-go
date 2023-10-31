// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/purchase"
)

// PurchaseCreate is the builder for creating a Purchase entity.
type PurchaseCreate struct {
	config
	mutation *PurchaseMutation
	hooks    []Hook
}

// SetProductID sets the "product_id" field.
func (pc *PurchaseCreate) SetProductID(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetProductID(u)
	return pc
}

// SetSupplierID sets the "supplier_id" field.
func (pc *PurchaseCreate) SetSupplierID(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetSupplierID(u)
	return pc
}

// SetQuantity sets the "quantity" field.
func (pc *PurchaseCreate) SetQuantity(i int64) *PurchaseCreate {
	pc.mutation.SetQuantity(i)
	return pc
}

// SetDate sets the "date" field.
func (pc *PurchaseCreate) SetDate(t time.Time) *PurchaseCreate {
	pc.mutation.SetDate(t)
	return pc
}

// SetTotalAmount sets the "total_amount" field.
func (pc *PurchaseCreate) SetTotalAmount(i int64) *PurchaseCreate {
	pc.mutation.SetTotalAmount(i)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PurchaseCreate) SetCreatedAt(t time.Time) *PurchaseCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableCreatedAt(t *time.Time) *PurchaseCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetCreatedBy sets the "created_by" field.
func (pc *PurchaseCreate) SetCreatedBy(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetCreatedBy(u)
	return pc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableCreatedBy(u *uuid.UUID) *PurchaseCreate {
	if u != nil {
		pc.SetCreatedBy(*u)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PurchaseCreate) SetUpdatedAt(t time.Time) *PurchaseCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableUpdatedAt(t *time.Time) *PurchaseCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetUpdatedBy sets the "updated_by" field.
func (pc *PurchaseCreate) SetUpdatedBy(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetUpdatedBy(u)
	return pc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableUpdatedBy(u *uuid.UUID) *PurchaseCreate {
	if u != nil {
		pc.SetUpdatedBy(*u)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PurchaseCreate) SetDeletedAt(t time.Time) *PurchaseCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableDeletedAt(t *time.Time) *PurchaseCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetDeletedBy sets the "deleted_by" field.
func (pc *PurchaseCreate) SetDeletedBy(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetDeletedBy(u)
	return pc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableDeletedBy(u *uuid.UUID) *PurchaseCreate {
	if u != nil {
		pc.SetDeletedBy(*u)
	}
	return pc
}

// SetIsDeleted sets the "is_deleted" field.
func (pc *PurchaseCreate) SetIsDeleted(i int64) *PurchaseCreate {
	pc.mutation.SetIsDeleted(i)
	return pc
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableIsDeleted(i *int64) *PurchaseCreate {
	if i != nil {
		pc.SetIsDeleted(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PurchaseCreate) SetID(u uuid.UUID) *PurchaseCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PurchaseCreate) SetNillableID(u *uuid.UUID) *PurchaseCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PurchaseMutation object of the builder.
func (pc *PurchaseCreate) Mutation() *PurchaseMutation {
	return pc.mutation
}

// Save creates the Purchase in the database.
func (pc *PurchaseCreate) Save(ctx context.Context) (*Purchase, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PurchaseCreate) SaveX(ctx context.Context) *Purchase {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PurchaseCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PurchaseCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PurchaseCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := purchase.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		v := purchase.DefaultCreatedBy()
		pc.mutation.SetCreatedBy(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := purchase.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedBy(); !ok {
		v := purchase.DefaultUpdatedBy()
		pc.mutation.SetUpdatedBy(v)
	}
	if _, ok := pc.mutation.DeletedAt(); !ok {
		v := purchase.DefaultDeletedAt()
		pc.mutation.SetDeletedAt(v)
	}
	if _, ok := pc.mutation.DeletedBy(); !ok {
		v := purchase.DefaultDeletedBy()
		pc.mutation.SetDeletedBy(v)
	}
	if _, ok := pc.mutation.IsDeleted(); !ok {
		v := purchase.DefaultIsDeleted
		pc.mutation.SetIsDeleted(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := purchase.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PurchaseCreate) check() error {
	if _, ok := pc.mutation.ProductID(); !ok {
		return &ValidationError{Name: "product_id", err: errors.New(`generated: missing required field "Purchase.product_id"`)}
	}
	if _, ok := pc.mutation.SupplierID(); !ok {
		return &ValidationError{Name: "supplier_id", err: errors.New(`generated: missing required field "Purchase.supplier_id"`)}
	}
	if _, ok := pc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`generated: missing required field "Purchase.quantity"`)}
	}
	if _, ok := pc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`generated: missing required field "Purchase.date"`)}
	}
	if _, ok := pc.mutation.TotalAmount(); !ok {
		return &ValidationError{Name: "total_amount", err: errors.New(`generated: missing required field "Purchase.total_amount"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Purchase.created_at"`)}
	}
	if _, ok := pc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`generated: missing required field "Purchase.created_by"`)}
	}
	if _, ok := pc.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`generated: missing required field "Purchase.is_deleted"`)}
	}
	return nil
}

func (pc *PurchaseCreate) sqlSave(ctx context.Context) (*Purchase, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PurchaseCreate) createSpec() (*Purchase, *sqlgraph.CreateSpec) {
	var (
		_node = &Purchase{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(purchase.Table, sqlgraph.NewFieldSpec(purchase.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.ProductID(); ok {
		_spec.SetField(purchase.FieldProductID, field.TypeUUID, value)
		_node.ProductID = value
	}
	if value, ok := pc.mutation.SupplierID(); ok {
		_spec.SetField(purchase.FieldSupplierID, field.TypeUUID, value)
		_node.SupplierID = value
	}
	if value, ok := pc.mutation.Quantity(); ok {
		_spec.SetField(purchase.FieldQuantity, field.TypeInt64, value)
		_node.Quantity = value
	}
	if value, ok := pc.mutation.Date(); ok {
		_spec.SetField(purchase.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := pc.mutation.TotalAmount(); ok {
		_spec.SetField(purchase.FieldTotalAmount, field.TypeInt64, value)
		_node.TotalAmount = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(purchase.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.CreatedBy(); ok {
		_spec.SetField(purchase.FieldCreatedBy, field.TypeUUID, value)
		_node.CreatedBy = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(purchase.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.UpdatedBy(); ok {
		_spec.SetField(purchase.FieldUpdatedBy, field.TypeUUID, value)
		_node.UpdatedBy = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(purchase.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pc.mutation.DeletedBy(); ok {
		_spec.SetField(purchase.FieldDeletedBy, field.TypeUUID, value)
		_node.DeletedBy = value
	}
	if value, ok := pc.mutation.IsDeleted(); ok {
		_spec.SetField(purchase.FieldIsDeleted, field.TypeInt64, value)
		_node.IsDeleted = value
	}
	return _node, _spec
}

// PurchaseCreateBulk is the builder for creating many Purchase entities in bulk.
type PurchaseCreateBulk struct {
	config
	err      error
	builders []*PurchaseCreate
}

// Save creates the Purchase entities in the database.
func (pcb *PurchaseCreateBulk) Save(ctx context.Context) ([]*Purchase, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Purchase, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PurchaseMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PurchaseCreateBulk) SaveX(ctx context.Context) []*Purchase {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PurchaseCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PurchaseCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
