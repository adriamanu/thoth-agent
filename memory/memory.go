package memory

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/adriamanu/thoth-agent/utils"
)

const meminfoFilePath string = "/proc/meminfo"

type MemInfo struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
	SwapTotal    int
	SwapFree     int
}

type MemoryStat struct {
	TotalMemory          int
	AvailableMemory      int
	UsedMemory           int
	UsedMemoryPercentage float64
}

func getKbSizeFromLine(line string) int {
	kbSize, err := strconv.Atoi(strings.Split(strings.Trim(line, " "), " ")[0])
	if err != nil {
		log.Fatal(err)
	}

	return kbSize
}

func parseMeminfoFile(meminfoFileContent string) MemInfo {
	var m MemInfo
	splittedMemInfo := utils.SplitFilesByLine(meminfoFileContent)

	for i := 0; i < len(splittedMemInfo); i++ {
		line := strings.Split(splittedMemInfo[i], ":")
		switch line[0] {
		case "MemTotal":
			m.MemTotal = getKbSizeFromLine(line[1])
		case "MemFree":
			m.MemFree = getKbSizeFromLine(line[1])
		case "MemAvailable":
			m.MemAvailable = getKbSizeFromLine(line[1])
		case "SwapTotal":
			m.SwapTotal = getKbSizeFromLine(line[1])
		case "SwapFree":
			m.SwapFree = getKbSizeFromLine(line[1])
		}
	}

	return m
}

func UsedMemory(available, total int) int {
	return total - available
}

func UsedPercentage(available, total int) float64 {
	return float64(UsedMemory(available, total)) / float64(available) * 100.0
}

func Stat() MemoryStat {
	meminfoFileContent, err := utils.ReadFile(meminfoFilePath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot open file %s - %v", meminfoFilePath, err))
	}

	meminfo := parseMeminfoFile(meminfoFileContent)

	return MemoryStat{
		TotalMemory:          meminfo.MemTotal,
		AvailableMemory:      meminfo.MemAvailable,
		UsedMemory:           UsedMemory(meminfo.MemAvailable, meminfo.MemTotal),
		UsedMemoryPercentage: UsedPercentage(meminfo.MemAvailable, meminfo.MemTotal),
	}
}
