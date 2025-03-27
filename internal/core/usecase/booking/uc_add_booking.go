package booking

import (
	"airplane/internal/constant"
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/domain/entities/po"
	"airplane/internal/enum"
	"airplane/internal/errs"
	"airplane/internal/tools/timelogger"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

func newAddBooking(in digIn) *AddBooking {
	return &AddBooking{
		in: in,
	}
}

type AddBooking struct {
	in digIn
}

func (uc *AddBooking) AddBooking(ctx context.Context, cond *bo.AddBookingCond) error {
	defer timelogger.LogTime(ctx)()

	// decrease the seat
	exist, err := uc.in.RedisRepository.FlightMaxSellCount.Exist(ctx, cond.FlightID)
	if err != nil || !exist {
		if err := uc.setFlightRemainSeat(ctx, cond.FlightID); err != nil {
			return err
		}
	}

	val, err := uc.in.RedisRepository.FlightMaxSellCount.DecrCount(ctx, cond.FlightID, time.Minute)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return err
	}
	if val < 0 {
		return errs.ErrFlightSoldOut
	}

	// push to mq
	cond.ID = uc.in.Snowflake.Generate().Int64()
	msgVal, err := json.Marshal(cond)
	if err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return errs.ErrParseFailed.TraceWrap(err)
	}
	if err := uc.in.Kafka.Produce(ctx, constant.KafkaTopicAddBooking, msgVal); err != nil {
		return errs.ErrMQFailed.TraceWrap(err)
	}

	return nil
}

func (uc *AddBooking) setFlightRemainSeat(ctx context.Context, flightID int64) error {
	defer timelogger.LogTime(ctx)()

	// get the flight
	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {
		flight, err := tx.FlightDAO().Get(ctx, flightID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		if err := uc.in.RedisRepository.FlightMaxSellCount.SetCount(
			ctx,
			flightID,
			int64(flight.SellableSeats),
			time.Minute,
		); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}
		return nil
	})
}

func (uc *AddBooking) HandleBooking(ctx context.Context, data []byte) error {
	defer timelogger.LogTime(ctx)()

	// unmarshal
	cond := &bo.AddBookingCond{}

	if err := json.Unmarshal(data, cond); err != nil {
		uc.in.Logger.AppLogger.Error(ctx, err)
		return errs.ErrParseFailed.TraceWrap(err)
	}

	// create booking
	booking := &po.Booking{
		ID:           cond.ID,
		FlightID:     cond.FlightID,
		CabinClassID: cond.CabinClassID,
		SeatID:       cond.SeatID,
		Status:       enum.BookingStatusPending,
		ExpiredAt:    time.Now().Add(time.Hour * 24),
	}

	return uc.in.DBRepository.Master().Transaction(func(tx *rdb.Database) error {

		// update the sellable seat
		flight, err := tx.FlightDAO().Get(ctx, cond.FlightID, false)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrRecordNotFound
			}
			return errs.ErrDBQueryFailed
		}
		if flight.DepartureTime.Before(time.Now().Add(-time.Hour)) || flight.SellableSeats <= 0 {
			return errs.ErrFlightNotAvailable
		}
		flight.SellableSeats--

		if err := tx.FlightDAO().UpdateSellableSeats(ctx, flight.ID, flight.SellableSeats); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		// get price
		class, err := tx.CabinClassDAO().Get(ctx, cond.CabinClassID, true)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			if errors.Is(err, errs.ErrRecordNotFound) {
				return errs.ErrDBQueryFailed
			}
			return err
		}
		booking.Price = class.Price

		// check seat
		if cond.SeatID != nil {
			seat, err := tx.SeatDAO().Get(ctx, *cond.SeatID, false)
			if err != nil {
				uc.in.Logger.AppLogger.Error(ctx, err)
				if errors.Is(err, errs.ErrRecordNotFound) {
					return errs.ErrDBQueryFailed
				}
				return err
			}
			if seat.Status != enum.SeatStatusAvailable {
				return errs.ErrSeatNotAvailable
			}
			if seat.CabinClassID != class.ID {
				return errs.ErrInvalidPayload
			}
		}

		// handle user
		userID, err := uc.in.User.GetUser.GetUser(ctx, cond.Email, cond.CountryCode, cond.PhoneNumber)
		if err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}
		booking.UserID = userID

		// create booking
		if err := tx.BookingDAO().Create(ctx, booking); err != nil {
			return err
		}

		// handle payment
		if _, err := uc.in.Payment.CreatePayment.CreatePayment(ctx, tx, &bo.CreatePaymentCond{
			BookingID: booking.ID,
			UserID:    userID,
			Amount:    booking.Price,
		}); err != nil {
			uc.in.Logger.AppLogger.Error(ctx, err)
			return err
		}

		return nil
	})
}
