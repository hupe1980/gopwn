package shellcode

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed amd64/*.asm amd64/**/*.asm
var tpls embed.FS

type shellcode struct {
	t *template.Template
}

func newShellcode(patterns ...string) (*shellcode, error) {
	t, err := template.ParseFS(tpls, patterns...)
	if err != nil {
		return nil, err
	}
	return &shellcode{t}, nil
}

func (s *shellcode) generate(name string, data interface{}) string {
	var buf bytes.Buffer
	if err := s.t.ExecuteTemplate(&buf, name, data); err != nil {
		panic(err)
	}
	return buf.String()
}
