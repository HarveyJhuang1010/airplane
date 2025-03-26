package interfaces

import (
	"airplane/internal/components/errortool"
	"airplane/internal/components/launcher"
	"airplane/internal/domain/entities/bo"
	"context"
)

type EmailClient interface {
	SendEmail(ctx context.Context, email bo.Email) errortool.Error
	launcher.IService
}
