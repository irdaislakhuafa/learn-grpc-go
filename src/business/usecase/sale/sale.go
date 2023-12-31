package sale

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	domSale "github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/sale"
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
	psql    *generated.Client
	cfg     config.Config
	domSale domSale.Interface
}

func Init(psql *generated.Client, cfg config.Config, domSale domSale.Interface) Interface {
	result := sale{
		psql:    psql,
		cfg:     cfg,
		domSale: domSale,
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
		sale, err := self.domSale.ToEntity(*v)
		if err != nil {
			return nil, err
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
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not uuid"))
	}

	sale, err := self.psql.Sale.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	result, err := self.domSale.ToEntity(*sale)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *sale) Create(ctx context.Context, params parameter.SaleCreateParam) (*entity.Sale, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	productID, err := uuid.Parse(params.ProductID)
	if err != nil {
		return nil, errors.Join(err, errors.New("product_id is not uuid"))
	}

	sale, err := tx.Sale.Create().
		SetProductID(productID).
		SetQuantity(params.Quantity).
		SetTotalAmount(params.TotalAmount).
		SetDate(time.Now()).

		// TODO: set authentication for GRPC
		SetCreatedBy(uuid.New()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domSale.ToEntity(*sale)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *sale) Update(ctx context.Context, params parameter.SaleUpdateParam) (*entity.Sale, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not uuid"))
	}

	productID, err := uuid.Parse(params.ProductID)
	if err != nil {
		return nil, errors.Join(err, errors.New("product_id is not uuid"))
	}

	err = tx.Sale.UpdateOneID(id).
		SetProductID(productID).
		SetQuantity(params.Quantity).
		SetTotalAmount(params.TotalAmount).
		SetDate(params.Date).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	sale, err := tx.Sale.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domSale.ToEntity(*sale)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *sale) Delete(ctx context.Context, params parameter.SaleDeleteParam) (*entity.Sale, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	if err := tx.Sale.UpdateOneID(id).SetIsDeleted(1).Exec(ctx); err != nil {
		return nil, err
	}

	sale, err := tx.Sale.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domSale.ToEntity(*sale)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
