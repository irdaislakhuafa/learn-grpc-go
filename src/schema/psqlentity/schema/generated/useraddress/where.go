// Code generated by ent, DO NOT EDIT.

package useraddress

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUserID, v))
}

// AddressID applies equality check predicate on the "address_id" field. It's identical to AddressIDEQ.
func AddressID(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldAddressID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldDeletedBy, v))
}

// IsDeleted applies equality check predicate on the "is_deleted" field. It's identical to IsDeletedEQ.
func IsDeleted(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldIsDeleted, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldUserID, v))
}

// AddressIDEQ applies the EQ predicate on the "address_id" field.
func AddressIDEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldAddressID, v))
}

// AddressIDNEQ applies the NEQ predicate on the "address_id" field.
func AddressIDNEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldAddressID, v))
}

// AddressIDIn applies the In predicate on the "address_id" field.
func AddressIDIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldAddressID, vs...))
}

// AddressIDNotIn applies the NotIn predicate on the "address_id" field.
func AddressIDNotIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldAddressID, vs...))
}

// AddressIDGT applies the GT predicate on the "address_id" field.
func AddressIDGT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldAddressID, v))
}

// AddressIDGTE applies the GTE predicate on the "address_id" field.
func AddressIDGTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldAddressID, v))
}

// AddressIDLT applies the LT predicate on the "address_id" field.
func AddressIDLT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldAddressID, v))
}

// AddressIDLTE applies the LTE predicate on the "address_id" field.
func AddressIDLTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldAddressID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotNull(FieldUpdatedAt))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotNull(FieldUpdatedBy))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v uuid.UUID) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotNull(FieldDeletedBy))
}

// IsDeletedEQ applies the EQ predicate on the "is_deleted" field.
func IsDeletedEQ(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldEQ(FieldIsDeleted, v))
}

// IsDeletedNEQ applies the NEQ predicate on the "is_deleted" field.
func IsDeletedNEQ(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNEQ(FieldIsDeleted, v))
}

// IsDeletedIn applies the In predicate on the "is_deleted" field.
func IsDeletedIn(vs ...int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldIn(FieldIsDeleted, vs...))
}

// IsDeletedNotIn applies the NotIn predicate on the "is_deleted" field.
func IsDeletedNotIn(vs ...int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldNotIn(FieldIsDeleted, vs...))
}

// IsDeletedGT applies the GT predicate on the "is_deleted" field.
func IsDeletedGT(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGT(FieldIsDeleted, v))
}

// IsDeletedGTE applies the GTE predicate on the "is_deleted" field.
func IsDeletedGTE(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldGTE(FieldIsDeleted, v))
}

// IsDeletedLT applies the LT predicate on the "is_deleted" field.
func IsDeletedLT(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLT(FieldIsDeleted, v))
}

// IsDeletedLTE applies the LTE predicate on the "is_deleted" field.
func IsDeletedLTE(v int64) predicate.UserAddress {
	return predicate.UserAddress(sql.FieldLTE(FieldIsDeleted, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserAddress) predicate.UserAddress {
	return predicate.UserAddress(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserAddress) predicate.UserAddress {
	return predicate.UserAddress(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserAddress) predicate.UserAddress {
	return predicate.UserAddress(sql.NotPredicates(p))
}
