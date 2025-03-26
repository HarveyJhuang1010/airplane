package payment

import (
	"airplane/internal/enum"
	"time"
)

type NotifyPaymentCond struct {
	// Payment ID
	ID int64 `json:"id,string"`
	// 3rd Party Transaction ID
	TransactionID *string `json:"transactionID"`
	// Payment Provider
	PaymentProvider *enum.PaymentProvider `json:"paymentProvider" enums:"stripe,paypal,line_pay,apple_pay,google_pay" swaggertype:"string"`
	// Payment Method
	PaymentMethod *enum.PaymentMethod `json:"paymentMethod" enums:"credit_card,debit_card,bank_transfer" swaggertype:"string"`
	// Payment Status
	PaymentStatus enum.PaymentStatus `json:"paymentStatus" enums:"pending,success,failed,cancelled" swaggertype:"string"`
	// Paid time
	PaidAt *time.Time `json:"paidAt" time_format:"2006-01-02T15:04:05Z07:00"`
}
