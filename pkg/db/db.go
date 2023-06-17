package db

import (
	"aramen-api/pkg/config"
	"aramen-api/pkg/db/seeder"
	"context"
	"fmt"
	"log"

	"aramen-api/cmd/api/ent"

	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection(cfg *config.AppConfig) (*ent.Client, func(), error) {
	client, err := ent.Open("mysql", getDsn(cfg.Database))
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	if err := client.Schema.Create(context.Background(), schema.WithHooks(func(next schema.Creator) schema.Creator {
		return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
			if err := next.Create(ctx, tables...); err != nil {
				return err
			}

			if err := seeder.New(client).Seed(context.Background()); err != nil {
				return err
			}

			return nil
		})
	})); err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		client.Close()
	}

	return client, cleanup, err
}

func getDsn(config config.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}
