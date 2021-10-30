package disk

import (
	"log"

	unix "golang.org/x/sys/unix"
)

const rootPath string = "/"

type DiskStat struct {
	TotalSpace     uint64
	AvailableSpace uint64
	UsedSpace      uint64
	UsedSpacePercentage float64
}

func Total(disk unix.Statfs_t) uint64 {
	return disk.Blocks * uint64(disk.Bsize)
}

func Available(disk unix.Statfs_t) uint64 {
	return disk.Bfree * uint64(disk.Bsize)
}

func Used(disk unix.Statfs_t) uint64 {
	return Total(disk) - Available(disk)
} 

func UsedPercentage(disk unix.Statfs_t) float64 {
	return float64(Used(disk)) / float64(Total(disk)) * 100.0
}

func Stat() DiskStat {
	var diskStat unix.Statfs_t
	err := unix.Statfs(rootPath, &diskStat)
	if err != nil {
		log.Fatal(err)
	}

	total := Total(diskStat)
	available := Available(diskStat)	
	used := Used(diskStat)
	usedPercentage := UsedPercentage(diskStat) 

	return DiskStat{
		TotalSpace: total,
		AvailableSpace: available,
		UsedSpace: used,
		UsedSpacePercentage: usedPercentage,
	}
}
