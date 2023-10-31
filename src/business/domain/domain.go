package domain

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/address"
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/sale"
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/user"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Domain struct {
	Sale    sale.Interface
	Address address.Interface
	User    user.Interface
}

func Init(psql *generated.Client, cfg config.Config) Domain {
	result := Domain{
		Sale:    sale.Init(psql, cfg),
		Address: address.Init(psql, cfg),
		User:    user.Init(psql, cfg),
	}
	return result
}
