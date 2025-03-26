package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

func (d *Database) FlightDAO() IFlightDao {
	return newFlightDao(d.Session)
}

func newFlightDao(session *Session) *flightDao {
	return &flightDao{
		Session: session,
	}
}

type IFlightDao interface {
	Create(ctx context.Context, cond ...*po.Flight) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.Flight, error)
	List(ctx context.Context, cond *po.FlightListCond) ([]*po.Flight, error)
	ListPager(ctx context.Context, cond *po.FlightListCond) (*po.Pagination, error)
	UpdateSellableSeats(ctx context.Context, id int64, sellableSeats int) error
}

type flightDao struct {
	*Session
	model po.Flight
}

// Create 新增操作紀錄
func (dao *flightDao) Create(ctx context.Context, cond ...*po.Flight) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *flightDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.Flight, error) {
	defer timelogger.LogTime(ctx)()

	var result po.Flight

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &result, nil
}

func (dao *flightDao) List(ctx context.Context, cond *po.FlightListCond) ([]*po.Flight, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.Flight

	query := dao.db.Scopes(dao.list(cond, cond.Pager))
	if cond.PreloadCabinClasses {
		query = query.Preload("CabinClasses")
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}

// ListPager get total count by pager
func (dao *flightDao) ListPager(ctx context.Context, cond *po.FlightListCond) (*po.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	var count int64

	if err := dao.db.
		Scopes(dao.list(cond, cond.Pager)).Count(&count).
		Error; err != nil {
		return nil, err
	}

	return po.NewPagination(cond.Pager, count), nil
}

func (dao *flightDao) list(cond *po.FlightListCond, pager *po.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !lo.IsNil(cond) && !lo.IsNil(cond.DepartureAirport) && !lo.IsEmpty(cond.DepartureAirport) {
			db = db.Where("departure_airport = ?", cond.DepartureAirport)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.ArrivalAirport) && !lo.IsEmpty(cond.ArrivalAirport) {
			db = db.Where("arrival_airport = ?", cond.ArrivalAirport)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.DepartureTimeStartAt) && !lo.IsEmpty(cond.DepartureTimeStartAt) {
			db = db.Where("departure_time >= ?", cond.DepartureTimeStartAt)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.DepartureTimeEndAt) && !lo.IsEmpty(cond.DepartureTimeEndAt) {
			db = db.Where("departure_time < ?", cond.DepartureTimeEndAt)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.Status) && len(cond.Status) > 0 {
			db = db.Where("status IN (?)", cond.Status)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.CanSell) && cond.CanSell {
			db = db.Where("sellable_seats > 0")
			db = db.Where("departure_time < ?", time.Now().Add(-time.Hour))
		}

		if pager == nil {
			return db.Model(dao.model)
		} else {
			return db.Model(dao.model).Scopes(dao.in.common.ParsePaging(pager))
		}
	}
}

func (dao *flightDao) UpdateSellableSeats(ctx context.Context, id int64, sellableSeats int) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Model(&dao.model).Where("id = ?", id).Update("sellable_seats", sellableSeats).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}
