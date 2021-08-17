package tubes

import (
	"os/exec"
)

type Process struct {
	tube
	cmd *exec.Cmd
}

type ProcessOptions struct {
	Path    string
	Args    []string
	NewLine byte
}

func NewProcess(argv []string, optFns ...func(o *ProcessOptions)) (*Process, error) {
	options := ProcessOptions{
		Path:    argv[0],
		Args:    argv[1:],
		NewLine: '\n',
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
		tube: tube{
			stdin:  stdin,
			stdout: stdout,
			stderr: stderr,

			newLine: options.NewLine,
		},
	}, nil
}

func (p *Process) Start() error {
	return p.cmd.Start()
}
