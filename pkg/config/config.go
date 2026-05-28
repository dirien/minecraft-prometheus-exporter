package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/alecthomas/kong"
	"github.com/prometheus/exporter-toolkit/web"
	"gopkg.in/yaml.v3"
)

// Config holds all flag/env/YAML-driven configuration. Tags:
//   - kong tags define the CLI surface (flag name, env var, default).
//   - yaml tags define the YAML overlay loaded from --mc.config-path.
//
// LoadFile() unmarshals YAML over the already-parsed flag values, so the YAML
// file takes precedence when present — matching the prior behaviour.
type Config struct {
	ListenAddresses        []string `kong:"name='web.listen-address',env='WEB_LISTEN_ADDRESS',default=':9150',help='Addresses on which to expose metrics and web interface. Repeatable for multiple addresses.'" yaml:"listen-address"`
	WebConfigFile          string   `kong:"name='web.config.file',env='WEB_CONFIG_FILE',default='',help='Path to configuration file that can enable TLS or authentication. See: https://github.com/prometheus/exporter-toolkit/blob/master/docs/web-configuration.md'" yaml:"-"`
	WebSystemdSocket       bool     `kong:"name='web.systemd-socket',env='WEB_SYSTEMD_SOCKET',help='Use systemd socket activation listeners instead of port listeners (Linux only).'" yaml:"-"`
	MetricsPath            string   `kong:"name='web.telemetry-path',env='WEB_TELEMETRY_PATH',default='/metrics',help='Path under which to expose metrics.'" yaml:"metrics-path"`
	DisableExporterMetrics bool     `kong:"name='web.disable-exporter-metrics',env='WEB_DISABLED_EXPORTER_METRICS',help='Disable collection of exporter metrics (like go_*).'" yaml:"disable-exporter-metrics"`

	ConfigPath     string `kong:"name='mc.config-path',env='MC_CONFIG_PATH',default='config.yml',help='Path to YAML file with config.'" yaml:"-"`
	WorldPath      string `kong:"name='mc.world',env='MC_WORLD',default='/minecraft/world',help='Path to the world folder.'" yaml:"world-path"`
	RconAddress    string `kong:"name='mc.rcon-address',env='MC_RCON_ADDRESS',default=':25575',help='Address of the Minecraft rcon.'" yaml:"rcon-address"`
	RconPassword   string `kong:"name='mc.rcon-password',env='MC_RCON_PASSWORD',help='Password of the Minecraft rcon.'" yaml:"rcon-password"`
	NameSource     string `kong:"name='mc.name-source',env='MC_NAME_SOURCE',default='mojang',help='How to retrieve names of players: offline, bukkit, mojang.'" yaml:"name-source"`
	ModServerStats string `kong:"name='mc.mod-server-stats',env='MC_MOD_SERVER_STATS',help='Set server for additional stats (papermc, purpurmc, forge, or fabric).'" yaml:"mod-server-stats"`

	LogLevel  string `kong:"name='log.level',env='LOG_LEVEL',default='info',help='Log level. One of: [debug, info, warn, error]'" yaml:"-"`
	LogFormat string `kong:"name='log.format',env='LOG_FORMAT',default='logfmt',help='Log format. One of: [logfmt, json]'" yaml:"-"`

	DisabledMetrics map[string]bool `kong:"-" yaml:"disabled-metrics"`

	Version kong.VersionFlag `kong:"name='version',short='V',help='Show version and exit.'" yaml:"-"`
}

// WebFlagConfig returns the exporter-toolkit FlagConfig assembled from the
// parsed values. Systemd socket activation is Linux-only; on other platforms
// the field is forced to false to match upstream behaviour.
func (c *Config) WebFlagConfig() *web.FlagConfig {
	systemd := c.WebSystemdSocket
	if runtime.GOOS != "linux" {
		systemd = false
	}
	return &web.FlagConfig{
		WebListenAddresses: &c.ListenAddresses,
		WebSystemdSocket:   &systemd,
		WebConfigFile:      &c.WebConfigFile,
	}
}

func (c *Config) LoadFile() {
	path, err := filepath.Abs(c.ConfigPath)
	if err != nil {
		return
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return
	}

	_ = yaml.Unmarshal(bytes, c)
}
