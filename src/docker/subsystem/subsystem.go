package subsystem

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
