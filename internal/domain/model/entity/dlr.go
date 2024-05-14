package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/octo-bot/internal/domain/model/object"
)

// OctoBot is a struct that represents the entity of a octoBot basic values.
type OctoBot struct {
	Id         uuid.UUID `json:"id"` // Id is the id of the octoBot.
	mo.OctoBot           // OctoBot is the basic values of the octoBot.

	// Only for view
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the create time.
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the update time.
}

// NewOctoBot creates a new *OctoBot.
func NewOctoBot(id uuid.UUID,
	octoBot mo.OctoBot) *OctoBot {
	return &OctoBot{
		Id:      id,
		OctoBot: octoBot,
	}
}

// NewEmptyOctoBot creates a new *OctoBot with empty values.
func NewEmptyOctoBot() *OctoBot {
	return &OctoBot{
		Id:      uuid.UUID{},
		OctoBot: *mo.NewEmptyOctoBot(),
	}
}

// String returns a string representation of the OctoBot.
func (s *OctoBot) String() string {
	return fmt.Sprintf("Id: %v, "+
		"OctoBot: %v",
		s.Id,
		s.OctoBot)
}

// Equals returns true if the OctoBot is equal to the other OctoBot.
func (s *OctoBot) Equals(other *OctoBot) bool {
	if s.Id != other.Id {
		return false
	}
	if !s.OctoBot.Equals(&other.OctoBot) {
		return false
	}
	return true
}

// Clone returns a clone of the OctoBot.
func (s *OctoBot) Clone() *OctoBot {
	return &OctoBot{
		Id:      s.Id,
		OctoBot: *s.OctoBot.Clone(),
	}
}

// IsEmpty returns true if the OctoBot is empty.
func (s *OctoBot) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if !s.OctoBot.IsEmpty() {
		return false
	}
	return true
}

// IsNotEmpty returns true if the OctoBot is not empty.
func (s *OctoBot) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the OctoBot.
func (s *OctoBot) Clear() {
	s.Id = uuid.UUID{}
	s.OctoBot.Clear()
}

// Validate validates the OctoBot.
func (s *OctoBot) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorOctoBotIsEmpty
	}
	if err := s.OctoBot.Validate(); err != nil {
		return err
	}
	return nil
}

// OctoBots contains a slice of *OctoBot and total number of octoBots.
type OctoBots struct {
	OctoBots  []OctoBot `json:"octoBots"`   // OctoBots is the slice of *OctoBot.
	TotalRows int64     `json:"total_rows"` // TotalRows is the total number of rows.
}
