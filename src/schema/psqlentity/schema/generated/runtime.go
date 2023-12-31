// Code generated by ent, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/address"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/purchase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/role"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/sale"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCountry is the schema descriptor for country field.
	addressDescCountry := addressFields[0].Descriptor()
	// address.CountryValidator is a validator for the "country" field. It is called by the builders before save.
	address.CountryValidator = addressDescCountry.Validators[0].(func(string) error)
	// addressDescProvince is the schema descriptor for province field.
	addressDescProvince := addressFields[1].Descriptor()
	// address.ProvinceValidator is a validator for the "province" field. It is called by the builders before save.
	address.ProvinceValidator = addressDescProvince.Validators[0].(func(string) error)
	// addressDescRegency is the schema descriptor for regency field.
	addressDescRegency := addressFields[2].Descriptor()
	// address.RegencyValidator is a validator for the "regency" field. It is called by the builders before save.
	address.RegencyValidator = addressDescRegency.Validators[0].(func(string) error)
	// addressDescSubDistrict is the schema descriptor for sub_district field.
	addressDescSubDistrict := addressFields[3].Descriptor()
	// address.SubDistrictValidator is a validator for the "sub_district" field. It is called by the builders before save.
	address.SubDistrictValidator = addressDescSubDistrict.Validators[0].(func(string) error)
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressFields[6].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescCreatedBy is the schema descriptor for created_by field.
	addressDescCreatedBy := addressFields[7].Descriptor()
	// address.DefaultCreatedBy holds the default value on creation for the created_by field.
	address.DefaultCreatedBy = addressDescCreatedBy.Default.(func() uuid.UUID)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressFields[8].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// addressDescUpdatedBy is the schema descriptor for updated_by field.
	addressDescUpdatedBy := addressFields[9].Descriptor()
	// address.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	address.DefaultUpdatedBy = addressDescUpdatedBy.Default.(func() uuid.UUID)
	// addressDescDeletedAt is the schema descriptor for deleted_at field.
	addressDescDeletedAt := addressFields[10].Descriptor()
	// address.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	address.DefaultDeletedAt = addressDescDeletedAt.Default.(func() time.Time)
	// addressDescDeletedBy is the schema descriptor for deleted_by field.
	addressDescDeletedBy := addressFields[11].Descriptor()
	// address.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	address.DefaultDeletedBy = addressDescDeletedBy.Default.(func() uuid.UUID)
	// addressDescIsDeleted is the schema descriptor for is_deleted field.
	addressDescIsDeleted := addressFields[12].Descriptor()
	// address.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	address.DefaultIsDeleted = addressDescIsDeleted.Default.(int64)
	// addressDescID is the schema descriptor for id field.
	addressDescID := addressFields[5].Descriptor()
	// address.DefaultID holds the default value on creation for the id field.
	address.DefaultID = addressDescID.Default.(func() uuid.UUID)
	purchaseFields := schema.Purchase{}.Fields()
	_ = purchaseFields
	// purchaseDescCreatedAt is the schema descriptor for created_at field.
	purchaseDescCreatedAt := purchaseFields[6].Descriptor()
	// purchase.DefaultCreatedAt holds the default value on creation for the created_at field.
	purchase.DefaultCreatedAt = purchaseDescCreatedAt.Default.(func() time.Time)
	// purchaseDescCreatedBy is the schema descriptor for created_by field.
	purchaseDescCreatedBy := purchaseFields[7].Descriptor()
	// purchase.DefaultCreatedBy holds the default value on creation for the created_by field.
	purchase.DefaultCreatedBy = purchaseDescCreatedBy.Default.(func() uuid.UUID)
	// purchaseDescUpdatedAt is the schema descriptor for updated_at field.
	purchaseDescUpdatedAt := purchaseFields[8].Descriptor()
	// purchase.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	purchase.DefaultUpdatedAt = purchaseDescUpdatedAt.Default.(func() time.Time)
	// purchaseDescUpdatedBy is the schema descriptor for updated_by field.
	purchaseDescUpdatedBy := purchaseFields[9].Descriptor()
	// purchase.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	purchase.DefaultUpdatedBy = purchaseDescUpdatedBy.Default.(func() uuid.UUID)
	// purchaseDescDeletedAt is the schema descriptor for deleted_at field.
	purchaseDescDeletedAt := purchaseFields[10].Descriptor()
	// purchase.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	purchase.DefaultDeletedAt = purchaseDescDeletedAt.Default.(func() time.Time)
	// purchaseDescDeletedBy is the schema descriptor for deleted_by field.
	purchaseDescDeletedBy := purchaseFields[11].Descriptor()
	// purchase.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	purchase.DefaultDeletedBy = purchaseDescDeletedBy.Default.(func() uuid.UUID)
	// purchaseDescIsDeleted is the schema descriptor for is_deleted field.
	purchaseDescIsDeleted := purchaseFields[12].Descriptor()
	// purchase.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	purchase.DefaultIsDeleted = purchaseDescIsDeleted.Default.(int64)
	// purchaseDescID is the schema descriptor for id field.
	purchaseDescID := purchaseFields[5].Descriptor()
	// purchase.DefaultID holds the default value on creation for the id field.
	purchase.DefaultID = purchaseDescID.Default.(func() uuid.UUID)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescCreatedAt is the schema descriptor for created_at field.
	roleDescCreatedAt := roleFields[3].Descriptor()
	// role.DefaultCreatedAt holds the default value on creation for the created_at field.
	role.DefaultCreatedAt = roleDescCreatedAt.Default.(func() time.Time)
	// roleDescCreatedBy is the schema descriptor for created_by field.
	roleDescCreatedBy := roleFields[4].Descriptor()
	// role.DefaultCreatedBy holds the default value on creation for the created_by field.
	role.DefaultCreatedBy = roleDescCreatedBy.Default.(func() uuid.UUID)
	// roleDescUpdatedAt is the schema descriptor for updated_at field.
	roleDescUpdatedAt := roleFields[5].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// roleDescUpdatedBy is the schema descriptor for updated_by field.
	roleDescUpdatedBy := roleFields[6].Descriptor()
	// role.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	role.DefaultUpdatedBy = roleDescUpdatedBy.Default.(func() uuid.UUID)
	// roleDescDeletedAt is the schema descriptor for deleted_at field.
	roleDescDeletedAt := roleFields[7].Descriptor()
	// role.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	role.DefaultDeletedAt = roleDescDeletedAt.Default.(func() time.Time)
	// roleDescDeletedBy is the schema descriptor for deleted_by field.
	roleDescDeletedBy := roleFields[8].Descriptor()
	// role.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	role.DefaultDeletedBy = roleDescDeletedBy.Default.(func() uuid.UUID)
	// roleDescIsDeleted is the schema descriptor for is_deleted field.
	roleDescIsDeleted := roleFields[9].Descriptor()
	// role.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	role.DefaultIsDeleted = roleDescIsDeleted.Default.(int64)
	// roleDescID is the schema descriptor for id field.
	roleDescID := roleFields[2].Descriptor()
	// role.DefaultID holds the default value on creation for the id field.
	role.DefaultID = roleDescID.Default.(func() uuid.UUID)
	saleFields := schema.Sale{}.Fields()
	_ = saleFields
	// saleDescCreatedAt is the schema descriptor for created_at field.
	saleDescCreatedAt := saleFields[5].Descriptor()
	// sale.DefaultCreatedAt holds the default value on creation for the created_at field.
	sale.DefaultCreatedAt = saleDescCreatedAt.Default.(func() time.Time)
	// saleDescCreatedBy is the schema descriptor for created_by field.
	saleDescCreatedBy := saleFields[6].Descriptor()
	// sale.DefaultCreatedBy holds the default value on creation for the created_by field.
	sale.DefaultCreatedBy = saleDescCreatedBy.Default.(func() uuid.UUID)
	// saleDescUpdatedAt is the schema descriptor for updated_at field.
	saleDescUpdatedAt := saleFields[7].Descriptor()
	// sale.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	sale.DefaultUpdatedAt = saleDescUpdatedAt.Default.(func() time.Time)
	// saleDescUpdatedBy is the schema descriptor for updated_by field.
	saleDescUpdatedBy := saleFields[8].Descriptor()
	// sale.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	sale.DefaultUpdatedBy = saleDescUpdatedBy.Default.(func() uuid.UUID)
	// saleDescDeletedAt is the schema descriptor for deleted_at field.
	saleDescDeletedAt := saleFields[9].Descriptor()
	// sale.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	sale.DefaultDeletedAt = saleDescDeletedAt.Default.(func() time.Time)
	// saleDescDeletedBy is the schema descriptor for deleted_by field.
	saleDescDeletedBy := saleFields[10].Descriptor()
	// sale.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	sale.DefaultDeletedBy = saleDescDeletedBy.Default.(func() uuid.UUID)
	// saleDescIsDeleted is the schema descriptor for is_deleted field.
	saleDescIsDeleted := saleFields[11].Descriptor()
	// sale.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	sale.DefaultIsDeleted = saleDescIsDeleted.Default.(int64)
	// saleDescID is the schema descriptor for id field.
	saleDescID := saleFields[4].Descriptor()
	// sale.DefaultID holds the default value on creation for the id field.
	sale.DefaultID = saleDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[2].Descriptor()
	// user.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	user.AgeValidator = userDescAge.Validators[0].(func(int) error)
	// userDescHobbies is the schema descriptor for hobbies field.
	userDescHobbies := userFields[3].Descriptor()
	// user.DefaultHobbies holds the default value on creation for the hobbies field.
	user.DefaultHobbies = userDescHobbies.Default.([]string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userFields[6].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(func() uuid.UUID)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userFields[8].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(func() uuid.UUID)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userFields[9].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(func() time.Time)
	// userDescDeletedBy is the schema descriptor for deleted_by field.
	userDescDeletedBy := userFields[10].Descriptor()
	// user.DefaultDeletedBy holds the default value on creation for the deleted_by field.
	user.DefaultDeletedBy = userDescDeletedBy.Default.(func() uuid.UUID)
	// userDescIsDeleted is the schema descriptor for is_deleted field.
	userDescIsDeleted := userFields[11].Descriptor()
	// user.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	user.DefaultIsDeleted = userDescIsDeleted.Default.(int64)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[4].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
