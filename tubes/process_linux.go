// +build linux

package tubes

import (
	"fmt"
	"os"
)

func (p *Process) CWD() string {
	cwd, err := os.Readlink(fmt.Sprintf("/proc/%d/cwd", p.Cmd.Process.Pid))
	if err != nil {
		panic(err)
	}
	return cwd
}
