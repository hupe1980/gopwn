// +build darwin

package tubes

/*
#include <libproc.h>

const char* getwd(int p, char* cwd) {
    pid_t pid;
    struct proc_vnodepathinfo vpi;

	pid = (pid_t) p;
	proc_pidinfo(pid, PROC_PIDVNODEPATHINFO, 0, &vpi, sizeof(vpi));

	cwd = vpi.pvi_cdir.vip_path;
	return cwd;
}
*/
import (
	"C"
)

func (p *Process) CWD() string {
	var cRetValue *C.char
	cRetValue = C.getwd(C.int(p.cmd.Process.Pid), cRetValue)
	return C.GoString(cRetValue)
}
