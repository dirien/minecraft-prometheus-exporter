package config

import (
	"os"
	"path/filepath"

	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ConfigPath      *string
	MetricsPath     *string         `yaml:"metrics-path"`
	WebConfig       *string         `yaml:"web-config"`
	ListenAddress   *string         `yaml:"listen-address"`
	WorldPath       *string         `yaml:"world-path"`
	RconAddress     *string         `yaml:"rcon-address"`
	RconPassword    *string         `yaml:"rcon-password"`
	NameSource      *string         `yaml:"name-source"`
	DisabledMetrics map[string]bool `yaml:"disabled-metrics"`
}

func NewConfg() *Config {
	var (
		webConfig     = webflag.AddFlags(kingpin.CommandLine)
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Envar("WEB_TELEMETRY_PATH").Default("/metrics").String()
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Envar("WEB_LISTEN_ADDRESS").Default(":9150").String()
		configPath    = kingpin.Flag("mc.config-path", "Path to YAML file with config.").Envar("MC_CONFIG_PATH").Default("config.yml").String()
		worldPath     = kingpin.Flag("mc.world", "Path the to world folder").Envar("MC_WORLD").Default("/minecraft/world").String()
		rconAddress   = kingpin.Flag("mc.rcon-address", "Address of the Minecraft rcon.").Envar("MC_RCON_ADDRESS").Default(":25575").String()
		rconPassword  = kingpin.Flag("mc.rcon-password", "Password of the Minecraft rcon.").Envar("MC_RCON_PASSWORD").String()
		nameSource    = kingpin.Flag("mc.name-source", "How to retrieve names of players: offline, bukkit, mojang.").Envar("MC_NAME_SOURCE").Default("mojang").String()
	)
	return &Config{
		ConfigPath:    configPath,
		MetricsPath:   metricsPath,
		WebConfig:     webConfig,
		ListenAddress: listenAddress,
		RconAddress:   rconAddress,
		RconPassword:  rconPassword,
		WorldPath:     worldPath,
		NameSource:    nameSource,
	}
}

func (c *Config) LoadFile() {
	path, err := filepath.Abs(*c.ConfigPath)
	if err != nil {
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return
	}
}
