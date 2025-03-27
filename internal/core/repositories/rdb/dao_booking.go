package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

func (d *Database) BookingDAO() IBookingDao {
	return newBookingDao(d.Session)
}

func newBookingDao(session *Session) *BookingDao {
	return &BookingDao{
		Session: session,
	}
}

type IBookingDao interface {
	Create(ctx context.Context, cond ...*po.Booking) error
	Get(ctx context.Context, id int64, forUpdate, preload bool) (*po.Booking, error)
	UpdateStatus(ctx context.Context, cond *po.BookingUpdateCond) error
	UpdateSeat(ctx context.Context, cond *po.BookingUpdateSeatCond) error
	CancelSeat(ctx context.Context, id int64) error
	List(ctx context.Context, cond *po.BookingListCond) ([]*po.Booking, error)
	ListPager(ctx context.Context, cond *po.BookingListCond) (*po.Pagination, error)
	GetExpired(ctx context.Context, expiredAt time.Time) ([]*po.Booking, error)
	GetOverBooking(ctx context.Context, flightID int64) ([]*po.Booking, error)
}

type BookingDao struct {
	*Session
	model po.Booking
}

// Create 新增操作紀錄
func (dao *BookingDao) Create(ctx context.Context, cond ...*po.Booking) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *BookingDao) Get(ctx context.Context, id int64, forUpdate, preload bool) (*po.Booking, error) {
	defer timelogger.LogTime(ctx)()

	var model po.Booking
	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}
	if preload {
		query = query.Preload("Flight").
			Preload("User").
			Preload("Class").
			Preload("Seat")
	}

	if err := query.First(&model).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *BookingDao) UpdateStatus(ctx context.Context, cond *po.BookingUpdateCond) error {
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

func (dao *BookingDao) UpdateSeat(ctx context.Context, cond *po.BookingUpdateSeatCond) error {
	defer timelogger.LogTime(ctx)()

	updates := func() map[string]interface{} {
		upd := make(map[string]interface{})
		if !lo.IsNil(cond.SeatID) {
			upd["seat_id"] = cond.SeatID
		}
		if !lo.IsNil(cond.CabinClassID) {
			upd["cabin_class_id"] = cond.CabinClassID
		}
		if !lo.IsNil(cond.Price) {
			upd["price"] = cond.Price
		}
		return upd
	}()

	if err := dao.db.Model(dao.model).
		Where("`id` = ?", cond.ID).
		Updates(updates).
		Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *BookingDao) CancelSeat(ctx context.Context, id int64) error {
	defer timelogger.LogTime(ctx)()

	upd := make(map[string]interface{})
	upd["seat_id"] = nil

	if err := dao.db.Model(dao.model).
		Where("`id` = ?", id).
		Updates(upd).
		Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *BookingDao) List(ctx context.Context, cond *po.BookingListCond) ([]*po.Booking, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.Booking

	query := dao.db.Scopes(dao.list(cond, cond.Pager))
	if err := query.Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}

// ListPager get total count by pager
func (dao *BookingDao) ListPager(ctx context.Context, cond *po.BookingListCond) (*po.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	var count int64

	if err := dao.db.
		Scopes(dao.list(cond, cond.Pager)).Count(&count).
		Error; err != nil {
		return nil, err
	}

	return po.NewPagination(cond.Pager, count), nil
}

func (dao *BookingDao) list(cond *po.BookingListCond, pager *po.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !lo.IsNil(cond) && !lo.IsNil(cond.FlightID) && !lo.IsEmpty(cond.FlightID) {
			db = db.Where("flight_id = ?", cond.FlightID)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.UserID) && !lo.IsEmpty(cond.UserID) {
			db = db.Where("user_id = ?", cond.UserID)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.Status) && len(cond.Status) > 0 {
			db = db.Where("status IN (?)", cond.Status)
		}

		if pager == nil {
			return db.Model(dao.model)
		} else {
			return db.Model(dao.model).Scopes(dao.in.common.ParsePaging(pager))
		}
	}
}

func (dao *BookingDao) GetExpired(ctx context.Context, expiredAt time.Time) ([]*po.Booking, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.Booking

	if err := dao.db.Where(
		"status = ? AND expired_at <= ?",
		enum.BookingStatusPending,
		expiredAt).Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}

func (dao *BookingDao) GetOverBooking(ctx context.Context, flightID int64) ([]*po.Booking, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.Booking

	if err := dao.db.Where(
		"flight_id = ? AND status = ?",
		flightID,
		enum.BookingStatusPending).
		Order("created_at ASC").
		Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}
