// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/address"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// Regency holds the value of the "regency" field.
	Regency string `json:"regency,omitempty"`
	// SubDistrict holds the value of the "sub_district" field.
	SubDistrict string `json:"sub_district,omitempty"`
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
	IsDeleted    int `json:"is_deleted,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldIsDeleted:
			values[i] = new(sql.NullInt64)
		case address.FieldCountry, address.FieldProvince, address.FieldRegency, address.FieldSubDistrict:
			values[i] = new(sql.NullString)
		case address.FieldCreatedAt, address.FieldUpdatedAt, address.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case address.FieldID, address.FieldCreatedBy, address.FieldUpdatedBy, address.FieldDeletedBy:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case address.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				a.Country = value.String
			}
		case address.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				a.Province = value.String
			}
		case address.FieldRegency:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field regency", values[i])
			} else if value.Valid {
				a.Regency = value.String
			}
		case address.FieldSubDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sub_district", values[i])
			} else if value.Valid {
				a.SubDistrict = value.String
			}
		case address.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case address.FieldCreatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value != nil {
				a.CreatedBy = *value
			}
		case address.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case address.FieldUpdatedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value != nil {
				a.UpdatedBy = *value
			}
		case address.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				a.DeletedAt = value.Time
			}
		case address.FieldDeletedBy:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value != nil {
				a.DeletedBy = *value
			}
		case address.FieldIsDeleted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_deleted", values[i])
			} else if value.Valid {
				a.IsDeleted = int(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Address.
// This includes values selected through modifiers, order, etc.
func (a *Address) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return NewAddressClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("generated: Address is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("country=")
	builder.WriteString(a.Country)
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(a.Province)
	builder.WriteString(", ")
	builder.WriteString("regency=")
	builder.WriteString(a.Regency)
	builder.WriteString(", ")
	builder.WriteString("sub_district=")
	builder.WriteString(a.SubDistrict)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(a.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", a.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("is_deleted=")
	builder.WriteString(fmt.Sprintf("%v", a.IsDeleted))
	builder.WriteByte(')')
	return builder.String()
}

// Addresses is a parsable slice of Address.
type Addresses []*Address
