package main

import (
	"context"
	"fmt"

	"github.com/irdaislakhuafa/learn-grpc-go/src/business/usecase"
	"github.com/irdaislakhuafa/learn-grpc-go/src/connection"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/params"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/flags"
)

const (
	configPath = "etc/cfg"
	configFile = "conf.json"
)

func main() {
	ctx := context.Background()
	env, err := flags.ParseFlags(configPath, configFile)
	if err != nil {
		panic(err)
	}

	cfg, err := config.ReadConfigFromFile[config.Config](fmt.Sprintf("%s/%s/%s", configPath, *env, configFile))
	if err != nil {
		panic(err)
	}

	psql, err := connection.ConnectPSQL(*cfg)
	if err != nil {
		panic(err)
	} else if err := psql.Schema.Create(ctx); err != nil {
		panic(err)
	}

	// init usecase
	uc := usecase.Init(psql, *cfg)

	rp, err := uc.User.GetListWithPagination(ctx, params.UserPaginationParam{Limit: 10, Page: 1, IsDeleted: 0})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", rp)
}
