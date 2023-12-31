package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	domAddress "github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/address"
	domUser "github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	psqlAddress "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/address"
	psqlUser "github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/operator"
)

type Interface interface {
	GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.User], error)
	Get(ctx context.Context, params parameter.UserGetParam) (*entity.User, error)
	Create(ctx context.Context, params parameter.UserCreateParam) (*entity.User, error)
	Update(ctx context.Context, params parameter.UserUpdateParam) (*entity.User, error)
	Delete(ctx context.Context, params parameter.UserDeleteParam) (*entity.User, error)
}

type user struct {
	psql       *generated.Client
	cfg        config.Config
	domUser    domUser.Interface
	domAddress domAddress.Interface
}

func Init(psql *generated.Client, cfg config.Config, domUser domUser.Interface, domAddress domAddress.Interface) Interface {
	result := user{
		psql:       psql,
		cfg:        cfg,
		domUser:    domUser,
		domAddress: domAddress,
	}
	return &result
}

func (self *user) GetListWithPagination(ctx context.Context, params parameter.PaginationParam) (*entity.ResponsePagination[entity.Pagination, []entity.User], error) {
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

	listAddress, err := self.psql.Address.Query().Where(psqlAddress.UserIDIn(listUserID...)).All(ctx)
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
		user, err := self.domUser.ToEntity(*v)
		if err != nil {
			return nil, err
		}

		address := mapListAddressByUserID[v.ID.String()]
		if user.Address, err = self.domAddress.ToEntity(address, user.ID); err != nil {
			return nil, err
		}

		listUserEntity = append(listUserEntity, user)
	}

	// count total data all user
	totalUser, err := self.psql.User.Query().Where(psqlUser.IsDeleted(params.IsDeleted)).Count(ctx)
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

func (self *user) Get(ctx context.Context, params parameter.UserGetParam) (*entity.User, error) {
	// parse id from parameter
	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New("id is not UUID"))
	}

	// get user by parsed id
	user, err := self.psql.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// find address of user
	address, err := self.psql.Address.Query().Where(psqlAddress.UserID(user.ID)).First(ctx)
	if err != nil {
		if !generated.IsNotFound(err) {
			return nil, err
		} else {
			address = new(generated.Address)
		}

	}

	result, err := self.domUser.ToEntity(*user)
	if err != nil {
		return nil, err
	}

	if result.Address, err = self.domAddress.ToEntity(*address, result.ID); err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *user) Create(ctx context.Context, params parameter.UserCreateParam) (*entity.User, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	user, err := tx.User.Create().
		SetName(params.Name).
		SetEmail(params.Email).
		SetAge(params.Age).
		SetHobbies(params.Hobbies).
		SetCreatedAt(time.Now()).

		// TODO: set authentication for GRPC
		SetCreatedBy(uuid.New()).
		Save(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot create user"))
	}

	address, err := tx.Address.Create().
		SetCountry(params.Address.Country).
		SetRegency(params.Address.Regency).
		SetProvince(params.Address.Province).
		SetSubDistrict(params.Address.SubDistrict).
		SetUserID(user.ID).

		// TODO: set authentication for GRPC
		SetCreatedBy(uuid.New()).
		Save(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot create address"))
	}

	if err := tx.Commit(); err != nil {
		return nil, errors.Join(err, errors.New("cannot commit create user"))
	}

	result, err := self.domUser.ToEntity(*user)
	if err != nil {
		return nil, err
	}

	if result.Address, err = self.domAddress.ToEntity(*address, result.ID); err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *user) Update(ctx context.Context, params parameter.UserUpdateParam) (*entity.User, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, errors.Join(err, errors.New(fmt.Sprintf("id is not uuid")))
	}

	err = tx.User.UpdateOneID(id).
		SetName(params.Name).
		SetEmail(params.Email).
		SetAge(params.Age).
		SetHobbies(params.Hobbies).
		SetUpdatedAt(time.Now()).

		// TODO: set authentication for GRPC
		SetUpdatedBy(uuid.New()).
		Exec(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot update user"))
	}

	user, err := tx.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	err = tx.Address.Update().
		SetCountry(params.Address.Country).
		SetRegency(params.Address.Regency).
		SetSubDistrict(params.Address.SubDistrict).
		SetProvince(params.Address.Province).
		SetUpdatedAt(time.Now()).

		// TODO: set authentication for GRPC
		SetUpdatedBy(uuid.New()).
		Where(psqlAddress.UserID(user.ID)).
		Exec(ctx)
	if err != nil {
		return nil, errors.Join(err, errors.New("cannot update address of user"))
	}

	address, err := tx.Address.Query().Where(psqlAddress.UserID(user.ID)).First(ctx)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domUser.ToEntity(*user)
	if err != nil {
		return nil, err
	}

	if result.Address, err = self.domAddress.ToEntity(*address, result.ID); err != nil {
		return nil, err
	}

	return &result, nil
}

func (self *user) Delete(ctx context.Context, params parameter.UserDeleteParam) (*entity.User, error) {
	tx, err := self.psql.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	id, err := uuid.Parse(params.ID)
	if err != nil {
		return nil, err
	}

	user, err := tx.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if user.IsDeleted == 1 {
		return nil, errors.New("user is already deleted")
	}

	if err := tx.User.UpdateOneID(id).SetIsDeleted(1).Exec(ctx); err != nil {
		return nil, errors.Join(err, errors.New("cannot delete user"))
	}

	address, err := tx.Address.Query().Where(psqlAddress.UserID(user.ID)).First(ctx)
	if err != nil {
		if !generated.IsNotFound(err) {
			return nil, err
		} else {
			address = new(generated.Address)
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	result, err := self.domUser.ToEntity(*user)
	if err != nil {
		return nil, err
	}

	if result.Address, err = self.domAddress.ToEntity(*address, result.ID); err != nil {
		return nil, err
	}
	return &result, nil
}
