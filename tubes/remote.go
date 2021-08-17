package tubes

import "net"

type Remote struct {
	tube
}

func NewRemote(addr string) (*Remote, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Remote{
		tube: tube{
			stdin:  c,
			stdout: c,
			stderr: c,
		},
	}, nil
}
