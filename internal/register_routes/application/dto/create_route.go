package dto

import (
	"io"
)

type CreateRouteInput struct {
	Filename string
	File 		 io.Reader
	Token    string
}