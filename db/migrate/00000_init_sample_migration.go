package migrate

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	_ = migrate.Register(func(ctx context.Context, db *mongo.Database) error {
		return nil
	}, func(ctx context.Context, db *mongo.Database) error {
		return nil
	})
}
