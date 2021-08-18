package gopwn

import (
	"debug/pe"
)

type PE struct {
	path string // Path to the file
	file *pe.File
}

func NewPE(path string) (*PE, error) {
	f, err := pe.Open(path)
	if err != nil {
		return nil, err
	}
	return &PE{
		path: path,
		file: f,
	}, nil
}

func (p *PE) Close() error {
	return p.file.Close()
}
