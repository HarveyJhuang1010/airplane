package rdb

import (
	"gorm.io/gorm"
)

type Database struct {
	*Session
}

func (d *Database) Transaction(fn func(tx *Database) error) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		db := &Database{&Session{in: d.in, db: tx}}
		return fn(db)
	})
}
