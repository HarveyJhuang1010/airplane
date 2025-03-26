package rdb

import (
	"errors"

	"airplane/internal/domain/entities/po"
	"airplane/internal/errs"
	"github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type common struct {
}

func (common) ErrorHandle(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return errs.ErrDuplicateRecord
	}
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errs.ErrRecordNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errs.ErrDuplicateRecord
	}

	return err
}

func (common) ParsePaging(pager *po.Pager) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pager != nil {
			db = db.Limit(pager.GetSize()).Offset(pager.GetOffset())
		}

		return db
	}
}

func (common) SetForUpdate() func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	}
}

// DirectDelete 直接刪除
func (common) DirectDelete() func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}
}

// WithDeleted 查詢軟刪除
func (common) WithDeleted() func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}

}

func (common) OrderBy(orderBy ...*po.OrderBy) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(orderBy) > 0 {
			for _, ob := range orderBy {
				if ob == nil {
					continue
				}
				db = db.Order("`" + ob.Column + "` " + lo.Ternary(ob.Desc, "DESC", "ASC"))
			}
		}

		return db
	}
}

func (common) TimeRange(column string, timeRange *po.TimeRange) func(dc *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if timeRange == nil {
			return db
		}

		if timeRange.Start != nil {
			db = db.Where(column+" >= ?", timeRange.Start)
		}
		if timeRange.End != nil {
			db = db.Where(column+" < ?", timeRange.End)
		}

		return db
	}
}
