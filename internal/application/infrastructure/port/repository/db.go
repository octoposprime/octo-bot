package application

import (
	"context"

	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {
	// SetLogger sets logging call-back function
	SetLogger(LogFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error))

	// GetOctoBotsByFilter returns the octoBots that match the given filter.
	GetOctoBotsByFilter(ctx context.Context, octoBotFilter me.OctoBotFilter) (me.OctoBots, error)

	// SaveOctoBot insert a new octoBot or update the existing one in the database.
	SaveOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error)

	// DeleteOctoBot soft-deletes the given octoBot in the database.
	DeleteOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error)
}
