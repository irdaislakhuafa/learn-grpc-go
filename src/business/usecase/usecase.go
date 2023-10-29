package usecase

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Usecase struct {
	User user.Interface
}

func Init(psql *generated.Client, cfg config.Config) Usecase {
	result := Usecase{
		User: user.Init(psql, cfg),
	}
	return result
}
