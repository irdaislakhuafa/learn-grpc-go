package purchase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	domPurchase "github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/purchase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	psqlPurchase "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/purchase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/pagination"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Purchase], error)
	Get(ctx context.Context, params parameter.PurchaseGetParam) (*entity.Purchase, error)
	Create(ctx context.Context, params parameter.PurchaseCreateParam) (*entity.Purchase, error)
	Update(ctx context.Context, params parameter.PurchaseUpdateParam) (*entity.Purchase, error)
	Delete(ctx context.Context, params parameter.PurchaseDeleteParam) (*entity.Purchase, error)
}

type purchase struct {
	psql        *generated.Client
	cfg         config.Config
	domPurchase domPurchase.Interface
}

func Init(psql *generated.Client, cfg config.Config, domPurchase domPurchase.Interface) Interface {
	result := purchase{
		psql:        psql,
		cfg:         cfg,
		domPurchase: domPurchase,
	}
	return &result
}

func (self *purchase) GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Purchase], error) {
	params, offset, err := pagination.ParseFromParam(params)
	if err != nil {
		return nil, err
	}

	listPurchase, err := self.psql.Purchase.Query().
		Limit(int(params.Limit)).
		Offset(int(offset)).
		Where(psqlPurchase.IsDeleted(params.IsDeleted)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	totalPurchase, err := self.psql.Purchase.Query().Count(ctx)
	if err != nil {
		return nil, err
	}

	listPurchaseEntity := []entity.Purchase{}
	for _, v := range listPurchase {
		purchase, err := self.domPurchase.ToEntity(*v)
		if err != nil {
			return nil, err
		}
		listPurchaseEntity = append(listPurchaseEntity, purchase)
	}

	totalPages := (totalPurchase / int(params.Limit))
	result := entity.ResponsePagination[entity.Pagination, []entity.Purchase]{
		Pagination: entity.Pagination{
			Total:       uint64(totalPurchase),
			Limit:       params.Limit,
			CurrentPage: params.Page,
			TotalPages:  uint64(totalPages),
		},
		Data: listPurchaseEntity,
	}

	return &result, nil
}

func (self *purchase) Get(ctx context.Context, params parameter.PurchaseGetParam) (*entity.Purchase, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	purchase, err := self.psql.Purchase.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	result, err := self.domPurchase.ToEntity(*purchase)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *purchase) Create(ctx context.Context, params parameter.PurchaseCreateParam) (*entity.Purchase, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	productID, err := uuid.Parse(params.ProductID)
	if err != nil {
		return nil, errors.Join(err, errors.New("product_id is not uuid"))
	}

	supplierID, err := uuid.Parse(params.SupplierID)
	if err != nil {
		return nil, errors.Join(err, errors.New("supplier_id is not uuid"))
	}

	purchase, err := tx.Purchase.Create().
		SetProductID(productID).
		SetSupplierID(supplierID).
		SetQuantity(params.Quantity).
		SetTotalAmount(params.TotalAmount).
		SetDate(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domPurchase.ToEntity(*purchase)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *purchase) Update(ctx context.Context, params parameter.PurchaseUpdateParam) (*entity.Purchase, error) {
	panic("not implemented") // TODO: Implement
}

func (self *purchase) Delete(ctx context.Context, params parameter.PurchaseDeleteParam) (*entity.Purchase, error) {
	panic("not implemented") // TODO: Implement
}
