package bo

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type CreatePaymentCond struct {
	BookingID int64
	Amount    decimal.Decimal
}

type Payment struct {
	ID              int64                 `json:"id"`
	BookingID       int64                 `json:"bookingID"`
	UserID          int64                 `json:"userID"`
	PaymentProvider *enum.PaymentProvider `json:"paymentProvider"`
	PaymentMethod   *enum.PaymentMethod   `json:"paymentMethod"`
	PaymentStatus   enum.PaymentStatus    `json:"paymentStatus"`
	Amount          decimal.Decimal       `json:"amount"`
	Currency        string                `json:"currency"`
	TransactionID   *string               `json:"transactionID"`
	PaymentURL      string                `json:"paymentURL"`
	ExpiredAt       time.Time             `json:"expiredAt"`
	PaidAt          *time.Time            `json:"paidAt"`
}

type ConfirmPaymentCond struct {
	ID              int64                 `json:"id"`
	TransactionID   *string               `json:"transactionID"`
	PaymentProvider *enum.PaymentProvider `json:"paymentProvider"`
	PaymentMethod   *enum.PaymentMethod   `json:"paymentMethod"`
	PaymentStatus   enum.PaymentStatus    `json:"paymentStatus"`
	PaidAt          *time.Time            `json:"paidAt"`
}
