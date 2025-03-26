package user

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func newGetUser(in dependence) *GetUser {
	return &GetUser{
		in: in,
	}
}

type GetUser struct {
	in dependence
}

func (uc *GetUser) GetUser(ctx context.Context, email, countryCode, phone string) (int64, error) {
	defer timelogger.LogTime(ctx)()

	if email == "" && countryCode == "" && phone == "" {
		return 0, errs.ErrInvalidPayload
	}

	var user *po.User
	users, err := uc.in.DBRepository.Master().UserDAO().List(ctx, &po.UserListCond{
		Email:            &email,
		PhoneCountryCode: &countryCode,
		PhoneNumber:      &phone,
	})
	if errors.Is(err, errs.ErrRecordNotFound) || len(users) == 0 {
		user = &po.User{
			ID:               uc.in.Snowflake.Generate().Int64(),
			Email:            email,
			PhoneCountryCode: countryCode,
			PhoneNumber:      phone,
			Status:           enum.UserStatusEnable,
			SecretKey:        uuid.New().String(),
		}
		if err := uc.in.DBRepository.Master().UserDAO().Create(ctx, user); err != nil {
			if errors.Is(err, errs.ErrDuplicateRecord) {
				users, err = uc.in.DBRepository.Master().UserDAO().List(ctx, &po.UserListCond{
					Email:            &email,
					PhoneCountryCode: &countryCode,
					PhoneNumber:      &phone,
				})
				if err != nil {
					return 0, err
				}
				if len(users) == 0 {
					return 0, errs.ErrRecordNotFound
				}
				user = users[0]
			} else {
				return 0, err
			}
		}
	} else if err != nil {
		return 0, err
	} else {
		user = users[0]
	}

	return user.ID, nil
}
