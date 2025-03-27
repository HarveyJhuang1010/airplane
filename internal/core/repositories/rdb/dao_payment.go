package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (d *Database) PaymentDAO() IPaymentDao {
	return newPaymentDao(d.Session)
}

func newPaymentDao(session *Session) *PaymentDao {
	return &PaymentDao{
		Session: session,
	}
}

type IPaymentDao interface {
	Create(ctx context.Context, cond ...*po.Payment) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.Payment, error)
	GetByBookingID(ctx context.Context, id int64, forUpdate bool) (*po.Payment, error)
	UpdateStatus(ctx context.Context, id int64, status enum.PaymentStatus) error
	UpdateAmount(ctx context.Context, id int64, amount decimal.Decimal) error
	UpdateResult(ctx context.Context, cond *po.PaymentUpdateResultCond) error
}

type PaymentDao struct {
	*Session
	model po.Payment
}

// Create 新增操作紀錄
func (dao *PaymentDao) Create(ctx context.Context, cond ...*po.Payment) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *PaymentDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.Payment, error) {
	defer timelogger.LogTime(ctx)()

	var model po.Payment

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *PaymentDao) GetByBookingID(ctx context.Context, id int64, forUpdate bool) (*po.Payment, error) {
	defer timelogger.LogTime(ctx)()

	var model po.Payment

	query := dao.db.Where(`booking_id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *PaymentDao) UpdateStatus(ctx context.Context, id int64, status enum.PaymentStatus) error {
	defer timelogger.LogTime(ctx)()

	updates := func() map[string]interface{} {
		upd := make(map[string]interface{})
		upd["status"] = status
		return upd
	}()

	if err := dao.db.Model(dao.model).
		Where("`id` = ?", id).Updates(updates).
		Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *PaymentDao) UpdateAmount(ctx context.Context, id int64, amount decimal.Decimal) error {
	defer timelogger.LogTime(ctx)()

	updates := func() map[string]interface{} {
		upd := make(map[string]interface{})
		upd["amount"] = amount
		return upd
	}()

	if err := dao.db.Model(dao.model).
		Where("`id` = ?", id).Updates(updates).
		Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *PaymentDao) UpdateResult(ctx context.Context, cond *po.PaymentUpdateResultCond) error {
	defer timelogger.LogTime(ctx)()

	updates := func() map[string]interface{} {
		upd := make(map[string]interface{})
		upd["status"] = cond.Status
		if !lo.IsNil(cond.TransactionID) {
			upd["transaction_id"] = cond.TransactionID
		}

		if !lo.IsNil(cond.Provider) {
			upd["provider"] = cond.Provider
		}

		if !lo.IsNil(cond.Method) {
			upd["method"] = cond.Method
		}

		if !lo.IsNil(cond.PaidAt) {
			upd["paid_at"] = cond.PaidAt
		}

		return upd
	}()

	if err := dao.db.Model(dao.model).
		Where("`id` = ?", cond.ID).Updates(updates).
		Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}
