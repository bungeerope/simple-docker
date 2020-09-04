package subsystem

import (
	"fmt"
	"io/ioutil"
	"path"
)

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err != nil {
		if res.MemoryLimit != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory.limit_in_bytes fail %v", err)
			}
		}
	}
}

func (s *MemorySubsystem) Name() string {
	panic("implement me")
}

func (s *MemorySubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (s *MemorySubsystem) Remove(path string) error {
	panic("implement me")
}
