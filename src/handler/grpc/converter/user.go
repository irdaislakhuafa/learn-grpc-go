package converter

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/parameter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/protobuf/generated/pb"
)

func ToUserPaginationParam(request *pb.UserPaginationRequest) (parameter.UserPaginationParam, error) {
	result := parameter.UserPaginationParam{
		Limit:     request.GetLimit(),
		Page:      request.GetPage(),
		IsDeleted: int64(request.GetIsDeleted()),
	}

	return result, nil
}

func ToUserResponsePaginationProto(request entity.ResponsePagination[entity.Pagination, []entity.User]) (*pb.UserResponsePagination, error) {
	result := &pb.UserResponsePagination{
		Pagination: &pb.Pagination{
			TotalData:   request.Pagination.Total,
			CurrentPage: request.Pagination.CurrentPage,
			Limit:       request.Pagination.Limit,
			TotalPages:  request.Pagination.TotalPages,
		},
		Data: []*pb.User{},
	}

	for _, v := range request.Data {
		user, err := ToUserProto(v)
		if err != nil {
			return nil, err
		}

		result.Data = append(result.Data, user)
	}

	return result, nil
}

func ToUserGetParam(request *pb.GetUserRequest) (parameter.UserGetParam, error) {
	result := parameter.UserGetParam{ID: request.Id}
	return result, nil
}

func ToUserProto(v entity.User) (*pb.User, error) {
	result := &pb.User{
		Id:      v.ID.String(),
		Name:    v.Name,
		Email:   v.Email,
		Age:     uint64(v.Age),
		Hobbies: v.Hobbies,
		Address: &pb.Address{
			Id:          v.Address.ID.String(),
			Country:     v.Address.Country,
			Province:    v.Address.Province,
			Regency:     v.Address.Regency,
			SubDistrict: v.Address.SubDistrict,
		},
	}
	return result, nil
}

func ToUserCreateParam(v *pb.CreateUserRequest) (parameter.UserCreateParam, error) {
	result := parameter.UserCreateParam{
		Name:    v.Name,
		Email:   v.Email,
		Age:     int(v.Age),
		Hobbies: v.Hobbies,
		Address: parameter.AddressCreateParam{
			Country:     v.Address.Country,
			Province:    v.Address.Province,
			Regency:     v.Address.Regency,
			SubDistrict: v.Address.SubDistrict,
		},
	}
	return result, nil
}

func ToUserUpdateParam(v *pb.UpdateUserRequest) (parameter.UserUpdateParam, error) {
	result := parameter.UserUpdateParam{
		ID:      v.GetId(),
		Name:    v.GetEmail(),
		Email:   v.GetEmail(),
		Age:     int(v.GetAge()),
		Hobbies: v.GetHobbies(),
		Address: parameter.AddressUpdateParam{
			ID:          v.Address.GetId(),
			Country:     v.Address.GetCountry(),
			Province:    v.Address.GetProvince(),
			Regency:     v.Address.GetRegency(),
			SubDistrict: v.Address.GetSubDistrict(),
		},
	}
	return result, nil
}
