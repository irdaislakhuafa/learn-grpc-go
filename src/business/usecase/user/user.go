package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/params"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/address"
	psqlUser "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/operator"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params params.UserPaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.User], error)
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

func (self *user) GetListWithPagination(ctx context.Context, params params.UserPaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.User], error) {
	// handle pagination offset
	offset := uint64(0)
	if params.Page <= 1 {
		offset = 0
	} else {
		offset = (params.Page - 1) * params.Limit
	}

	// handle zero divider
	if params.Limit <= 0 {
		return nil, errors.New("minimum limit is 1")
	}

	// get list user from db
	listUser, err := self.psql.User.Query().
		Limit(int(params.Limit)).
		Where(psqlUser.IsDeleted(params.IsDeleted)).
		Offset(int(offset)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	// find list address for list user
	listUserID := []uuid.UUID{}
	for _, v := range listUser {
		listUserID = append(listUserID, v.ID)
	}

	listAddress, err := self.psql.Address.Query().Where(address.UserIDIn(listUserID...)).All(ctx)
	if err != nil {
		return nil, err
	}

	mapListAddressByUserID := map[string]generated.Address{}
	for _, v := range listAddress {
		mapListAddressByUserID[v.UserID.String()] = *v
	}

	// mapping list user from sql to entity
	listUserEntity := []entity.User{}
	for _, v := range listUser {
		address := mapListAddressByUserID[v.ID.String()]
		user := entity.User{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Age:       v.Age,
			Hobbies:   v.Hobbies,
			CreatedAt: v.CreatedAt,
			CreatedBy: v.CreatedBy,
			UpdatedAt: v.CreatedAt,
			UpdatedBy: v.UpdatedBy,
			DeletedAt: v.DeletedAt,
			DeletedBy: v.DeletedBy,
			Address: entity.Address{
				ID:          address.ID,
				Country:     address.Country,
				Province:    address.Province,
				Regency:     address.Regency,
				SubDistrict: address.SubDistrict,
				UserID:      v.ID,
				CreatedAt:   address.CreatedAt,
				CreatedBy:   address.CreatedBy,
				UpdatedAt:   address.UpdatedAt,
				UpdatedBy:   address.UpdatedBy,
				DeletedAt:   address.DeletedAt,
				DeletedBy:   address.DeletedBy,
				IsDeleted:   address.IsDeleted,
			},
			IsDeleted: v.IsDeleted,
		}

		listUserEntity = append(listUserEntity, user)
	}

	// count total data all user
	totalUser, err := self.psql.User.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := uint64(totalUser / int(params.Limit))
	result := &entity.ResponsePagination[entity.Pagination, []entity.User]{
		Pagination: entity.Pagination{
			Total:       uint64(totalUser),
			Limit:       params.Limit,
			CurrentPage: params.Page,
			TotalPages:  uint64(operator.Ternary(totalPages <= 0, 1, int(totalPages))),
		},
		Data: listUserEntity,
	}

	return result, nil
}
