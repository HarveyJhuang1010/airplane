package po

import (
	"gorm.io/gorm"
	"time"
)

type At struct {
	CreatedAt time.Time       `gorm:"column:created_at"`
	UpdatedAt time.Time       `gorm:"column:updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at"`
}

type OrderBy struct {
	Column string
	Desc   bool // true: DESC, false: ASC
}

type TimeRange struct {
	Start *time.Time
	End   *time.Time
}

func NewTimeRange(start, end *time.Time) *TimeRange {
	return &TimeRange{
		Start: start,
		End:   end,
	}
}
