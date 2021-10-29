package main

import (
	proc "github.com/adriamanu/thoth-agent/proc"
)

func main() {
	proc.Meminfo()
	proc.ProcStat()
}
