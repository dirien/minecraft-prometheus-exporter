package minecraft_prometheus_exporter

import (
	"github.com/go-kit/kit/log/level"
	"github.com/minecraft-prometheus-exporter/pkg/config"
	"github.com/minecraft-prometheus-exporter/pkg/exporter"
	"github.com/minecraft-prometheus-exporter/pkg/template"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

func Run() {

	var (
		webConfig     = webflag.AddFlags(kingpin.CommandLine)
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9150").String()
		worldPath     = kingpin.Flag("mc.world", "").Default("/Users/dirien/Tools/repos/minectl/tmp/world").String()
		rconAddress   = kingpin.Flag("mc.rcon-address", "Address of the Minecraft rcon.").Default(":25575").String()
		rconPassword  = kingpin.Flag("mc.rcon-password", "Password of the Minecraft rcon.").Required().String()
	)

	config := config.NewConfg()
	promlogConfig := &promlog.Config{
		Level: &promlog.AllowedLevel{},
	}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("minecraft_prometheus_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting minecraft_prometheus_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "build", version.BuildContext())

	prometheus.MustRegister(version.NewCollector("minecraft_prometheus_exporter"))
	prometheus.MustRegister(exporter.New(*rconAddress, *rconPassword, *worldPath, logger))

	http.Handle(*config.MetricsPath, promhttp.Handler())
	temp := template.NewIndexTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.Execute(w, config)
	})

	level.Info(logger).Log("msg", "Listening on address", "address", *listenAddress)
	srv := &http.Server{Addr: *listenAddress}
	if err := web.ListenAndServe(srv, *webConfig, logger); err != nil {
		level.Error(logger).Log("msg", "Error running HTTP server", "err", err)
		os.Exit(1)
	}
}
