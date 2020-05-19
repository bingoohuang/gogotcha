package cgroup

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/dustin/go-humanize"
	"golang.org/x/sync/errgroup"
)

// Limit limits the resources by cgroups.
func (c Cgroup) Limit() error {
	var g errgroup.Group

	memoryBytes := ""
	tcpBytes := ""

	if c.Memory != "" {
		b, _ := humanize.ParseBytes(c.Memory)
		memoryBytes = fmt.Sprintf("%d", b)
	}

	if c.TCPMemory != "" {
		b, _ := humanize.ParseBytes(c.TCPMemory)
		tcpBytes = fmt.Sprintf("%d", b)
	}

	g.Go(func() error { return c.write("cpuset", "cpuset.cpus", c.CPUSet) })
	g.Go(func() error { return c.write("cpu", "cpu.cfs_quota_us", c.CPURate) })
	g.Go(func() error { return c.write("memory", "memory.limit_in_bytes", memoryBytes) })
	g.Go(func() error { return c.write("memory", "memory.kmem.tcp.limit_in_bytes", tcpBytes) })
	g.Go(func() error { return c.write("devices", "devices.deny", c.DevicesDeny) })
	g.Go(func() error { return c.write("devices", "devices.allow", c.DeviceAllow) })
	g.Go(func() error { return c.write("blkio", "blkio.throttle.read_bps_device", c.BlkReadBps) })
	g.Go(func() error { return c.write("blkio", "blkio.throttle.write_bps_device", c.BlkWriteBps) })

	if err := g.Wait(); err != nil {
		logrus.Warnf("fail to limit %v", err)
	}

	return err
}

func (c Cgroup) write(subp, f, data string) error {
	if data == "" {
		return nil
	}

	pid := strconv.Itoa(os.Getpid())
	cp := filepath.Join("/sys/fs/cgroup", subp, filepath.Base(os.Args[0]))
	if err := checkPath(cp); err != nil {
		return err
	}

	w := func(filename, data string) error {
		f := filepath.Join(cp, filename)
		logrus.Infof("write %s with data %s", filename, data)

		return ioutil.WriteFile(f, []byte(data), 0600)
	}

	var g errgroup.Group

	g.Go(func() error { return w(f, data) })
	g.Go(func() error { return w("tasks", pid) })

	return g.Wait()
}

func checkPath(p string) error {
	_, err := os.Stat(p)
	if os.IsNotExist(err) {
		logrus.Infof("try to MkdirAll %s", p)

		return os.MkdirAll(p, 755)
	}

	return err
}
