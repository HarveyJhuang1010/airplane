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
	Provider *enum.PaymentProvider `json:"provider" enums:"stripe,paypal,line_pay,apple_pay,google_pay" swaggertype:"string"`
	// Payment Method
	Method *enum.PaymentMethod `json:"method" enums:"credit_card,debit_card,bank_transfer" swaggertype:"string"`
	// Payment Status
	Status enum.PaymentStatus `json:"status" enums:"pending,success,failed,cancelled" swaggertype:"string"`
	// Paid time
	PaidAt *time.Time `json:"paidAt" time_format:"2006-01-02T15:04:05Z07:00"`
}
