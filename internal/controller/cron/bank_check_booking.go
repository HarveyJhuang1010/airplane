package cron

import (
	"airplane/internal/components/logger"
	"airplane/internal/domain/interfaces"
	"context"
)

type checkBookingTask struct {
	in digIn
}

var _ interfaces.CronTask = (*checkBookingTask)(nil)

func newCheckBookingTask(in digIn) interfaces.CronTask {
	return &checkBookingTask{
		in: in,
	}
}

func (t *checkBookingTask) Name() string {
	return "check_booking"
}

func (t *checkBookingTask) Schedule() string {
	return "0 */5 * * * *" // 5 minutes for test
}

func (t *checkBookingTask) Run() {
	ctx := context.Background()
	if err := t.in.Booking.CheckBooking.CheckBooking(ctx); err != nil {
		t.in.Logger.AppLogger.Error(ctx, err, logger.String("task", t.Name()))
	}
}
