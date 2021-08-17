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

type ProcessOptions struct {
	Path  string
	Args  []string
	Delim byte
}

func NewProcess(argv []string, optFns ...func(o *ProcessOptions)) (*Process, error) {
	options := ProcessOptions{
		Path: argv[0],
		Args: argv[1:],
	}
	for _, fn := range optFns {
		fn(&options)
	}
	cmd := exec.Command(options.Path, options.Args...)

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
