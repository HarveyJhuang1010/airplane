package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (d *Database) SeatDAO() ISeatDao {
	return newSeatDao(d.Session)
}

func newSeatDao(session *Session) *SeatDao {
	return &SeatDao{
		Session: session,
	}
}

type ISeatDao interface {
	Create(ctx context.Context, cond ...*po.Seat) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.Seat, error)
	Update(ctx context.Context, cond *po.SeatUpdateCond) error
	List(ctx context.Context, cond *po.SeatListCond) ([]*po.Seat, error)
	ListPager(ctx context.Context, cond *po.SeatListCond) (*po.Pagination, error)
}

type SeatDao struct {
	*Session
	model po.Seat
}

// Create 新增操作紀錄
func (dao *SeatDao) Create(ctx context.Context, cond ...*po.Seat) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *SeatDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.Seat, error) {
	defer timelogger.LogTime(ctx)()

	var model po.Seat

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model, id).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *SeatDao) Update(ctx context.Context, cond *po.SeatUpdateCond) error {
	defer timelogger.LogTime(ctx)()

	updates := func() map[string]interface{} {
		upd := make(map[string]interface{})
		if !lo.IsNil(cond.Status) {
			upd["status"] = cond.Status
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

func (dao *SeatDao) List(ctx context.Context, cond *po.SeatListCond) ([]*po.Seat, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.Seat

	query := dao.db.Scopes(dao.list(cond, cond.Pager))
	if err := query.Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}

// ListPager get total count by pager
func (dao *SeatDao) ListPager(ctx context.Context, cond *po.SeatListCond) (*po.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	var count int64

	if err := dao.db.
		Scopes(dao.list(cond, cond.Pager)).Count(&count).
		Error; err != nil {
		return nil, err
	}

	return po.NewPagination(cond.Pager, count), nil
}

func (dao *SeatDao) list(cond *po.SeatListCond, pager *po.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !lo.IsNil(cond) && !lo.IsNil(cond.FlightID) && !lo.IsEmpty(cond.FlightID) {
			db = db.Where("flight_id = ?", cond.FlightID)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.Status) && !lo.IsEmpty(cond.Status) {
			db = db.Where("status = ?", cond.Status)
		}

		if pager == nil {
			return db.Model(dao.model)
		} else {
			return db.Model(dao.model).Scopes(dao.in.common.ParsePaging(pager))
		}
	}
}
