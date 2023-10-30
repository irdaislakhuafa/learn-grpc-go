package converter

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/entity"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/params"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/protobuf/generated/pb"
)

func FromUserPaginationParam(request *pb.UserPaginationRequest) (params.UserPaginationParam, error) {
	result := params.UserPaginationParam{
		Limit:     request.GetLimit(),
		Page:      request.GetPage(),
		IsDeleted: int64(request.GetIsDeleted()),
	}

	return result, nil
}

func ToUserResponsePagination(request entity.ResponsePagination[entity.Pagination, []entity.User]) (*pb.UserResponsePagination, error) {
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
		user := pb.User{
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

		result.Data = append(result.Data, &user)
	}

	return result, nil
}
