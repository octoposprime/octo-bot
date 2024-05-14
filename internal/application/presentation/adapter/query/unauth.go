package application

import "context"

func (a QueryAdapter) Calculate(ctx context.Context, value int64) (int64, error) {
	return a.Service.Calculate(ctx, value)
}
