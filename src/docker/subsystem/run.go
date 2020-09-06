package subsystem

import (
	"github.com/bungeerope/simple-docker/src/docker/subsystem/pipline"
	logger "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func Run(tty bool, commands []string, res *ResourceConfig) {
	parent, writePipe := pipline.NewParentProcess(tty)
	if parent == nil {
		logger.Errorf("New Parent Process error")
		return
	}
	if err := parent.Start(); err != nil {
		logger.Error(err)
	}

	//cgroupManager := cgroup.NewResourceConfig("simple-docker")
	//defer cgroupManager.Destroy()
	//cgroupManager.Set(res)
	//cgroupManager.Apply(parent.Process.Pid)
	sendInitCommand(commands, writePipe)
	parent.Wait()
	os.Exit(0)
}

func sendInitCommand(commands []string, pipe *os.File) {
	commandString := strings.Join(commands, " ")
	logger.Infof("command all is %s", commandString)
	pipe.WriteString(commandString)
	pipe.Close()
}
