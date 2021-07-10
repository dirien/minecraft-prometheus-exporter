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
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

func Run() {

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
	prometheus.MustRegister(exporter.New(*config.RconAddress, *config.RconPassword, *config.WorldPath, logger))

	http.Handle(*config.MetricsPath, promhttp.Handler())
	template := template.NewIndexTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Execute(w, config)
	})

	level.Info(logger).Log("msg", "Listening on address", "address", *config.ListenAddress)
	srv := &http.Server{Addr: *config.ListenAddress}
	if err := web.ListenAndServe(srv, *config.WebConfig, logger); err != nil {
		level.Error(logger).Log("msg", "Error running HTTP server", "err", err)
		os.Exit(1)
	}
}
