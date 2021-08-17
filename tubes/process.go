package tubes

import (
	"bufio"
	"io"
	"os/exec"
)

type Process struct {
	cmd    *exec.Cmd
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser
	Delim  byte
}

func NewProcess(path string, args ...string) (*Process, error) {
	cmd := exec.Command(path, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	return &Process{
		cmd: cmd,

		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,

		// Character sent with methods like SendLine() or used for RecvLine()
		Delim: '\n',
	}, nil
}

func (p *Process) Start() error {
	return p.cmd.Start()
}

func (p *Process) SendLine(input interface{}) (int, error) {
	b := Bytes(input)
	b = append(b, p.Delim)
	return p.Stdin.Write(b)
}

func (p *Process) RecvLine() ([]byte, error) {
	rd := bufio.NewReader(p.Stdout)
	return rd.ReadBytes(p.Delim)
}
