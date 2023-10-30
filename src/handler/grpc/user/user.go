package user

import (
	"context"

	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/handler/grpc/converter"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/protobuf/generated/pb"
)

type userService struct {
	pb.UnimplementedUserServiceServer
	user user.Interface
}

func Init(user user.Interface) pb.UserServiceServer {
	result := userService{
		user: user,
	}
	return &result
}

func (self *userService) GetUsers(ctx context.Context, request *pb.UserPaginationRequest) (*pb.UserResponsePagination, error) {
	args, err := converter.FromUserPaginationParam(request)
	if err != nil {
		return nil, err
	}

	result, err := self.user.GetListWithPagination(ctx, args)
	if err != nil {
		return nil, err

	}

	response, err := converter.ToUserResponsePagination(*result)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (self *userService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	panic("not implemented") // TODO: Implement
}

func (self *userService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.User, error) {
	panic("not implemented") // TODO: Implement
}

func (self *userService) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	panic("not implemented") // TODO: Implement
}

func (self *userService) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.User, error) {
	panic("not implemented") // TODO: Implement
}
