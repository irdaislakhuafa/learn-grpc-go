// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/purchase"
)

// Purchase is the model entity for the Purchase schema.
type Purchase struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID uuid.UUID `json:"product_id,omitempty"`
	// SupplierID holds the value of the "supplier_id" field.
	SupplierID uuid.UUID `json:"supplier_id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int64 `json:"quantity,omitempty"`
	// Date holds the value of the "date" field.
	Date time.Time `json:"date,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount int64 `json:"total_amount,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy uuid.UUID `json:"created_by,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy uuid.UUID `json:"updated_by,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy uuid.UUID `json:"deleted_by,omitempty"`
	// IsDeleted holds the value of the "is_deleted" field.
	IsDeleted    int64 `json:"is_deleted,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Purchase) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case purchase.FieldQuantity, purchase.FieldTotalAmount, purchase.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case purchase.FieldDate, purchase.FieldCreatedAt, purchase.FieldUpdatedAt, purchase.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case purchase.FieldID, purchase.FieldProductID, purchase.FieldSupplierID, purchase.FieldCreatedBy, purchase.FieldUpdatedBy, purchase.FieldDeletedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Purchase fields.
func (pu *Purchase) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case purchase.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pu.ID = *value
			}
		case purchase.FieldProductID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value != nil {
				pu.ProductID = *value
			}
		case purchase.FieldSupplierID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field supplier_id", values[i])
			} else if value != nil {
				pu.SupplierID = *value
			}
		case purchase.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				pu.Quantity = value.Int64
			}
		case purchase.FieldDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				pu.Date = value.Time
			}
		case purchase.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				pu.TotalAmount = value.Int64
			}
		case purchase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pu.CreatedAt = value.Time
			}
		case purchase.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				pu.CreatedBy = *value
			}
		case purchase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pu.UpdatedAt = value.Time
			}
		case purchase.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				pu.UpdatedBy = *value
			}
		case purchase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pu.DeletedAt = value.Time
			}
		case purchase.FieldDeletedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value != nil {
				pu.DeletedBy = *value
			}
		case purchase.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				pu.IsDeleted = value.Int64
			}
		default:
			pu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Purchase.
// This includes values selected through modifiers, order, etc.
func (pu *Purchase) Value(name string) (ent.Value, error) {
	return pu.selectValues.Get(name)
}

// Update returns a builder for updating this Purchase.
// Note that you need to call Purchase.Unwrap() before calling this method if this Purchase
// was returned from a transaction, and the transaction was committed or rolled back.
func (pu *Purchase) Update() *PurchaseUpdateOne {
	return NewPurchaseClient(pu.config).UpdateOne(pu)
}

// Unwrap unwraps the Purchase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pu *Purchase) Unwrap() *Purchase {
	_tx, ok := pu.config.driver.(*txDriver)
	if !ok {
		panic("generated: Purchase is not a transactional entity")
	}
	pu.config.driver = _tx.drv
	return pu
}

// String implements the fmt.Stringer.
func (pu *Purchase) String() string {
	var builder strings.Builder
	builder.WriteString("Purchase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pu.ID))
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", pu.ProductID))
	builder.WriteString(", ")
	builder.WriteString("supplier_id=")
	builder.WriteString(fmt.Sprintf("%v", pu.SupplierID))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", pu.Quantity))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(pu.Date.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(fmt.Sprintf("%v", pu.TotalAmount))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pu.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pu.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(pu.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", pu.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", pu.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// Purchases is a parsable slice of Purchase.
type Purchases []*Purchase