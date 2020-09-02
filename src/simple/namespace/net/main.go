// +build linux

// Network Namespace 用来隔离网络设备、IP地址端口等网络栈。
// 可以提供独立虚拟的网络设备，同时Namespace内的应用进程可以绑定到指定端口，各Namespace间的端口将不冲突。
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
		Cloneflags: syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
