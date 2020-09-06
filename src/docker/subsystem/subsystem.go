package subsystem

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// 限制cgroup
type ResourceConfig struct {
	MemoryLimit string
	CpuSet      string
	CpuShare    string
}

// 定义Subsystem接口
type Subsystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	SubsystemIns = []Subsystem{
		&CpusetSubsystem{},
		&CpuSubsystem{},
		&MemorySubsystem{},
	}
)

const cgroupRoot = `/proc/self/mountinfo`

func FindCgroupMountPoint(subsystem string) string {

	file, err := os.Open(cgroupRoot)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, "")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint(subsystem)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err != nil {
				return "", fmt.Errorf("error create cgroup : %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("create cgroup error: %v", err)
	}
}
