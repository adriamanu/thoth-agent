package server

import (
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
	Cpu    cpu.CPUStat       `json:"cpu"`
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

func visualizationHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./server/chart.html")
}

func Server() {
	http.HandleFunc("/", healthCheckHandler)
	http.HandleFunc("/resource-usage", resourceUsageHandler)
	http.HandleFunc("/visualization", visualizationHandler)

	log.Fatal(http.ListenAndServe(":11805", nil))
}
