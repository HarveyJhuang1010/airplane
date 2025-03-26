package bo

import "airplane/internal/enum"

type User struct {
	ID               int64           `json:"id"`
	Email            string          `json:"email"`
	PhoneCountryCode string          `json:"phoneCountryCode"`
	PhoneNumber      string          `json:"phoneNumber"`
	Status           enum.UserStatus `json:"status"`
}
