package user

import (
	"context"

	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/params"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	psqlUser "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/responses"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params params.UserPaginationParam) (*responses.ResponsePagination[responses.Pagination, []*generated.User], error)
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

func (self *user) GetListWithPagination(ctx context.Context, params params.UserPaginationParam) (*responses.ResponsePagination[responses.Pagination, []*generated.User], error) {
	// handle pagination offset
	offset := uint64(0)
	if params.Page <= 1 {
		offset = 0
	} else {
		offset = (params.Page - 1) * params.Limit
	}

	listUser, err := self.psql.User.Query().
		Limit(int(params.Limit)).
		Where(psqlUser.IsDeleted(params.IsDeleted)).
		Offset(int(offset)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	totalUser, err := self.psql.User.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	result := &responses.ResponsePagination[responses.Pagination, []*generated.User]{
		Pagination: responses.Pagination{
			Total:       uint64(totalUser),
			Limit:       params.Limit,
			CurrentPage: params.Page,
			TotalPages:  uint64(totalUser / int(params.Limit)),
		},
		Data: listUser,
	}

	return result, nil
}
