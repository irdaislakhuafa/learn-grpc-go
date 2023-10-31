package role

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	domRole "github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/role"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	psqlRole "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/role"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/operator"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/pagination"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Role], error)
	Get(ctx context.Context, params parameter.RoleGetParam) (*entity.Role, error)
	Create(ctx context.Context, params parameter.RoleCreateParam) (*entity.Role, error)
	Update(ctx context.Context, params parameter.RoleUpdateParam) (*entity.Role, error)
	Delete(ctx context.Context, params parameter.RoleDeleteParam) (*entity.Role, error)
}

type role struct {
	psql    *generated.Client
	cfg     config.Config
	domRole domRole.Interface
}

func Init(psql *generated.Client, cfg config.Config, domRole domRole.Interface) Interface {
	result := role{
		psql:    psql,
		cfg:     cfg,
		domRole: domRole,
	}
	return &result
}

func (self *role) GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.Role], error) {
	params, offset, err := pagination.ParseFromParam(params)
	if err != nil {
		return nil, err
	}

	listRole, err := self.psql.Role.Query().
		Limit(int(params.Limit)).
		Offset(int(offset)).
		Where(psqlRole.IsDeleted(params.IsDeleted)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	totalRole, err := self.psql.Role.Query().Where(psqlRole.IsDeleted(params.IsDeleted)).Count(ctx)
	if err != nil {
		return nil, err
	}

	listRoleEntity := []entity.Role{}
	for _, v := range listRole {
		role, err := self.domRole.ToEntity(*v)
		if err != nil {
			return nil, err
		}
		listRoleEntity = append(listRoleEntity, role)
	}

	totalPages := uint64(totalRole / int(params.Limit))
	result := entity.ResponsePagination[entity.Pagination, []entity.Role]{
		Pagination: entity.Pagination{
			Total:       uint64(totalRole),
			Limit:       params.Limit,
			CurrentPage: params.Limit,
			TotalPages:  uint64(operator.Ternary(totalPages <= 0, 1, int(totalPages))),
		},
		Data: listRoleEntity,
	}

	return &result, nil
}

func (self *role) Get(ctx context.Context, params parameter.RoleGetParam) (*entity.Role, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not uuid"))
	}

	role, err := self.psql.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	result, err := self.domRole.ToEntity(*role)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *role) Create(ctx context.Context, params parameter.RoleCreateParam) (*entity.Role, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	role, err := tx.Role.Create().
		SetName(params.Name).
		SetDescription(params.Description).

		// TODO: set authentication for GRPC
		SetCreatedBy(uuid.New()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domRole.ToEntity(*role)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *role) Update(ctx context.Context, params parameter.RoleUpdateParam) (*entity.Role, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not uuid"))
	}

	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Role.Update().
		SetName(params.Name).
		SetDescription(params.Description).
		SetUpdatedAt(time.Now()).

		// TODO: set authentication for GRPC
		SetUpdatedBy(uuid.New()).
		Where(psqlRole.ID(id)).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	role, err := tx.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domRole.ToEntity(*role)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *role) Delete(ctx context.Context, params parameter.RoleDeleteParam) (*entity.Role, error) {
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not uuid"))
	}

	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	role, err := tx.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if role.IsDeleted == 1 {
		return nil, errors.New("role is already deleted")
	}

	if err := tx.Role.UpdateOneID(id).SetIsDeleted(1).Exec(ctx); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domRole.ToEntity(*role)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
