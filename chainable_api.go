package ctxgorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

/*
	Chainable API
*/
func (db *DB) Where(query any, args ...any) *DB {
	return &DB{db.d.Where(query, args...)}
}

func (db *DB) ID(ids ...any) *DB {
	return &DB{db.d.Where("id IN (?)", ids)}
}

func (db *DB) Model(value any) *DB {
	return &DB{db.d.Model(value)}
}

func (db *DB) Table(name string, args ...any) *DB {
	return &DB{db.d.Table(name, args...)}
}

func (db *DB) Distinct(args ...any) *DB {
	return &DB{db.d.Distinct(args...)}
}

func (db *DB) Select(query any, args ...any) *DB {
	return &DB{db.d.Select(query, args...)}
}

func (db *DB) Omit(columns ...string) *DB {
	return &DB{db.d.Omit(columns...)}
}

func (db *DB) Not(query any, args ...any) *DB {
	return &DB{db.d.Not(query, args...)}
}

func (db *DB) Or(query any, args ...any) *DB {
	return &DB{db.d.Or(query, args...)}
}

func (db *DB) Joins(query string, args ...any) *DB {
	return &DB{db.d.Joins(query, args...)}
}

func (db *DB) Group(name string) *DB {
	return &DB{db.d.Group(name)}
}

func (db *DB) Having(query any, args ...any) *DB {
	return &DB{db.d.Having(query, args...)}
}

func (db *DB) Order(value any) *DB {
	return &DB{db.d.Order(value)}
}

// Scopes 原来会在最后执行，这里改为按顺序直接执行，语义上更明确一些
func (db *DB) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *DB {
	tx := db.d
	for _, fn := range funcs {
		tx = fn(tx)
	}
	return &DB{d: tx}
}

func (db *DB) Preload(query string, args ...any) *DB {
	return &DB{db.d.Preload(query, args...)}
}

func (db *DB) Attrs(attrs ...any) *DB {
	return &DB{db.d.Attrs(attrs...)}
}

func (db *DB) Assign(attrs ...any) *DB {
	return &DB{db.d.Assign(attrs...)}
}

func (db *DB) Unscoped() *DB {
	return &DB{db.d.Unscoped()}
}

func (db *DB) Raw(sql string, values ...any) *DB {
	return &DB{db.d.Raw(sql, values...)}
}

func (db *DB) Limit(limit int) *DB {
	return &DB{db.d.Limit(limit)}
}

func (db *DB) Offset(offset int) *DB {
	return &DB{db.d.Offset(offset)}
}

func (db *DB) Clauses(conds ...clause.Expression) *DB {
	return &DB{db.d.Clauses(conds...)}
}
