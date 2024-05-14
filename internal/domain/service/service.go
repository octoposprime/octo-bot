package domain

import (
	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
)

// This is the service layer of the domain layer.
type Service struct {
}

// NewService creates a new *Service.
func NewService() *Service {
	return &Service{}
}

// ValidateOctoBot validates the octoBot.
func (s *Service) ValidateOctoBot(octoBot *me.OctoBot) error {
	if err := octoBot.Validate(); err != nil {
		return err
	}
	return nil
}
