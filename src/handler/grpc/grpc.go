package grpc

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type GRPC struct {
	uc  usecase.Usecase
	cfg config.Config
}

func Init(uc usecase.Usecase, cfg config.Config) GRPC {
	result := GRPC{
		uc:  uc,
		cfg: cfg,
	}
	return result
}
