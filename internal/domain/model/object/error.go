package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorOctoBotIsEmpty,
	ErrorOctoBotOctoBotDataIsEmpty,
}

const (
	ErrId          string = "id"
	ErrOctoBot     string = "octoBot"
	ErrOctoBotData string = "octoBotdata"
)

const (
	ErrEmpty         string = "empty"
	ErrTooShort      string = "tooshort"
	ErrTooLong       string = "toolong"
	ErrNotValid      string = "notvalid"
	ErrInactive      string = "inactive"
	ErrAlreadyExists string = "alreadyexists"
)

var (
	ErrorNone                      error = nil
	ErrorOctoBotIsEmpty            error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrOctoBot + smodel.ErrSep + ErrEmpty)
	ErrorOctoBotOctoBotDataIsEmpty error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrOctoBot + smodel.ErrSep + ErrOctoBotData + smodel.ErrSep + ErrEmpty)
)

func GetErrors() []error {
	return ERRORS
}
