package main

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/irdaislakhuafa/learn-grpc-go/src/connection"
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

}
