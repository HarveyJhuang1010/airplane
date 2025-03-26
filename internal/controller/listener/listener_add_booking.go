package listener

import (
	"airplane/internal/constant"
	"airplane/internal/domain/interfaces"
	"context"
)

type addBookingListener struct {
	in digIn
}

func newAddBookingListener(in digIn) interfaces.Listener {
	return &addBookingListener{
		in: in,
	}
}

func (l *addBookingListener) Listen(ctx context.Context) {
	l.in.Kafka.Consume(
		ctx,
		constant.KafkaGroupAddBooking,
		constant.KafkaTopicAddBooking,
		l.in.Booking.AddBooking.HandleBooking)
}

func (l *addBookingListener) Name() string {
	return "add_booking_listener"
}
