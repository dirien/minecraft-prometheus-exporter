package config

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	MetricsPath *string
}

func NewConfg() *Config {
	var (
		metricsPath = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
	)
	return &Config{
		MetricsPath: metricsPath,
	}
}
