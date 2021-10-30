package main

import (
	"fmt"
	"time"
)

func main() {
	for {
	meminfo := Meminfo()
	memoryPercent := float64(meminfo.MemTotal - meminfo.MemAvailable)/float64(meminfo.MemTotal) * 100.0
	fmt.Printf("Memory percentage used: %.3f\n", memoryPercent)

	cpuPercent := getCPUPercentage()
	fmt.Printf("CPU percentage used: %.3f\n", cpuPercent)

	disk := DiskUsage()
	fmt.Printf("Disk percentage used: %.3f\n\n",disk.usedPercentage)

	// time.Sleep(1 * time.Second)
	time.Sleep(200 * time.Millisecond)
}
}
