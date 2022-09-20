package ctxgorm

import (
	"gorm.io/gorm"
)

type DB struct {
	d *gorm.DB
}

func (db *DB) DB() *gorm.DB {
	return db.d
}
