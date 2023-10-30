package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/handler/grpc/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/protobuf/generated/pb"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	goGRPC "google.golang.org/grpc"
)

type GRPC struct {
	uc  usecase.Usecase
	cfg config.Config
}

func InitAndRun(uc usecase.Usecase, cfg config.Config) GRPC {
	result := GRPC{
		uc:  uc,
		cfg: cfg,
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", result.cfg.Service.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen GRPC server at port %v, %v", result.cfg.Service.GRPC.Port, err)
	}

	server := goGRPC.NewServer()

	result.register(server)

	log.Printf("GRPC Server started at port %v", result.cfg.Service.GRPC.Port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to server GRPC server at port %v, %v", result.cfg.Service.GRPC.Port, err)
	}

	return result
}

func (self *GRPC) register(server *goGRPC.Server) {
	pb.RegisterUserServiceServer(server, user.Init(self.uc.User))
}
