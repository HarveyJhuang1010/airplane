package rdb

import (
	"gorm.io/gorm"
)

func newRepository(in dependence) *Repository {
	return &Repository{
		in: in,
	}
}

type Repository struct {
	in dependence
}

func (c *Repository) Master() *Database {
	return &Database{&Session{in: c.in, db: c.in.DBMaster.Session(&gorm.Session{})}}
}

func (c *Repository) ReadOnly() *Database {
	return &Database{&Session{in: c.in, db: c.in.DBMaster.Session(&gorm.Session{})}}
}

type Session struct {
	in dependence
	db *gorm.DB
}
