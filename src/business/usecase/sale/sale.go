package sale

import (
	"context"

	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	psqlSale "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/sale"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/operator"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/pagination"
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
	params, offset, err := pagination.ParseFromParam(params)
	if err != nil {
		return nil, err
	}

	listSale, err := self.psql.Sale.Query().
		Limit(int(params.Limit)).
		Offset(int(offset)).
		Where(psqlSale.IsDeleted(params.IsDeleted)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	totalSale, err := self.psql.Sale.Query().Where(psqlSale.IsDeleted(params.IsDeleted)).Count(ctx)
	if err != nil {
		return nil, err
	}

	listSaleEntity := []entity.Sale{}
	for _, v := range listSale {
		sale := entity.Sale{
			ID:          v.ID,
			ProductID:   v.ProductID,
			Quantity:    v.Quantity,
			TotalAmount: v.TotalAmount,
			Date:        v.Date,
			CreatedAt:   v.CreatedAt,
			CreatedBy:   v.CreatedBy,
			UpdatedAt:   v.UpdatedAt,
			UpdatedBy:   v.UpdatedBy,
			DeletedAt:   v.DeletedAt,
			DeletedBy:   v.DeletedBy,
			IsDeleted:   v.IsDeleted,
		}
		listSaleEntity = append(listSaleEntity, sale)
	}

	totalPages := uint64(totalSale / int(params.Limit))
	result := entity.ResponsePagination[entity.Pagination, []entity.Sale]{
		Pagination: entity.Pagination{
			Total:       uint64(totalSale),
			Limit:       params.Limit,
			CurrentPage: params.Page,
			TotalPages:  operator.Ternary[uint64](totalPages <= 0, 1, totalPages),
		},
		Data: listSaleEntity,
	}

	return &result, nil
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
