package dto

import (
	"io"
)

type SendRequestInput struct {
	Path        string
	ServiceName string
	Body        io.Reader
}

type SendRequestOutput struct {
	Response string
	Status   int
	Header   string
}
