package presentation

import (
	"fmt"

	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	mo "github.com/octoposprime/octo-bot/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/octoBot"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// OctoBot is a struct that represents the dto of a octoBot basic values.
type OctoBot struct {
	proto *pb.OctoBot
}

// NewOctoBot creates a new *OctoBot.
func NewOctoBot(pb *pb.OctoBot) *OctoBot {
	return &OctoBot{
		proto: pb,
	}
}

// String returns a string representation of the OctoBot.
func (s *OctoBot) String() string {
	return fmt.Sprintf("Id: %v, "+
		"OctoBotData: %v, "+
		"OctoBotType: %v, "+
		"OctoBotStatus: %v, "+
		"Tags: %v",
		s.proto.Id,
		s.proto.OctoBotData,
		s.proto.OctoBotType,
		s.proto.OctoBotStatus,
		s.proto.Tags)
}

// NewOctoBotFromEntity creates a new *OctoBot from entity.
func NewOctoBotFromEntity(entity me.OctoBot) *OctoBot {
	return &OctoBot{
		&pb.OctoBot{
			Id:            entity.Id.String(),
			OctoBotData:   entity.OctoBotData,
			OctoBotType:   pb.OctoBotType(entity.OctoBotType),
			OctoBotStatus: pb.OctoBotStatus(entity.OctoBotStatus),
			Tags:          entity.Tags,

			// Only for view
			CreatedAt: timestamppb.New(entity.CreatedAt),
			UpdatedAt: timestamppb.New(entity.UpdatedAt),
		},
	}
}

// ToPb returns a protobuf representation of the OctoBot.
func (s *OctoBot) ToPb() *pb.OctoBot {
	return s.proto
}

// ToEntity returns a entity representation of the OctoBot.
func (s *OctoBot) ToEntity() *me.OctoBot {
	return &me.OctoBot{
		Id: tuuid.FromString(s.proto.Id),
		OctoBot: mo.OctoBot{
			OctoBotData:   s.proto.OctoBotData,
			OctoBotType:   mo.OctoBotType(s.proto.OctoBotType),
			OctoBotStatus: mo.OctoBotStatus(s.proto.OctoBotStatus),
			Tags:          s.proto.Tags,
		},
	}
}

type OctoBots struct {
	OctoBots  []*OctoBot `json:"octoBots"`
	TotalRows int64      `json:"total_rows"`
}

// NewOctoBotsFromEntities creates a new []*OctoBot from entities.
func NewOctoBotFromEntities(entities me.OctoBots) OctoBots {
	octoBots := make([]*OctoBot, len(entities.OctoBots))
	for i, entity := range entities.OctoBots {
		octoBots[i] = NewOctoBotFromEntity(entity)
	}

	return OctoBots{
		OctoBots:  octoBots,
		TotalRows: entities.TotalRows,
	}
}

// ToPbs returns a protobuf representation of the OctoBots.
func (s OctoBots) ToPbs() *pb.OctoBots {
	octoBots := make([]*pb.OctoBot, len(s.OctoBots))
	for i, octoBot := range s.OctoBots {
		octoBots[i] = octoBot.proto
	}
	return &pb.OctoBots{
		OctoBots:  octoBots,
		TotalRows: s.TotalRows,
	}
}
