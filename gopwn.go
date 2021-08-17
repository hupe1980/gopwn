package gopwn

import (
	"github.com/hupe1980/gopwn/tubes"
)

func Process(path string, args ...string) *tubes.Process {
	p, err := tubes.NewProcess(path, args...)
	if err != nil {
		panic(err)
	}
	if err := p.Start(); err != nil {
		panic(err)
	}
	return p
}
