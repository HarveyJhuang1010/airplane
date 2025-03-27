package bo

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type CreatePaymentCond struct {
	BookingID int64
	UserID    int64
	Amount    decimal.Decimal
}

type Payment struct {
	ID              int64                 `json:"id"`
	BookingID       int64                 `json:"bookingID"`
	UserID          int64                 `json:"userID"`
	PaymentProvider *enum.PaymentProvider `json:"provider"`
	PaymentMethod   *enum.PaymentMethod   `json:"method"`
	PaymentStatus   enum.PaymentStatus    `json:"status"`
	Amount          decimal.Decimal       `json:"amount"`
	Currency        string                `json:"currency"`
	TransactionID   *string               `json:"transactionID"`
	PaymentURL      string                `json:"paymentURL"`
	ExpiredAt       time.Time             `json:"expiredAt"`
	PaidAt          *time.Time            `json:"paidAt"`
}

type ConfirmPaymentCond struct {
	ID            int64                 `json:"id"`
	TransactionID *string               `json:"transactionID"`
	Provider      *enum.PaymentProvider `json:"provider"`
	Method        *enum.PaymentMethod   `json:"method"`
	Status        enum.PaymentStatus    `json:"status"`
	PaidAt        *time.Time            `json:"paidAt"`
}
