package main

import (
	"fmt"
	"github.com/bungeerope/simple-docker/src/docker/container"
	"github.com/bungeerope/simple-docker/src/docker/subsystem"
	logger "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `simple-docker is a simple container runtime implementation.
				Enjoy it, just for fun.`
const runUsage = `Create a container with namespace and cgroup limit simple-docker run -it [command]`

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user`s process in container.Do not call it outside.",
	Action: func(content *cli.Context) error {
		logger.Infof("init come on")
		cmd := content.Args().Get(0)
		logger.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

var runCommand = cli.Command{
	Name:  "run",
	Usage: runUsage,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.BoolFlag{
			Name:  "c",
			Usage: "cpuset limit",
		},
		cli.BoolFlag{
			Name:  "cs",
			Usage: "cpushare limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}

		var commands []string
		for _, arg := range context.Args() {
			commands = append(commands, arg)
		}

		tty := context.Bool("ti")
		resConf := &subsystem.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("c"),
			CpuShare:    context.String("cs"),
		}

		// container运行参数
		// cmd := context.Args()[0]
		// container.Run(tty, cmd)

		subsystem.Run(tty, commands, resConf)
		return nil
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "simple-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		logger.SetFormatter(&logger.JSONFormatter{})
		logger.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
