package infrastructure

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	map_repo "github.com/octoposprime/octo-bot/pkg/infrastructure/mapper/repository"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
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

	err := dbClient.DbClient.AutoMigrate(&map_repo.OctoBot{})
	if err != nil {
		panic(err)
	}

	return adapter
}

// SetLogger sets logging call-back function
func (a *DbAdapter) SetLogger(LoggerFunc func(ctx context.Context, logData *pb_logging.LogData) (*pb_logging.LoggingResult, error)) {
	a.Log = LoggerFunc
}

// GetOctoBotsByFilter returns the octoBots that match the given filter.
func (a DbAdapter) GetOctoBotsByFilter(ctx context.Context, octoBotFilter me.OctoBotFilter) (me.OctoBots, error) {
	var octoBotsDbMapper map_repo.OctoBots
	var filter map_repo.OctoBot
	qry := a.DbClient
	if octoBotFilter.Id.String() != "" && octoBotFilter.Id != (uuid.UUID{}) {
		filter.ID = octoBotFilter.Id
	}
	if octoBotFilter.OctoBotName != "" {
		filter.OctoBotName = octoBotFilter.OctoBotName
	}
	if octoBotFilter.Email != "" {
		filter.Email = octoBotFilter.Email
	}
	if octoBotFilter.OctoBotType != 0 {
		filter.OctoBotType = int(octoBotFilter.OctoBotType)
	}
	if octoBotFilter.OctoBotStatus != 0 {
		filter.OctoBotStatus = int(octoBotFilter.OctoBotStatus)
	}
	if len(octoBotFilter.Tags) > 0 {
		filter.Tags = octoBotFilter.Tags
	}
	if octoBotFilter.FirstName != "" {
		filter.FirstName = octoBotFilter.FirstName
	}
	if octoBotFilter.LastName != "" {
		filter.LastName = octoBotFilter.LastName
	}
	if !octoBotFilter.CreatedAtFrom.IsZero() && !octoBotFilter.CreatedAtTo.IsZero() {
		qry = qry.Where("created_at between ? and ?", octoBotFilter.CreatedAtFrom, octoBotFilter.CreatedAtTo)
	}
	if !octoBotFilter.UpdatedAtFrom.IsZero() && !octoBotFilter.UpdatedAtTo.IsZero() {
		qry = qry.Where("updated_at between ? and ?", octoBotFilter.UpdatedAtFrom, octoBotFilter.UpdatedAtTo)
	}
	if octoBotFilter.SearchText != "" {
		qry = qry.Where(
			qry.Where("UPPER(octoBot_name) LIKE UPPER(?)", "%"+octoBotFilter.SearchText+"%").
				Or("UPPER(email) LIKE UPPER(?)", "%"+octoBotFilter.SearchText+"%").
				Or("UPPER(array_to_string(tags, ',')) LIKE UPPER(?)", "%"+octoBotFilter.SearchText+"%"),
		)
	}
	qry = qry.Where(filter)
	var totalRows int64
	result := qry.Model(&map_repo.OctoBot{}).Where(filter).Count(&totalRows)
	if result.Error != nil {
		octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetOctoBotsByFilter", octoBotId, result.Error.Error()))
		totalRows = 0
	}
	if octoBotFilter.Limit != 0 {
		qry = qry.Limit(octoBotFilter.Limit)
	}
	if octoBotFilter.Offset != 0 {
		qry = qry.Offset(octoBotFilter.Offset)
	}
	if octoBotFilter.SortType != "" && octoBotFilter.SortField != 0 {
		sortStr := map_repo.OctoBotSortMap[octoBotFilter.SortField]
		if octoBotFilter.SortType == "desc" || octoBotFilter.SortType == "DESC" {
			sortStr += " desc"
		} else {
			sortStr += " asc"
		}
		qry = qry.Order(sortStr)
	}
	result = qry.Find(&octoBotsDbMapper)
	if result.Error != nil {
		octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "GetOctoBotsByFilter", octoBotId, result.Error.Error()))
		return me.OctoBots{}, result.Error
	}
	return me.OctoBots{
		OctoBots:  octoBotsDbMapper.ToEntities(),
		TotalRows: totalRows,
	}, nil
}

// SaveOctoBot insert a new octoBot or update the existing one in the database.
func (a DbAdapter) SaveOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	octoBotDbMapper := map_repo.NewOctoBotFromEntity(octoBot)
	qry := a.DbClient
	if octoBot.Id.String() != "" && octoBot.Id != (uuid.UUID{}) {
		qry = qry.Omit("created_at")
	}
	octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	if octoBotDbMapper.ID != (uuid.UUID{}) {
		octoBotDbMapper.UpdatedBy, _ = uuid.Parse(octoBotId)
	} else {
		octoBotDbMapper.CreatedBy, _ = uuid.Parse(octoBotId)
	}
	result := qry.Save(&octoBotDbMapper)
	if result.Error != nil {
		octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "SaveOctoBot", octoBotId, result.Error.Error()))
		return me.OctoBot{}, result.Error
	}
	return *octoBotDbMapper.ToEntity(), nil
}

// DeleteOctoBot soft-deletes the given octoBot in the database.
func (a DbAdapter) DeleteOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	octoBotDbMapper := map_repo.NewOctoBotFromEntity(octoBot)
	octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
	octoBotDbMapper.DeletedBy, _ = uuid.Parse(octoBotId)
	result := a.DbClient.Delete(&octoBotDbMapper)
	if result.Error != nil {
		octoBotId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteOctoBot", octoBotId, result.Error.Error()))
		return me.OctoBot{}, result.Error
	}
	return *octoBotDbMapper.ToEntity(), nil
}
