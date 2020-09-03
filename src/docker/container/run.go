// +build linux

package container

import (
	logger "github.com/sirupsen/logrus"
	"os"
)

func Run(tty bool, command string) {
	parent := NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logger.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}
