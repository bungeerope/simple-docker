// +build linux

package container

import (
	logger "github.com/sirupsen/logrus"
	"os"
	"syscall"
)

func RunContainerInitProcess(command string, common interface{}) error {
	logger.Infof("command %s", command)

	// syscall.MS_NOEXEC: 在本文件系统中不允许运行其他应用程序
	// syscall.MS_NOUSER: 运行应用程序是不允许set_user_id、set_group_id
	// syscall.MS_NODEV: 所有mount参数使用默认设定的参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOUSER | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	// syscall.Exec() 调用了Linux kernel中的int execve(const *char filename,char *const argv[],char *const envp[])
	// 作用: 执行当前filename对应的程序。会覆盖当前进程的镜像，数据和堆栈等信息，包括PID，都将被要运行的程序覆盖
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logger.Errorf(err.Error())
	}
	return nil
}
