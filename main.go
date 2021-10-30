package main

import (
	"fmt"
	"time"

	cpu "github.com/adriamanu/thoth-agent/cpu"
	disk "github.com/adriamanu/thoth-agent/disk"
	memory "github.com/adriamanu/thoth-agent/memory"
)

func main() {
	for {
	memory := memory.Stat()
	fmt.Printf("Memory percentage used: %.3f\n", memory.UsedMemoryPercentage)

	cpu := cpu.Stat()
	fmt.Printf("CPU percentage used: %.3f\n", cpu.UsedPercentage)

	disk := disk.Stat()
	fmt.Printf("Disk percentage used: %.3f\n\n",disk.UsedSpacePercentage)

	// time.Sleep(1 * time.Second)
	time.Sleep(200 * time.Millisecond)
}
}
