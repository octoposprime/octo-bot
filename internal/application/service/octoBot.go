package application

import (
	"context"

	"github.com/google/uuid"
	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	mo "github.com/octoposprime/octo-bot/internal/domain/model/object"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// GetOctoBotsByFilter returns the octoBots that match the given filter.
func (a *Service) GetOctoBotsByFilter(ctx context.Context, octoBotFilter me.OctoBotFilter) (me.OctoBots, error) {
	return a.DbPort.GetOctoBotsByFilter(ctx, octoBotFilter)
}

// CreateOctoBot sends the given octoBot to the repository of the infrastructure layer for creating a new octoBot.
func (a *Service) CreateOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	user.Id = uuid.UUID{}
	if err := a.ValidateUser(&user); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	if err := a.CheckUserNameRules(&user); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	if err := a.CheckEmailRules(&user); err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	var userEmailCheckFilter me.UserFilter
	userEmailCheckFilter.Email = user.Email
	emailExistsUsers, err := a.GetUsersByFilter(ctx, userEmailCheckFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	if emailExistsUsers.TotalRows > 0 {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		err := mo.ErrorUserEmailIsExists
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	var userNameCheckFilter me.UserFilter
	userNameCheckFilter.UserName = user.UserName
	nameExistsUsers, err := a.GetUsersByFilter(ctx, userNameCheckFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	if nameExistsUsers.TotalRows > 0 {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		err := mo.ErrorUserUsernameIsExists
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "CreateUser", userId, err.Error()))
		return me.User{}, err
	}
	if user.UserStatus == mo.UserStatusNONE {
		user.UserStatus = mo.UserStatusACTIVE
	}
	return a.DbPort.SaveUser(ctx, user)
}

// UpdateOctoBotBase sends the given base values of the octoBot to the repository of the infrastructure layer for updating base values of octoBot data.
func (a *Service) UpdateOctoBotBase(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.Tags = user.Tags
		dbUser.OctoBotType = user.OctoBotType
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdateOctoBotCore sends the given core values of the octoBot to the repository of the infrastructure layer for updating core values of octoBot data.
func (a *Service) UpdateOctoBotCore(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.OctoBotData = user.OctoBotData
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// UpdateOctoBotStatus sends the given status value of the octoBot to the repository of the infrastructure layer for updating status of octoBot data.
func (a *Service) UpdateOctoBotStatus(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	if user.Id.String() == "" || user.Id == (uuid.UUID{}) {
		err := mo.ErrorUserIdIsEmpty
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	var userFilter me.UserFilter
	userFilter.Id = user.Id
	users, err := a.GetUsersByFilter(ctx, userFilter)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
		return me.User{}, err
	}
	if users.TotalRows > 0 {
		dbUser := users.Users[0]
		dbUser.UserStatus = user.UserStatus
		if err := a.ValidateUser(&dbUser); err != nil {
			userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
			go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "UpdateUserStatus", userId, err.Error()))
			return me.User{}, err
		}
		return a.DbPort.SaveUser(ctx, dbUser)
	} else {
		return user, mo.ErrorUserNotFound
	}
}

// DeleteOctoBot sends the given octoBot to the repository of the infrastructure layer for deleting data.
func (a *Service) DeleteOctoBot(ctx context.Context, octoBot me.OctoBot) (me.OctoBot, error) {
	var err error
	user, err = a.DbPort.DeleteUser(ctx, user)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUser", userId, err.Error()))
		return me.User{}, err
	}

	err = a.RedisPort.DeleteUserPasswordByUserId(ctx, user.Id)
	if err != nil {
		userId, _ := ctx.Value(smodel.QueryKeyUid).(string)
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "DeleteUser", userId, err.Error()))
		return me.User{}, err
	}
	return user, err
}
