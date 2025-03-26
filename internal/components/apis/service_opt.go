package apis

import (
	"airplane/internal/components/logger"
)

func ServiceOption() *serviceOption {
	return new(serviceOption)
}

type serviceOption func(s *Service)

func (serviceOption) WithLogger(logger logger.ILogger) func(e *Service) {
	return func(s *Service) {
		s.Logger = logger
	}
}

func (serviceOption) WithConfig(config ServiceConfig) func(e *Service) {
	return func(s *Service) {
		s.Config = config
	}
}
