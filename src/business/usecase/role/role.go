package role

import (
	"context"

	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params parameter.RolePaginationParam) (entity.ResponsePagination[entity.Pagination, []entity.Role], error)
	Get()
}
