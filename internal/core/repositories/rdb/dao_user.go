package rdb

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func (d *Database) UserDAO() IUserDao {
	return newUserDao(d.Session)
}

func newUserDao(session *Session) *UserDao {
	return &UserDao{
		Session: session,
	}
}

type IUserDao interface {
	Create(ctx context.Context, cond ...*po.User) error
	Get(ctx context.Context, id int64, forUpdate bool) (*po.User, error)
	List(ctx context.Context, cond *po.UserListCond) ([]*po.User, error)
	ListPager(ctx context.Context, cond *po.UserListCond) (*po.Pagination, error)
}

type UserDao struct {
	*Session
	model po.User
}

// Create 新增操作紀錄
func (dao *UserDao) Create(ctx context.Context, cond ...*po.User) error {
	defer timelogger.LogTime(ctx)()

	if err := dao.db.Create(&cond).Error; err != nil {
		return dao.Session.in.common.ErrorHandle(err)
	}

	return nil
}

func (dao *UserDao) Get(ctx context.Context, id int64, forUpdate bool) (*po.User, error) {
	defer timelogger.LogTime(ctx)()

	var model po.User

	query := dao.db.Where(`id = ?`, id)
	if forUpdate {
		query = query.Scopes(dao.in.common.SetForUpdate())
	}

	if err := query.First(&model, id).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return &model, nil
}

func (dao *UserDao) List(ctx context.Context, cond *po.UserListCond) ([]*po.User, error) {
	defer timelogger.LogTime(ctx)()

	var result []*po.User

	query := dao.db.Scopes(dao.list(cond, cond.Pager))
	if err := query.Find(&result).Error; err != nil {
		return nil, dao.Session.in.common.ErrorHandle(err)
	}

	return result, nil
}

// ListPager get total count by pager
func (dao *UserDao) ListPager(ctx context.Context, cond *po.UserListCond) (*po.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	var count int64

	if err := dao.db.
		Scopes(dao.list(cond, cond.Pager)).Count(&count).
		Error; err != nil {
		return nil, err
	}

	return po.NewPagination(cond.Pager, count), nil
}

func (dao *UserDao) list(cond *po.UserListCond, pager *po.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !lo.IsNil(cond) && !lo.IsNil(cond.Email) && !lo.IsEmpty(cond.Email) {
			db = db.Where("email = ?", cond.Email)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.PhoneCountryCode) && !lo.IsEmpty(cond.PhoneCountryCode) {
			db = db.Where("phone_country_code = ?", cond.PhoneCountryCode)
		}
		if !lo.IsNil(cond) && !lo.IsNil(cond.PhoneNumber) && !lo.IsEmpty(cond.PhoneNumber) {
			db = db.Where("phone_number = ?", cond.PhoneNumber)
		}

		if !lo.IsNil(cond) && !lo.IsNil(cond.Status) && len(cond.Status) > 0 {
			db = db.Where("status = ?", cond.Status)
		}

		if pager == nil {
			return db.Model(dao.model)
		} else {
			return db.Model(dao.model).Scopes(dao.in.common.ParsePaging(pager))
		}
	}
}
