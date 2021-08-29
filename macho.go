package gopwn

import (
	"bytes"
	"debug/macho"
	"fmt"
	"os"
)

type MACHO struct {
	file *macho.File
}

func NewMACHO(path string) (*MACHO, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewMACHOFromReader(fh)
}

func NewMACHOFromBytes(b []byte) (*MACHO, error) {
	r := bytes.NewReader(b)
	return NewMACHOFromReader(r)
}

func NewMACHOFromReader(r BinaryReader) (*MACHO, error) {
	f, err := macho.NewFile(r)
	if err != nil {
		return nil, err
	}
	return &MACHO{
		file: f,
	}, nil
}

func (m *MACHO) Close() error {
	return m.file.Close()
}

func (m *MACHO) Caves(caveSize int) []Cave {
	var caves []Cave
	for _, s := range m.file.Sections {
		data, _ := s.Data()
		caves = append(caves, searchCaves(fmt.Sprintf("%s.%s", s.Seg, s.Name), data, uint64(s.Offset), s.Addr, s.Size, parseMACHOFlags(s.Flags), caveSize)...)
	}
	return caves
}

func parseMACHOFlags(f uint32) string {
	return "TODO"
}

func (m *MACHO) Strings(optFns ...func(o *StringsOptions)) []string {
	options := StringsOptions{}
	for _, fn := range optFns {
		fn(&options)
	}

	var sections []dataReader
	if len(options.Sections) > 0 {
		for _, name := range options.Sections {
			sections = append(sections, m.file.Section(name))
		}
	} else {
		for _, s := range m.file.Sections {
			sections = append(sections, s)
		}
	}
	return parseStrings(sections)
}
