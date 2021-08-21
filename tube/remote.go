package tube

import "net"

type Remote struct {
	tube
	conn net.Conn
}

func NewRemote(network, addr string) (*Remote, error) {
	c, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return &Remote{
		conn: c,
		tube: tube{
			stdin:  c,
			stdout: c,
			stderr: c,
		},
	}, nil
}

func (r *Remote) Close() error {
	return r.conn.Close()
}
