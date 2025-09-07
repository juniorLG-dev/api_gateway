package service

import (
	"path/filepath"
	"strings"
)

type File struct {
	value string
}

func NewCheckFile(filename string) *File {
	return &File{
		value: filename,
	}
}

func (f *File) Check(ext string) bool {
	return strings.ToLower(filepath.Ext(f.value)) == ext
}