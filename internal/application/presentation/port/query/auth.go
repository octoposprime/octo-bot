package application

import "context"

type AuthQueryPort interface {
	Users(ctx context.Context, filter string) ([]string, error)
}
