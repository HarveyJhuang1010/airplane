package interfaces

import "context"

type Service interface {
	Run(ctx context.Context, stop context.CancelFunc)
	Shutdown(ctx context.Context)
}
