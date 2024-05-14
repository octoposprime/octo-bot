package presentation

import (
	"context"

	dto "github.com/octoposprime/octo-bot/pkg/presentation/dto"
	pb_octoBot "github.com/octoposprime/op-be-shared/pkg/proto/pb/octoBot"
)

// GetOctoBotsByFilter returns the octoBots that match the given filter.
func (a *Grpc) GetOctoBotsByFilter(ctx context.Context, filter *pb_octoBot.OctoBotFilter) (*pb_octoBot.OctoBots, error) {
	octoBots, err := a.queryHandler.GetOctoBotsByFilter(ctx, *dto.NewOctoBotFilter(filter).ToEntity())
	return dto.NewOctoBotFromEntities(octoBots).ToPbs(), err
}

// CreateOctoBot sends the given octoBot to the application layer for creating new octoBot.
func (a *Grpc) CreateOctoBot(ctx context.Context, octoBot *pb_octoBot.OctoBot) (*pb_octoBot.OctoBot, error) {
	data, err := a.commandHandler.CreateOctoBot(ctx, *dto.NewOctoBot(octoBot).ToEntity())
	return dto.NewOctoBotFromEntity(data).ToPb(), err
}

// UpdateOctoBotBase sends the given octoBot to the application layer for updating octoBot's base values.
func (a *Grpc) UpdateOctoBotBase(ctx context.Context, octoBot *pb_octoBot.OctoBot) (*pb_octoBot.OctoBot, error) {
	data, err := a.commandHandler.UpdateOctoBotBase(ctx, *dto.NewOctoBot(octoBot).ToEntity())
	return dto.NewOctoBotFromEntity(data).ToPb(), err
}

// UpdateOctoBotCore sends the given octoBot to the application layer for updating octoBot's core values.
func (a *Grpc) UpdateOctoBotCore(ctx context.Context, octoBot *pb_octoBot.OctoBot) (*pb_octoBot.OctoBot, error) {
	data, err := a.commandHandler.UpdateOctoBotCore(ctx, *dto.NewOctoBot(octoBot).ToEntity())
	return dto.NewOctoBotFromEntity(data).ToPb(), err
}

// UpdateOctoBotStatus sends the given octoBot to the application layer for updating octoBot status.
func (a *Grpc) UpdateOctoBotStatus(ctx context.Context, octoBot *pb_octoBot.OctoBot) (*pb_octoBot.OctoBot, error) {
	data, err := a.commandHandler.UpdateOctoBotStatus(ctx, *dto.NewOctoBot(octoBot).ToEntity())
	return dto.NewOctoBotFromEntity(data).ToPb(), err
}

// DeleteOctoBot sends the given octoBot to the application layer for deleting data.
func (a *Grpc) DeleteOctoBot(ctx context.Context, octoBot *pb_octoBot.OctoBot) (*pb_octoBot.OctoBot, error) {
	data, err := a.commandHandler.DeleteOctoBot(ctx, *dto.NewOctoBot(octoBot).ToEntity())
	return dto.NewOctoBotFromEntity(data).ToPb(), err
}
