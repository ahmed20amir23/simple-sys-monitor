package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Response struct for JSON output
type Metrics struct {
	CPUUsage    float64 `json:"cpu_usage_percent"`
	MemoryUsage float64 `json:"memory_usage_percent"`
}

func getMetrics() Metrics {
	// CPU usage
	cpuPercent, _ := cpu.Percent(0, false)

	// Memory usage
	vmStat, _ := mem.VirtualMemory()

	return Metrics{
		CPUUsage:    cpuPercent[0],
		MemoryUsage: vmStat.UsedPercent,
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := getMetrics()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func main() {
	http.HandleFunc("/metrics", metricsHandler)

	log.Println("✅ Server running on http://localhost:8080/metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
