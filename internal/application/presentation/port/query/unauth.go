package application

import "context"

type UnAuthQueryPort interface {
	Calculate(ctx context.Context, value int64) (int64, error)
}
