package cgroup

// Cgroup defines the configuration to limit resources by cgroups.
type Cgroup struct {
	CPUSet      string
	CPURate     string
	Memory      string
	TCPMemory   string
	DevicesDeny string
	DeviceAllow string
	BlkReadBps  string
	BlkWriteBps string
}
