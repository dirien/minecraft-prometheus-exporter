package config

import (
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
)

type Config struct {
	MetricsPath   *string
	WebConfig     *string
	ListenAddress *string
	WorldPath     *string
	RconAddress   *string
	RconPassword  *string
}

func NewConfg() *Config {
	var (
		webConfig     = webflag.AddFlags(kingpin.CommandLine)
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9150").String()
		worldPath     = kingpin.Flag("mc.world", "Path the to world folder").Default("/minecraft/world").String()
		rconAddress   = kingpin.Flag("mc.rcon-address", "Address of the Minecraft rcon.").Default(":25575").String()
		rconPassword  = kingpin.Flag("mc.rcon-password", "Password of the Minecraft rcon.").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
	)
	return &Config{
		MetricsPath:   metricsPath,
		WebConfig:     webConfig,
		ListenAddress: listenAddress,
		RconAddress:   rconAddress,
		RconPassword:  rconPassword,
		WorldPath:     worldPath,
	}
}
