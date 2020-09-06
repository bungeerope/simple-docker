package subsystem

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory.limit_in_bytes fail %v", err)
			}
		}
		return nil
	} else {
		return fmt.Errorf("get cgroup path fail %v", err)
	}
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}

func (s *MemorySubsystem) Apply(cgroupPath string, pid int) error {
	if subSysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		if err := ioutil.WriteFile(path.Join(subSysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err == nil {
			return nil
		} else {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
	} else {
		return fmt.Errorf("set cgroup proc fail %v", err)
	}
}

func (s *MemorySubsystem) Remove(cgroupPath string) error {
	if subSysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		return os.Remove(subSysCgroupPath)
	} else {
		return err
	}
}
