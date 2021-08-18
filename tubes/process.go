package tubes

import (
	"io"
	"os"
	"os/exec"
)

type Process struct {
	tube
	cmd *exec.Cmd
}

type ProcessOptions struct {
	Env     []string
	Dir     string
	NewLine byte
}

func NewProcess(argv []string, optFns ...func(o *ProcessOptions)) (*Process, error) {
	options := ProcessOptions{
		NewLine: '\n',
	}
	for _, fn := range optFns {
		fn(&options)
	}

	cmd := exec.Command(argv[0], argv[1:]...)
	cmd.Env = append(os.Environ(), options.Env...)
	cmd.Dir = options.Dir

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
			stdin:   stdin,
			stdout:  stdout,
			stderr:  stderr,
			newLine: options.NewLine,
		},
	}, nil
}

// Start starts the specified command but does not wait for it to complete.
func (p *Process) Start() error {
	return p.cmd.Start()
}

func (p Process) Interactive() error {
	go io.Copy(p.tube.stdin, os.Stdin)
	go io.Copy(os.Stdout, p.tube.stdout)
	go io.Copy(os.Stderr, p.tube.stderr)

	// Wait for the process to exit
	return p.cmd.Wait()
}

// Kill causes the Process to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
func (p *Process) Kill() error {
	return p.cmd.Process.Kill()
}

// Signal sends a signal to the Process.
// Sending Interrupt on Windows is not implemented.
func (p *Process) Signal(sig os.Signal) error {
	return p.cmd.Process.Signal(sig)
}
