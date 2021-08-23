package shellcode

import (
	_ "embed"
)

type AMD64 struct {
	*shellcode
	Linux *amd64Linux
}

func NewAMD64() (*AMD64, error) {
	sc, err := newShellcode("amd64/*.asm")
	if err != nil {
		return nil, err
	}
	lsc, err := newShellcode("amd64/linux/*.asm")
	if err != nil {
		return nil, err
	}
	return &AMD64{
		shellcode: sc,
		Linux: &amd64Linux{
			shellcode: lsc,
		},
	}, nil
}

func (a *AMD64) NOP() string {
	return a.generate("nop.asm", nil)
}

type amd64Linux struct {
	*shellcode
}

func (l *amd64Linux) SH() string {
	return l.generate("sh.asm", nil)
}
