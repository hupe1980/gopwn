package gopwn

import (
	"github.com/hupe1980/gopwn/tubes"
)

func Process(argv []string, optFns ...func(o *tubes.ProcessOptions)) *tubes.Process {
	p, err := tubes.NewProcess(argv, optFns...)
	if err != nil {
		panic(err)
	}
	if err := p.Start(); err != nil {
		panic(err)
	}
	return p
}

func Remotee(addr string) *tubes.Remote {
	r, err := tubes.NewRemote(addr)
	if err != nil {
		panic(err)
	}
	return r
}

func Listen(addr string) *tubes.Listener {
	l, err := tubes.NewListener(addr)
	if err != nil {
		panic(err)
	}
	return l
}
