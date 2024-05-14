package application

import "context"

func (a QueryAdapter) Users(ctx context.Context, filter string) ([]string, error) {
	return a.Service.Users(ctx, filter)
}
