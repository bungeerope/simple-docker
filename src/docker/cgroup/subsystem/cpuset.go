package subsystem

type CpusetSubsystem struct {
}

func (c *CpusetSubsystem) Name() string {
	return ""
}

func (c *CpusetSubsystem) Set(path string, res *ResourceConfig) error {
	return nil
}

func (c *CpusetSubsystem) Apply(path string, pid int) error {
	return nil
}

func (c *CpusetSubsystem) Remove(path string) error {
	return nil
}
