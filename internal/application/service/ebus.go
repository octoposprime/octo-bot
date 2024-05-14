package application

import (
	"context"

	me "github.com/octoposprime/octo-bot/internal/domain/model/entity"
	smodel "github.com/octoposprime/op-be-shared/pkg/model"
	pb_logging "github.com/octoposprime/op-be-shared/pkg/proto/pb/logging"
)

// This is the event listener handler of the application layer.
func (a *Service) EventListen() *Service {
	go a.Listen(context.Background(), smodel.ChannelCreateOctoBot, a.EventListenerCallBack)
	go a.Listen(context.Background(), smodel.ChannelDeleteOctoBot, a.EventListenerCallBack)
	return a
}

// This is a call-back function of the event listener handler of the application layer.
func (a *Service) EventListenerCallBack(channelName string, octoBot me.OctoBot) {
	if channelName == smodel.ChannelCreateOctoBot {
		a.CreateOctoBot(context.Background(), octoBot)
	} else if channelName == smodel.ChannelDeleteOctoBot {
		a.DeleteOctoBot(context.Background(), octoBot)
	} else {
		go a.Log(context.Background(), me.NewLogData().GenerateLogData(pb_logging.LogType_LogTypeERROR, "EventListenerCallBack", channelName, smodel.ErrorChannelNameNotValid.Error()))
	}
}
