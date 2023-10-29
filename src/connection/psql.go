package connection

import (
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"github.com/irdaislakhuafa/learn-grpc-go/src/schema/psqlentity/schema/generated"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/config"
	"github.com/irdaislakhuafa/learn-grpc-go/src/utils/operator"
	_ "github.com/lib/pq"
)

func ConnectPSQL(cfg config.Config) (*generated.Client, error) {
	p := cfg.Database.PSQL
	format := "host=%s port=%d user=%s dbname=%s password=%s sslmode=%s"
	datasource := fmt.Sprintf(format, p.Host, p.Port, p.User, p.DB, p.Password, operator.Ternary(p.SSL, "enable", "disable"))
	client, err := generated.Open(dialect.Postgres, datasource)
	if err != nil {
		return nil, errors.Join(errors.New("cannot connect to postgresql db"), err)
	}

	return client, nil
}
