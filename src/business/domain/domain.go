package domain

import (
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/domain/sale"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

type Domain struct {
	Sale sale.Interface
}

func Init(psql *generated.Client, cfg config.Config) Domain {
	result := Domain{
		Sale: sale.Init(psql, cfg),
	}
	return result
}
