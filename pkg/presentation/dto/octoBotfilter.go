package presentation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	mo "github.com/octoposprime/octo-bot/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/octoBot"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// OctoBotFilter is a struct that represents the filter dto of a octoBot.
type OctoBotFilter struct {
	proto *pb.OctoBotFilter
}

// NewOctoBotFilter creates a new *OctoBotFilter.
func NewOctoBotFilter(pb *pb.OctoBotFilter) *OctoBotFilter {
	return &OctoBotFilter{
		proto: pb,
	}
}

// String returns a string representation of the OctoBotFilter.
func (s *OctoBotFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
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
		s.proto.Id,
		s.proto.OctoBotType,
		s.proto.OctoBotStatus,
		s.proto.Tags,
		s.proto.CreatedAtFrom,
		s.proto.CreatedAtTo,
		s.proto.UpdatedAtFrom,
		s.proto.UpdatedAtTo,
		s.proto.SearchText,
		s.proto.SortType,
		s.proto.SortField,
		s.proto.Limit,
		s.proto.Offset)
}

// NewOctoBotFilterFromEntity creates a new *OctoBotFilter from entity.
func NewOctoBotFilterFromEntity(entity me.OctoBotFilter) *OctoBotFilter {
	id := entity.Id.String()
	octoBotType := pb.OctoBotType(entity.OctoBotType)
	octoBotStatus := pb.OctoBotStatus(entity.OctoBotStatus)
	tags := entity.Tags
	createdAtFrom := timestamppb.New(entity.CreatedAtFrom)
	createdAtTo := timestamppb.New(entity.CreatedAtTo)
	updatedAtFrom := timestamppb.New(entity.UpdatedAtFrom)
	updatedAtTo := timestamppb.New(entity.UpdatedAtTo)
	searchText := entity.SearchText
	sortType := entity.SortType
	sortField := pb.OctoBotSortField(entity.SortField)
	limit := int32(entity.Limit)
	offset := int32(entity.Offset)
	return &OctoBotFilter{
		&pb.OctoBotFilter{
			Id:            &id,
			OctoBotType:   &octoBotType,
			OctoBotStatus: &octoBotStatus,
			Tags:          tags,
			CreatedAtFrom: createdAtFrom,
			CreatedAtTo:   createdAtTo,
			UpdatedAtFrom: updatedAtFrom,
			UpdatedAtTo:   updatedAtTo,
			SearchText:    &searchText,
			SortType:      &sortType,
			SortField:     &sortField,
			Limit:         &limit,
			Offset:        &offset,
		},
	}
}

// ToEntity returns a entity representation of the OctoBotFilter.
func (s *OctoBotFilter) ToEntity() *me.OctoBotFilter {
	id := uuid.UUID{}
	if s.proto.Id != nil {
		id = tuuid.FromString(*s.proto.Id)
	}
	octoBotType := 0
	if s.proto.OctoBotType != nil {
		octoBotType = int(*s.proto.OctoBotType)
	}
	octoBotStatus := 0
	if s.proto.OctoBotStatus != nil {
		octoBotStatus = int(*s.proto.OctoBotStatus)
	}
	tags := []string{}
	if s.proto.Tags != nil {
		tags = s.proto.Tags
	}
	createdAtFrom := time.Time{}
	if s.proto.CreatedAtFrom != nil {
		createdAtFrom = s.proto.CreatedAtFrom.AsTime()
	}
	createdAtTo := time.Time{}
	if s.proto.CreatedAtTo != nil {
		createdAtTo = s.proto.CreatedAtTo.AsTime()
	}
	updatedAtFrom := time.Time{}
	if s.proto.UpdatedAtFrom != nil {
		updatedAtFrom = s.proto.UpdatedAtFrom.AsTime()
	}
	updatedAtTo := time.Time{}
	if s.proto.UpdatedAtTo != nil {
		updatedAtTo = s.proto.UpdatedAtTo.AsTime()
	}
	searchText := ""
	if s.proto.SearchText != nil {
		searchText = string(*s.proto.SearchText)
	}
	sortType := ""
	if s.proto.SortType != nil {
		sortType = string(*s.proto.SortType)
	}
	sortField := 0
	if s.proto.SortField != nil {
		sortField = int(*s.proto.SortField)
	}
	limit := 0
	if s.proto.Limit != nil {
		limit = int(*s.proto.Limit)
	}
	offset := 0
	if s.proto.Offset != nil {
		offset = int(*s.proto.Offset)
	}
	return &me.OctoBotFilter{
		Id:            id,
		OctoBotType:   mo.OctoBotType(octoBotType),
		OctoBotStatus: mo.OctoBotStatus(octoBotStatus),
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     mo.OctoBotSortField(sortField),
		Limit:         limit,
		Offset:        offset,
	}
}
