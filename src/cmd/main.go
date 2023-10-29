package main

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/connection"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/params"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
)

func main() {
	ctx := context.Background()
	cfg := config.Config{
		Database: config.Database{
			PSQL: config.SQL{
				Driver:   dialect.Postgres,
				Host:     "localhost",
				Port:     5432,
				DB:       "learn_grpc_go",
				User:     "postgres",
				Password: "postgres",
				SSL:      false,
			},
		},
	}

	psql, err := connection.ConnectPSQL(cfg)
	if err != nil {
		panic(err)
	} else if err := psql.Schema.Create(ctx); err != nil {
		panic(err)
	}

	// init usecase
	uc := usecase.Init(psql, cfg)

	rp, err := uc.User.GetListWithPagination(ctx, params.UserPaginationParam{Limit: 10, Page: 1, IsDeleted: 0})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rp)
}
