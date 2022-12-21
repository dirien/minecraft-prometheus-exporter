package config

import (
	"os"
	"path/filepath"

	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ConfigPath             *string
	FlagConfig             *web.FlagConfig
	MetricsPath            *string         `yaml:"metrics-path"`
	ListenAddress          *string         `yaml:"listen-address"`
	WorldPath              *string         `yaml:"world-path"`
	RconAddress            *string         `yaml:"rcon-address"`
	RconPassword           *string         `yaml:"rcon-password"`
	NameSource             *string         `yaml:"name-source"`
	DisableExporterMetrics *bool           `yaml:"disable-exporter-metrics"`
	DisabledMetrics        map[string]bool `yaml:"disabled-metrics"`
	ModServerStats         *string         `yaml:"mod-server-stats"`
}

func NewConfg() *Config {
	var (
		flagConfig              = webflag.AddFlags(kingpin.CommandLine, ":9150")
		metricsPath             = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Envar("WEB_TELEMETRY_PATH").Default("/metrics").String()
		disabledExporterMetrics = kingpin.Flag("web.disable-exporter-metrics", "Disabling collection of exporter metrics (like go_*)").Envar("WEB_DISABLED_EXPORTER_METRICS").Bool()
		configPath              = kingpin.Flag("mc.config-path", "Path to YAML file with config.").Envar("MC_CONFIG_PATH").Default("config.yml").String()
		worldPath               = kingpin.Flag("mc.world", "Path the to world folder").Envar("MC_WORLD").Default("/minecraft/world").String()
		rconAddress             = kingpin.Flag("mc.rcon-address", "Address of the Minecraft rcon.").Envar("MC_RCON_ADDRESS").Default(":25575").String()
		rconPassword            = kingpin.Flag("mc.rcon-password", "Password of the Minecraft rcon.").Envar("MC_RCON_PASSWORD").String()
		nameSource              = kingpin.Flag("mc.name-source", "How to retrieve names of players: offline, bukkit, mojang.").Envar("MC_NAME_SOURCE").Default("mojang").String()
		modServerStats          = kingpin.Flag("mc.mod-server-stats", "Set server for additional stats (papermc, purpurmc or forge)").Envar("MC_MOD_SERVER_STATS").String()
	)
	return &Config{
		ConfigPath:             configPath,
		DisableExporterMetrics: disabledExporterMetrics,
		MetricsPath:            metricsPath,
		FlagConfig:             flagConfig,
		RconAddress:            rconAddress,
		RconPassword:           rconPassword,
		WorldPath:              worldPath,
		NameSource:             nameSource,
		ModServerStats:         modServerStats,
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
