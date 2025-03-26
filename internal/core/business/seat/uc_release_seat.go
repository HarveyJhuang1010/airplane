package seat

import (
	"airplane/internal/core/repositories/rdb"
	"airplane/internal/tools/timelogger"
	"context"
)

func newReleaseSeat(in dependence) *ReleaseSeat {
	return &ReleaseSeat{
		in: in,
	}
}

type ReleaseSeat struct {
	in dependence
}

func (uc *ReleaseSeat) ReleaseSeat(ctx context.Context, tx *rdb.Database, seatID int64) error {
	defer timelogger.LogTime(ctx)()

	// Implement the business logic of ReleaseSeat here
	return nil
}
