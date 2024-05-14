package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/octo-bot/internal/domain/model/object"
)

// OctoBotFilter is a struct that represents the filter of a octoBot.
type OctoBotFilter struct {
	Id            uuid.UUID        `json:"id"`             // Id is the id of the octoBot.
	OctoBotData   string           `json:"octoBot_name"`   // OctoBotData is the octoBot name of the octoBot.
	OctoBotType   mo.OctoBotType   `json:"octoBot_type"`   // OctoBotType is the type of the octoBot.
	OctoBotStatus mo.OctoBotStatus `json:"octoBot_status"` // OctoBotStatus is the status of the octoBot.
	Tags          []string         `json:"tags"`           // Tags is the tags of the octoBot.

	CreatedAtFrom time.Time `json:"created_at_from"` // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	CreatedAtTo   time.Time `json:"created_at_to"`   // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	UpdatedAtFrom time.Time `json:"updated_at_from"` // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.
	UpdatedAtTo   time.Time `json:"updated_at_to"`   // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.

	SearchText string              `json:"search_text"` // SearchText is the full-text search value.
	SortType   string              `json:"sort_type"`   // SortType is the sorting type (ASC,DESC).
	SortField  mo.OctoBotSortField `json:"sort_field"`  // SortField is the sorting field of the octoBot.

	Limit  int `json:"limit"`  // Limit provides to limitation row size.
	Offset int `json:"offset"` // Offset provides a starting row number of the limitation.
}

// NewOctoBotFilter creates a new *OctoBotFilter.
func NewOctoBotFilter(id uuid.UUID,
	octoBotData string,
	octoBotType mo.OctoBotType,
	octoBotStatus mo.OctoBotStatus,
	tags []string,
	createdAtFrom time.Time,
	createdAtTo time.Time,
	updatedAtFrom time.Time,
	updatedAtTo time.Time,
	searchText string,
	sortType string,
	sortField mo.OctoBotSortField,
	limit int,
	offset int) *OctoBotFilter {
	return &OctoBotFilter{
		Id:            id,
		OctoBotData:   octoBotData,
		OctoBotType:   octoBotType,
		OctoBotStatus: octoBotStatus,
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     sortField,
		Limit:         limit,
		Offset:        offset,
	}
}

// NewEmptyOctoBotFilter creates a new *OctoBotFilter with empty values.
func NewEmptyOctoBotFilter() *OctoBotFilter {
	return &OctoBotFilter{
		Id:            uuid.UUID{},
		OctoBotData:   "",
		OctoBotType:   mo.OctoBotTypeNONE,
		OctoBotStatus: mo.OctoBotStatusNONE,
		Tags:          []string{},
		CreatedAtFrom: time.Time{},
		CreatedAtTo:   time.Time{},
		UpdatedAtFrom: time.Time{},
		UpdatedAtTo:   time.Time{},
		SearchText:    "",
		SortType:      "",
		SortField:     mo.OctoBotSortFieldNONE,
		Limit:         0,
		Offset:        0,
	}
}

// String returns a string representation of the OctoBotFilter.
func (s *OctoBotFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
		"OctoBotData: %v, "+
		"OctoBotType: %v, "+
		"OctoBotStatus: %v, "+
		"Tags: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.Id,
		s.OctoBotData,
		s.OctoBotType,
		s.OctoBotStatus,
		s.Tags,
		s.CreatedAtFrom,
		s.CreatedAtTo,
		s.UpdatedAtFrom,
		s.UpdatedAtTo,
		s.SearchText,
		s.SortType,
		s.SortField,
		s.Limit,
		s.Offset)
}
