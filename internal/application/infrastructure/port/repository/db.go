package application

import "context"

// DbPort is a port for Hexagonal Architecture Pattern.
// It is used to communicate with the database.
type DbPort interface {

	// GetUsersByFilter returns the users that match the given filter.
	GetUsersByFilter(ctx context.Context, filter string) ([]string, error)
}
