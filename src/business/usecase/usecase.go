package usecase

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/domain"
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase/sale"
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Usecase struct {
	User user.Interface
	Sale sale.Interface
}

func Init(psql *generated.Client, cfg config.Config, dom domain.Domain) Usecase {
	result := Usecase{
		User: user.Init(psql, cfg),
		Sale: sale.Init(psql, cfg, dom.Sale),
	}
	return result
}
