package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/tools/timelogger"
	"context"
)

func (d *Database) CabinClassDAO() ICabinClassDao {
	return newCabinClassDao(d.Session)
}

func newCabinClassDao(session *Session) *CabinClassDao {
	return &CabinClassDao{
		Session: session,
	}
}

type ICabinClassDao interface {
	Create(ctx context.Context, cond ...*po.CabinClass) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.CabinClass, error)
	Update(ctx context.Context, id int64, remainCount int) error
}

type CabinClassDao struct {
	*Session
	model po.CabinClass
}

// Create 新增操作紀錄
func (dao *CabinClassDao) Create(ctx context.Context, cond ...*po.CabinClass) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *CabinClassDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.CabinClass, error) {
	defer timelogger.LogTime(ctx)()

	var result po.CabinClass

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &result, nil
}

func (dao *CabinClassDao) Update(ctx context.Context, id int64, remainSeats int) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Model(&dao.model).Where("id = ?", id).Update("remain_seats", remainSeats).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}
