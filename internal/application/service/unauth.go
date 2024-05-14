package application

import "context"

func (a *Service) Calculate(ctx context.Context, value int64) (int64, error) {
	return value * 2, nil
}
