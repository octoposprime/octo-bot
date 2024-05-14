package domain

import (
	"fmt"
)

// OctoBot is a struct that represents the object of a octoBot basic values.
type OctoBot struct {
	OctoBotData   string        `json:"octoBot_data"`   // OctoBotData is the data of the octoBot.
	OctoBotType   OctoBotType   `json:"octoBot_type"`   // OctoBotType is the type of the octoBot.
	OctoBotStatus OctoBotStatus `json:"octoBot_status"` // OctoBotStatus is the status of the octoBot.
	Tags          []string      `json:"tags"`           // Tags is the tags of the octoBot.
}

// NewOctoBot creates a new *OctoBot.
func NewOctoBot(octoBotData string,
	octoBotType OctoBotType,
	octoBotStatus OctoBotStatus,
	tags []string) *OctoBot {
	return &OctoBot{
		OctoBotData:   octoBotData,
		OctoBotType:   octoBotType,
		OctoBotStatus: octoBotStatus,
		Tags:          tags,
	}
}

// NewEmptyOctoBot creates a new *OctoBot with empty values.
func NewEmptyOctoBot() *OctoBot {
	return &OctoBot{
		OctoBotData:   "",
		OctoBotType:   OctoBotTypeNONE,
		OctoBotStatus: OctoBotStatusNONE,
		Tags:          []string{},
	}
}

// String returns a string representation of the OctoBot.
func (s *OctoBot) String() string {
	return fmt.Sprintf("OctoBotData: %v, "+
		"OctoBotType: %v, "+
		"OctoBotStatus: %v, "+
		"Tags: %v",
		s.OctoBotData,
		s.OctoBotType,
		s.OctoBotStatus,
		s.Tags)
}

// Equals returns true if the OctoBot is equal to the other OctoBot.
func (s *OctoBot) Equals(other *OctoBot) bool {
	if s.OctoBotData != other.OctoBotData {
		return false
	}
	if s.OctoBotType != other.OctoBotType {
		return false
	}
	if s.OctoBotStatus != other.OctoBotStatus {
		return false
	}
	for i := range s.Tags {
		if s.Tags[i] != other.Tags[i] {
			return false
		}
	}
	return true
}

// Clone returns a clone of the OctoBot.
func (s *OctoBot) Clone() *OctoBot {
	return &OctoBot{
		OctoBotData:   s.OctoBotData,
		OctoBotType:   s.OctoBotType,
		OctoBotStatus: s.OctoBotStatus,
		Tags:          s.Tags,
	}
}

// IsEmpty returns true if the OctoBot is empty.
func (s *OctoBot) IsEmpty() bool {
	if s.OctoBotData != "" {
		return false
	}
	if s.OctoBotType != OctoBotTypeNONE {
		return false
	}
	if s.OctoBotStatus != OctoBotStatusNONE {
		return false
	}
	if len(s.Tags) != 0 {
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
	s.OctoBotData = ""
	s.OctoBotType = OctoBotTypeNONE
	s.OctoBotStatus = OctoBotStatusNONE
	s.Tags = []string{}
}

// Validate validates the OctoBot.
func (s *OctoBot) Validate() error {
	if s.IsEmpty() {
		return ErrorOctoBotIsEmpty
	}
	if s.OctoBotData == "" {
		return ErrorOctoBotOctoBotDataIsEmpty
	}
	return nil
}
