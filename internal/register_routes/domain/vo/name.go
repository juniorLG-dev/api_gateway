package vo

import (
	"gateway/internal/configuration/handler_err"

	"regexp"
)

type Name struct {
	Value string
}

func NewName(name string) (*Name, *handler_err.InfoErr) {
	re := regexp.MustCompile(`^[A-Za-z0-9_-]+$`)

	if re.MatchString(name) {
		return &Name{
			Value: name,
		}, &handler_err.InfoErr{}
	}

	return &Name{}, &handler_err.InfoErr{
		Message: "your service name can only contain letters (upper or lower case), numbers (0-9) and these two special characters: \"_\" and \"-\". It cannot have spaces",
		Err: handler_err.ErrInvalidInput,
	}
}