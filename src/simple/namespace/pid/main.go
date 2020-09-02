// +build linux

// 通过PID Namespace将针对进程ID进行隔离
// 在宿主机查询进程树看到的进程ID与通过PID Namespace隔离后的进程应用查看PID存在不一致的情况
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 通过新进程fork一个'sh' terminal
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
