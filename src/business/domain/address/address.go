package address

import (
	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	ToEntity(params generated.Address, userID uuid.UUID) (entity.Address, error)
}

type address struct {
	psql *generated.Client
	cfg  config.Config
}

func Init(psql *generated.Client, cfg config.Config) Interface {
	result := address{
		psql: psql,
		cfg:  cfg,
	}
	return &result
}

func (self *address) ToEntity(params generated.Address, userID uuid.UUID) (entity.Address, error) {
	result := entity.Address{
		ID:          params.ID,
		Country:     params.Country,
		Province:    params.Province,
		Regency:     params.Regency,
		SubDistrict: params.SubDistrict,
		UserID:      userID,
		CreatedAt:   params.CreatedAt,
		CreatedBy:   params.CreatedBy,
		UpdatedAt:   params.UpdatedAt,
		UpdatedBy:   params.UpdatedBy,
		DeletedAt:   params.DeletedAt,
		DeletedBy:   params.DeletedBy,
		IsDeleted:   params.IsDeleted,
	}

	return result, nil
}
