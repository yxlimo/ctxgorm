package ctxgorm

import (
	"context"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ModelWrapper[T any] struct {
	model *T
	tx    *DB
}

func Model[T any](ctx context.Context, model *T) *ModelWrapper[T] {
	return &ModelWrapper[T]{tx: Ctx(ctx).Model(model), model: model}
}

/*
	Chainable API
*/
func (m *ModelWrapper[T]) Where(query any, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Where(query, args...)
	return m
}

func (m *ModelWrapper[T]) ID(ids ...any) *ModelWrapper[T] {
	m.tx = m.tx.ID(ids...)
	return m
}

func (m *ModelWrapper[T]) IncrID(ids ...uint64) *ModelWrapper[T] {
	m.tx = m.tx.Where("id IN (?)", ids)
	return m
}

func (m *ModelWrapper[T]) Distinct(args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Distinct(args...)
	return m
}

func (m *ModelWrapper[T]) Select(query any, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Select(query, args...)
	return m
}

func (m *ModelWrapper[T]) Omit(columns ...string) *ModelWrapper[T] {
	m.tx = m.tx.Omit(columns...)
	return m
}

func (m *ModelWrapper[T]) Not(query any, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Not(query, args...)
	return m
}

func (m *ModelWrapper[T]) Or(query any, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Or(query, args...)
	return m
}

func (m *ModelWrapper[T]) Joins(query string, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Joins(query, args...)
	return m
}

func (m *ModelWrapper[T]) Group(name string) *ModelWrapper[T] {
	m.tx = m.tx.Group(name)
	return m
}

func (m *ModelWrapper[T]) Having(query any, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Having(query, args...)
	return m
}

func (m *ModelWrapper[T]) Order(value any) *ModelWrapper[T] {
	m.tx = m.tx.Order(value)
	return m
}

// Scopes 原来会在最后执行，这里改为按顺序直接执行，语义上更明确一些
func (m *ModelWrapper[T]) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *ModelWrapper[T] {
	for _, fn := range funcs {
		m.tx.d = fn(m.tx.d)
	}
	return m
}

func (m *ModelWrapper[T]) Preload(query string, args ...any) *ModelWrapper[T] {
	m.tx = m.tx.Preload(query, args...)
	return m
}

func (m *ModelWrapper[T]) Attrs(attrs ...any) *ModelWrapper[T] {
	m.tx = m.tx.Attrs(attrs...)
	return m
}

func (m *ModelWrapper[T]) Assign(attrs ...any) *ModelWrapper[T] {
	m.tx = m.tx.Assign(attrs...)
	return m
}

func (m *ModelWrapper[T]) Unscoped() *ModelWrapper[T] {
	m.tx = m.tx.Unscoped()
	return m
}

func (m *ModelWrapper[T]) Raw(sql string, values ...any) *ModelWrapper[T] {
	m.tx = m.tx.Raw(sql, values...)
	return m
}

func (m *ModelWrapper[T]) Limit(limit int) *ModelWrapper[T] {
	m.tx = m.tx.Limit(limit)
	return m
}

func (m *ModelWrapper[T]) Offset(offset int) *ModelWrapper[T] {
	m.tx = m.tx.Offset(offset)
	return m
}

func (m *ModelWrapper[T]) Clauses(conds ...clause.Expression) *ModelWrapper[T] {
	m.tx = m.tx.Clauses(conds...)
	return m
}

/*
	Finisher API
*/

func (m *ModelWrapper[T]) Create(value ...*T) error {
	if len(value) > 0 {
		return m.tx.Create(value)
	}
	if err := m.tx.Create(m.model); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) CreateInBatches(value []*T, batchSize int) error {
	if err := m.tx.CreateInBatches(value, batchSize); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

// // Save update value in database, if the value doesn't have primary key, will insert it
func (m *ModelWrapper[T]) Save(value ...*T) error {
	if len(value) > 0 {
		return m.tx.Save(value)
	}
	if err := m.tx.Save(m.model); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) First() (*T, error) {
	if err := m.tx.First(m.model); err != nil {
		return m.model, errors.WithStackDepth(err, 1)
	}
	return m.model, nil
}

func (m *ModelWrapper[T]) Take() (*T, error) {
	if err := m.tx.Take(m.model); err != nil {
		return m.model, errors.WithStackDepth(err, 1)
	}
	return m.model, nil
}

func (m *ModelWrapper[T]) Last() (*T, error) {
	if err := m.tx.Last(m.model); err != nil {
		return nil, errors.WithStackDepth(err, 1)
	}
	return m.model, nil
}

func (m *ModelWrapper[T]) Find() ([]*T, error) {
	dest := []*T{}
	if err := m.tx.Find(&dest); err != nil {
		return nil, errors.WithStackDepth(err, 1)
	}
	return dest, nil
}

func (m *ModelWrapper[T]) FirstOrInit() (*T, error) {
	if err := m.tx.FirstOrInit(m.model); err != nil {
		return nil, errors.WithStackDepth(err, 1)
	}
	return m.model, nil
}

func (m *ModelWrapper[T]) FirstOrCreate() (*T, error) {
	if err := m.tx.FirstOrCreate(m.model); err != nil {
		return nil, errors.WithStackDepth(err, 1)
	}
	return m.model, nil
}

func (m *ModelWrapper[T]) Update(column string, value any) error {
	if err := m.tx.Update(column, value); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) Updates(values map[string]any) error {
	if err := m.tx.Updates(values); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) UpdateColumn(column string, value any) error {
	if err := m.tx.UpdateColumn(column, value); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) Delete() error {
	if err := m.tx.Delete(m.model); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (m *ModelWrapper[T]) Count() (int64, error) {
	count, err := m.tx.Count()
	if err != nil {
		return 0, errors.WithStackDepth(err, 1)
	}
	return count, nil
}

func (m *ModelWrapper[T]) Pluck(column string, dest any) error {
	if err := m.tx.Pluck(m.model, column, dest); err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}
