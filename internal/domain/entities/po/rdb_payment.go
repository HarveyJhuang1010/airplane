package po

import (
	"airplane/internal/enum"
	"github.com/shopspring/decimal"
	"time"
)

type Payment struct {
	ID            int64                 `gorm:"primaryKey" json:"id"`
	BookingID     int64                 `gorm:"uniqueIndex;not null" json:"bookingID"`
	UserID        int64                 `gorm:"not null;index" json:"userID"`
	Provider      *enum.PaymentProvider `gorm:"type:varchar(50);default null" json:"provider"`
	Method        *enum.PaymentMethod   `gorm:"type:varchar(50);default null" json:"method"`
	Status        enum.PaymentStatus    `gorm:"type:varchar(20);not null" json:"status"`
	Amount        decimal.Decimal       `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency      string                `gorm:"type:varchar(10);default:'TWD'" json:"currency"`
	TransactionID *string               `gorm:"type:varchar(100)" json:"transactionID"`
	PaymentURL    string                `gorm:"type:text" json:"paymentURL"`
	ExpiredAt     time.Time             `gorm:"not null" json:"expiredAt"`
	PaidAt        *time.Time            `json:"paid_at,omitempty" json:"paidAt"`

	Booking Booking `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`

	At
}

func (Payment) TableName() string {
	return "payment"
}

type ExtraPayment struct {
	ID            int64                 `gorm:"primaryKey" json:"id"`
	BookingID     int64                 `gorm:"uniqueIndex;not null" json:"bookingID"`
	UserID        int64                 `gorm:"not null;index" json:"userID"`
	Provider      *enum.PaymentProvider `gorm:"type:varchar(50);not null" json:"provider"`
	Method        *enum.PaymentMethod   `gorm:"type:varchar(50);not null" json:"method"`
	Status        enum.PaymentStatus    `gorm:"type:varchar(20);not null" json:"status"`
	Amount        decimal.Decimal       `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency      string                `gorm:"type:varchar(10);default:'TWD'" json:"currency"`
	TransactionID *string               `gorm:"type:varchar(100)" json:"transactionID"`
	PaymentURL    string                `gorm:"type:text" json:"paymentURL"`
	ExpiredAt     time.Time             `gorm:"not null" json:"expiredAt"`
	PaidAt        *time.Time            `json:"paid_at,omitempty" json:"paidAt"`

	Booking Booking `gorm:"foreignKey:BookingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User    User    `gorm:"foreignKey:UserID;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`

	At
}

func (ExtraPayment) TableName() string {
	return "extra_payment"
}

type PaymentUpdateResultCond struct {
	ID            int64
	TransactionID *string
	Provider      *enum.PaymentProvider `gorm:"type:varchar(50);default null" json:"provider"`
	Method        *enum.PaymentMethod   `gorm:"type:varchar(50);default null" json:"method"`
	Status        enum.PaymentStatus
	PaidAt        *time.Time
}
