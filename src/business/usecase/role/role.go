package role

import (
	"context"
	"errors"

	"github.com/google/uuid"
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
	Create(ctx context.Context, params parameter.RoleCreateParam) (*entity.User, error)
	Update(ctx context.Context, params parameter.RoleUpdateParam) (*entity.User, error)
	Delete(ctx context.Context, params parameter.RoleDeleteParam) (*entity.User, error)
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
		role := entity.Role{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
			CreatedBy:   v.CreatedBy,
			UpdatedAt:   v.UpdatedAt,
			UpdatedBy:   v.UpdatedBy,
			DeletedAt:   v.DeletedAt,
			DeletedBy:   v.DeletedBy,
			IsDeleted:   v.IsDeleted,
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

	result := entity.Role{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		CreatedBy:   role.CreatedBy,
		UpdatedAt:   role.UpdatedAt,
		UpdatedBy:   role.UpdatedBy,
		DeletedAt:   role.DeletedAt,
		DeletedBy:   role.DeletedBy,
		IsDeleted:   role.IsDeleted,
	}

	return &result, nil
}

func (self *role) Create(ctx context.Context, params parameter.RoleCreateParam) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (self *role) Update(ctx context.Context, params parameter.RoleUpdateParam) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}

func (self *role) Delete(ctx context.Context, params parameter.RoleDeleteParam) (*entity.User, error) {
	panic("not implemented") // TODO: Implement
}
