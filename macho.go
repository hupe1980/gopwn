package gopwn

import (
	"debug/macho"
)

type MACHO struct {
	path string // Path to the file
	file *macho.File
}

func NewMACHO(path string) (*MACHO, error) {
	f, err := macho.Open(path)
	if err != nil {
		return nil, err
	}
	return &MACHO{
		path: path,
		file: f,
	}, nil
}

func (m *MACHO) Close() error {
	return m.file.Close()
}
