package infrastructure

import (
	"context"

	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

type DbAdapter struct {
	*tgorm.GormClient
	Log func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)
}

func NewDbAdapter(dbClient *tgorm.GormClient) DbAdapter {
	adapter := DbAdapter{
		dbClient,
		Log,
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *DbAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}
