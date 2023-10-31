package role

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	ToEntity(params generated.Role) (entity.Role, error)
}

type role struct {
	psql *generated.Client
	cfg  config.Config
}

func Init(psql *generated.Client, cfg config.Config) Interface {
	result := role{
		psql: psql,
		cfg:  cfg,
	}
	return &result
}

func (self *role) ToEntity(params generated.Role) (entity.Role, error) {
	result := entity.Role{
		ID:          params.ID,
		Name:        params.Name,
		Description: params.Description,
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
