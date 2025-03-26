package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/tools/timelogger"
	"context"
)

func (d *Database) ExtraPaymentDAO() IExtraPaymentDao {
	return newExtraPaymentDao(d.Session)
}

func newExtraPaymentDao(session *Session) *ExtraPaymentDao {
	return &ExtraPaymentDao{
		Session: session,
	}
}

type IExtraPaymentDao interface {
	Create(ctx context.Context, cond ...*po.ExtraPayment) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.ExtraPayment, error)
	GetByBookingID(ctx context.Context, id int64, forUpdate bool) (*po.ExtraPayment, error)
}

type ExtraPaymentDao struct {
	*Session
	model po.ExtraPayment
}

// Create 新增操作紀錄
func (dao *ExtraPaymentDao) Create(ctx context.Context, cond ...*po.ExtraPayment) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *ExtraPaymentDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.ExtraPayment, error) {
	defer timelogger.LogTime(ctx)()

	var model po.ExtraPayment

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model, id).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *ExtraPaymentDao) GetByBookingID(ctx context.Context, id int64, forUpdate bool) (*po.ExtraPayment, error) {
	defer timelogger.LogTime(ctx)()

	var model po.ExtraPayment

	query := dao.db.Where(`booking_id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model, id).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}
