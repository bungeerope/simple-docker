package subsystem

type CpuSubsystem struct {
}

func (c *CpuSubsystem) Name() string {
	return ""
}

func (c *CpuSubsystem) Set(path string, res *ResourceConfig) error {
	return nil
}

func (c *CpuSubsystem) Apply(path string, pid int) error {
	return nil
}

func (c *CpuSubsystem) Remove(path string) error {
	return nil
}
