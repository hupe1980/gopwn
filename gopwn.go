package gopwn

import (
	"github.com/hupe1980/gopwn/tubes"
)

func Process(argv []string, optFns ...func(o *tubes.ProcessOptions)) (*tubes.Process, error) {
	p, err := tubes.NewProcess(argv, optFns...)
	if err != nil {
		return nil, err
	}
	if err := p.Start(); err != nil {
		return nil, err
	}
	return p, nil
}

func Remotee(addr string) (*tubes.Remote, error) {
	return tubes.NewRemote(addr)
}

func Listen(addr string) (*tubes.Listener, error) {
	return tubes.NewListener(addr)
}
