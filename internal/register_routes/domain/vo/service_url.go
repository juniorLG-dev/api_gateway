package vo

import (
	"regexp"

	"gateway/internal/configuration/handler_err"
)

type ServiceURL struct {
	Value string
}

func NewServiceURL(url string) (*ServiceURL, *handler_err.InfoErr) {
	re := regexp.MustCompile(`^(http|https)://`)

	if re.MatchString(url) {
		return &ServiceURL{
			Value: url,
		}, &handler_err.InfoErr{}
	}

	return &ServiceURL{}, &handler_err.InfoErr{
		Message: "enter a valid url",
		Err: handler_err.ErrInvalidInput,
	}
}