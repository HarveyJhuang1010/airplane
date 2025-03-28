package interfaces

import "context"

type Listener interface {
	// Listen to the channel and process the message
	Listen(ctx context.Context)
	Name() string
}
