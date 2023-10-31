package sale

import (
	"context"

	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Sale], error)
	Get(ctx context.Context, params parameter.SaleGetParam) (*entity.Sale, error)
	Create(ctx context.Context, params parameter.SaleCreateParam) (*entity.Sale, error)
	Update(ctx context.Context, params parameter.SaleUpdateParam) (*entity.Sale, error)
	Delete(ctx context.Context, params parameter.SaleDeleteParam) (*entity.Sale, error)
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

func (self *sale) GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Sale], error) {
	panic("not implemented") // TODO: Implement
}

func (self *sale) Get(ctx context.Context, params parameter.SaleGetParam) (*entity.Sale, error) {
	panic("not implemented") // TODO: Implement
}

func (self *sale) Create(ctx context.Context, params parameter.SaleCreateParam) (*entity.Sale, error) {
	panic("not implemented") // TODO: Implement
}

func (self *sale) Update(ctx context.Context, params parameter.SaleUpdateParam) (*entity.Sale, error) {
	panic("not implemented") // TODO: Implement
}

func (self *sale) Delete(ctx context.Context, params parameter.SaleDeleteParam) (*entity.Sale, error) {
	panic("not implemented") // TODO: Implement
}
