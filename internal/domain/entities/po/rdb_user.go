package po

import (
	"airplane/internal/enum"
)

// User represents the `user` table
type User struct {
	ID               int64           `gorm:"column:id" json:"id"`
	Email            string          `gorm:"column:email;uniqueIndex:idx_email_phone" json:"email"`
	PhoneCountryCode string          `gorm:"column:phone_country_code;uniqueIndex:idx_email_phone" json:"phoneCountryCode"`
	PhoneNumber      string          `gorm:"column:phone_number;uniqueIndex:idx_email_phone" json:"phoneNumber"`
	Status           enum.UserStatus `gorm:"column:status" json:"status"`
	SecretKey        string          `gorm:"column:secret_key" json:"secretKey"`
	At
}

func (User) TableName() string {
	return "user"
}

type UserListCond struct {
	*Pager
	Email            *string
	PhoneCountryCode *string
	PhoneNumber      *string
	Status           []enum.UserStatus
}
