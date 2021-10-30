package main

import (
	"log"

	unix "golang.org/x/sys/unix"
)

const rootPath string = "/"

type Disk struct {
	total     uint64
	available uint64
	used      uint64
	usedPercentage float64
}

func Total(disk unix.Statfs_t) uint64 {
	return disk.Blocks * uint64(disk.Bsize)
}

func Available(disk unix.Statfs_t) uint64 {
	return disk.Bfree * uint64(disk.Bsize)
}

func Used(total, available uint64) uint64 {
	return total - available
} 

func UsedPercentage(total, available uint64) float64 {
	return float64(available) / float64(total) * 100.0
}

func DiskUsage() Disk {
	var diskStat unix.Statfs_t
	err := unix.Statfs(rootPath, &diskStat)
	if err != nil {
		log.Fatal(err)
	}

	total := Total(diskStat)
	available := Available(diskStat)	
	used := Used(total, available)
	usedPercentage := UsedPercentage(total, available) 

	return Disk{
		total: total,
		available: available,
		used: used,
		usedPercentage: usedPercentage,
	}
}
