package user

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	ToEntity(params generated.User) (entity.User, error)
}

type user struct {
	psql *generated.Client
	cfg  config.Config
}

func Init(psql *generated.Client, cfg config.Config) Interface {
	result := user{
		psql: psql,
		cfg:  cfg,
	}
	return &result
}

func (self *user) ToEntity(params generated.User) (entity.User, error) {
	result := entity.User{
		ID:        params.ID,
		Name:      params.Name,
		Email:     params.Email,
		Age:       params.Age,
		Hobbies:   params.Hobbies,
		CreatedAt: params.CreatedAt,
		CreatedBy: params.CreatedBy,
		UpdatedAt: params.CreatedAt,
		UpdatedBy: params.UpdatedBy,
		DeletedAt: params.DeletedAt,
		DeletedBy: params.DeletedBy,
		Address:   entity.Address{},
		IsDeleted: params.IsDeleted,
	}
	return result, nil
}
