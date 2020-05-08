//+build !linux

package cgroup

// Limit limits the resources by cgroups.
func (c Cgroup) Limit() error {
	return nil
}
