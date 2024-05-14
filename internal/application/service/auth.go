package application

import "context"

func (a *Service) Users(ctx context.Context, filter string) ([]string, error) {
	return a.DbPort.GetUsersByFilter(ctx, filter)
}
