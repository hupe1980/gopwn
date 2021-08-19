// +build linux

package tubes

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/v3/process"
)

func (p *Process) CWD() (string, error) {
	process, err := process.NewProcess(int32(p.cmd.Process.Pid))
	if err != nil {
		return "", err
	}
	return process.Cwd()
}
