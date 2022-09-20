package ctxgorm

import (
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

/*
	Finisher API
*/

func (db *DB) Create(value any) error {
	err := db.d.Create(value).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) CreateInBatches(value any, batchSize int) error {
	err := db.d.CreateInBatches(value, batchSize).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

// Save update value in database, if the value doesn't have primary key, will insert it
func (db *DB) Save(value any) error {
	err := db.d.Save(value).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Exec(sql string, values ...any) error {
	err := db.d.Exec(sql, values).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) First(dest any) error {
	err := db.d.First(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Take(dest any) error {
	err := db.d.Take(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Last(dest any) error {
	err := db.d.Last(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Find(dest any) error {
	err := db.d.Find(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) FindInBatches(dest any, batchSize int, fc func(tx *gorm.DB, batch int) error) error {
	err := db.d.FindInBatches(dest, batchSize, fc).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) FirstOrInit(dest any) error {
	err := db.d.FirstOrInit(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) FirstOrCreate(dest any) error {
	err := db.d.FirstOrCreate(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Update(column string, value any) error {
	err := db.d.Update(column, value).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Updates(values any) error {
	err := db.d.Updates(values).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) UpdateColumn(column string, value any) error {
	err := db.d.UpdateColumn(column, value).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Delete(value any) error {
	err := db.d.Delete(value).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

func (db *DB) Count() (int64, error) {
	var count int64
	if err := db.d.Count(&count).Error; err != nil {
		return 0, errors.WithStackDepth(err, 1)
	}
	return count, nil
}

func (db *DB) Scan(dest any) error {
	err := db.d.Scan(dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}

// Pluck used to query single column from a model as a map
//     var ages []int64
//     db.Pluck(&users, "age", &ages)
func (db *DB) Pluck(model any, column string, dest any) error {
	err := db.d.Find(model).Pluck(column, dest).Error
	if err != nil {
		return errors.WithStackDepth(err, 1)
	}
	return nil
}
