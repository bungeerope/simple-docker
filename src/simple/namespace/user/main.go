// +build linux

// User Namespace主要用于隔离用户及用户组
package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER,
		// Question: 以下代码是issue，待解决。将会出现`fork/exec /usr/bin/sh: no such file or directory`的错误。
		// Answer: 以下代码OK。出现上述问题是因为`mount -m proc proc /proc`后未执行`umount`导致。
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1,
				HostID:      1,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1,
				HostID:      1,
				Size:        1,
			},
		},
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(-1)
}
