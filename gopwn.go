package gopwn

import (
	"github.com/hupe1980/gopwn/bins"
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

func Remote(network, addr string) (*tubes.Remote, error) {
	return tubes.NewRemote(network, addr)
}

func Listen(addr string) (*tubes.Listener, error) {
	return tubes.NewListener(addr)
}

func ELF(path string) (*bins.ELF, error) {
	return bins.NewELF(path)
}

func PE(path string) (*bins.PE, error) {
	return bins.NewPE(path)
}

func MACHO(path string) (*bins.MACHO, error) {
	return bins.NewMACHO(path)
}
