package listener

import (
	"airplane/internal/constant"
	"airplane/internal/domain/interfaces"
	"context"
)

type confirmBookingListener struct {
	in digIn
}

func newConfirmBookingListener(in digIn) interfaces.Listener {
	return &confirmBookingListener{
		in: in,
	}
}

func (l *confirmBookingListener) Listen(ctx context.Context) {
	l.in.Kafka.Consume(
		ctx,
		constant.KafkaGroupConfirmBooking,
		constant.KafkaTopicConfirmBooking,
		l.in.Booking.ConfirmBooking.ConfirmBooking)
}

func (l *confirmBookingListener) Name() string {
	return "confirm_booking_listener"
}
