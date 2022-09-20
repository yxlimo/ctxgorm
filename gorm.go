package ctxgorm

import (
	"context"

	"gorm.io/gorm"
)

var dbKey = struct{}{}

func ToContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey, db)
}

func Ctx(ctx context.Context) *DB {
	db, ok := ctx.Value(dbKey).(*gorm.DB)
	if !ok {
		panic("gorm db not exist in context")
	}
	return &DB{d: db.WithContext(ctx)}
}

func Transaction(ctx context.Context, fn func(context.Context) error) error {
	db := ctx.Value(dbKey).(*gorm.DB)
	if err := db.Transaction(func(tx *gorm.DB) error {
		newCtx := ToContext(ctx, tx)
		return fn(newCtx)
	}); err != nil {
		return err
	}
	return nil
}
