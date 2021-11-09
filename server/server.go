package server

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/adriamanu/thoth-agent/cpu"
	"github.com/adriamanu/thoth-agent/disk"
	"github.com/adriamanu/thoth-agent/memory"
)

type healthStruct struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	healthAsJson, err := json.Marshal(healthStruct{"alive", time.Now().String()})
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "%s", healthAsJson)
}

type resourceUsage struct {
	Cpu    []cpu.ProcStat    `json:"cpu"`
	Memory memory.MemoryStat `json:"memory"`
	Disk   disk.DiskStat     `json:"disk"`
	Date   string            `json:"date"`
}

func resourceUsageHandler(res http.ResponseWriter, req *http.Request) {
	c := cpu.Stat()
	m := memory.Stat()
	d := disk.Stat()
	ru, err := json.Marshal(resourceUsage{c, m, d, time.Now().String()})
	if err != nil {
		log.Fatal(err)
	}
	res.WriteHeader(200)
	res.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(res, "%s", ru)
}

//go:embed index.html
var index []byte

func visualizationHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	fmt.Fprintf(res, "%s", string(index))
}

func Server() {
	http.HandleFunc("/", visualizationHandler)
	http.HandleFunc("/heartbeat", healthCheckHandler)
	http.HandleFunc("/resource-usage", resourceUsageHandler)

	log.Fatal(http.ListenAndServe(":11805", nil))
}
