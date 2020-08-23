package overview

import (
	"github.com/bitcav/nitr-core/cpu"
	"github.com/bitcav/nitr-core/host"
	"github.com/bitcav/nitr-core/ram"
)

type Overview struct {
	Host     host.HostInfo `json:"host"`
	CPUUsage float64       `json:"cpuUsage"`
	RAM      ram.RAM       `json:"ram"`
}

func Info() Overview {
	return Overview{
		Host:     host.Info(),
		CPUUsage: cpu.CpuUsage(),
		RAM:      ram.Info(),
	}
}
