package purchase

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	ToEntity(params generated.Purchase) (entity.Purchase, error)
}

type purchase struct {
	psql *generated.Client
	cfg  config.Config
}

func Init(psql *generated.Client, cfg config.Config) Interface {
	result := purchase{
		psql: psql,
		cfg:  cfg,
	}
	return &result
}

func (self *purchase) ToEntity(params generated.Purchase) (entity.Purchase, error) {
	result := entity.Purchase{
		ID:          params.ID,
		ProductID:   params.ProductID,
		SupplierID:  params.SupplierID,
		Quantity:    params.Quantity,
		Date:        params.Date,
		TotalAmount: params.TotalAmount,
		CreatedAt:   params.CreatedAt,
		CreatedBy:   params.CreatedBy,
		UpdatedAt:   params.UpdatedAt,
		UpdatedBy:   params.UpdatedBy,
		DeletedAt:   params.CreatedAt,
		DeletedBy:   params.DeletedBy,
		IsDeleted:   params.IsDeleted,
	}

	return result, nil
}
