package domain

var ERRORS []error = []error{
	ErrorNone,
}

var (
	ErrorNone error = nil
)

func GetErrors() []error {
	return ERRORS
}
