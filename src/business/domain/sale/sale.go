package sale

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	ToEntity(params generated.Sale) (entity.Sale, error)
}

type sale struct {
	psql *generated.Client
	cfg  config.Config
}

func Init(psql *generated.Client, cfg config.Config) Interface {
	result := sale{
		psql: psql,
		cfg:  cfg,
	}
	return &result
}

func (self *sale) ToEntity(params generated.Sale) (entity.Sale, error) {
	result := entity.Sale{
		ID:          params.ID,
		ProductID:   params.ProductID,
		Quantity:    params.Quantity,
		TotalAmount: params.TotalAmount,
		Date:        params.Date,
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
