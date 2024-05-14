package application

import (
	ip_repo "github.com/octoposprime/octo-bot/internal/application/infrastructure/port/repository"
	ds "github.com/octoposprime/octo-bot/internal/domain/service"
)

// Service is an application service.
// It manages the business logic of the application.
type Service struct {
	*ds.Service
	ip_repo.DbPort
}

// NewService creates a new *Service.
func NewService(domainService *ds.Service, dbRepository ip_repo.DbPort) *Service {
	service := &Service{
		domainService,
		dbRepository,
	}
	service.Migrate()
	return service
}
