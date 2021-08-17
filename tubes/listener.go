package tubes

import "net"

type Listener struct {
	tube
	listener net.Listener
	con      net.Conn
}

func NewListener(addr string) (*Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{
		listener: l,
	}, nil
}

func (l *Listener) WaitForConnection() error {
	c, err := l.listener.Accept()
	if err != nil {
		return err
	}
	l.con = c
	return nil
}

func (l *Listener) Close() error {
	return l.listener.Close()
}
