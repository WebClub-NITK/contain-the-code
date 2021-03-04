package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

// docker run ubuntu sleep 100
// docker         run <image> <cmd> <params>
// go run main.go run         <cmd> <params>

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Bad parameter")
	}
}

func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	// go run main.go child /bin/bash
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cg()

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("ubuntufs")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	// /bin/bash
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	syscall.Unmount("proc", 0)
}

func cg() {
	pids := "/sys/fs/cgroup/pids"

	// We will create a "cgroup rule"
	os.Mkdir(pids+"/container", 0755)

	// Set max pid
	os.WriteFile(pids+"/container/pids.max", []byte("20"), 0700)

	// Add current process to cgroup
	os.WriteFile(pids+"/container/cgroup.procs", []byte(strconv.Itoa(os.Getpid())), 0700)

	os.WriteFile(pids+"/container/notify_on_release", []byte("1"), 0700)
}

/*

go run main.go run /bin/bash
	|
	|
	V
go run main.go child /bin/bash [New UTS Namespace]
	|
	|
	V
/bin/bash [Same UTS Namespace]

*/
