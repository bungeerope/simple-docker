package cgroup

import (
	"github.com/bungeerope/simple-docker/src/docker/cgroup/subsystem"
	logger "github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path           string
	ResourceConfig *subsystem.ResourceConfig
}

func NewResourceConfig(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

func (c *CgroupManager) Set(res *subsystem.ResourceConfig) error {
	for _, subSysIn := range subsystem.SubsystemIns {
		subSysIn.Set(c.Path, res)
	}
	return nil
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subsystem.SubsystemIns {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subsystem.SubsystemIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			logger.Warnf("remove cgroup fail: %v", err)
		}
	}
	return nil
}
