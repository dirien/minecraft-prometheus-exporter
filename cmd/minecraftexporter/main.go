package minecraftexporter

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kong"
	"github.com/minecraft-exporter/pkg/config"
	"github.com/minecraft-exporter/pkg/exporter"
	"github.com/minecraft-exporter/pkg/template"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	v2 "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promslog"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
)

func Run() {
	cfg := &config.Config{}
	kong.Parse(cfg,
		kong.Name("minecraft_exporter"),
		kong.Description("Prometheus exporter for Minecraft Java Edition servers."),
		kong.UsageOnError(),
		kong.Vars{"version": version.Print("minecraft_exporter")},
	)

	cfg.LoadFile()

	promlogConfig := &promslog.Config{
		Level:  promslog.NewLevel(),
		Format: promslog.NewFormat(),
	}
	if err := promlogConfig.Level.Set(cfg.LogLevel); err != nil {
		fmt.Fprintf(os.Stderr, "invalid --log.level %q: %v\n", cfg.LogLevel, err)
		os.Exit(2)
	}
	if err := promlogConfig.Format.Set(cfg.LogFormat); err != nil {
		fmt.Fprintf(os.Stderr, "invalid --log.format %q: %v\n", cfg.LogFormat, err)
		os.Exit(2)
	}
	logger := promslog.New(promlogConfig)

	fmt.Print(`_  _ _ _  _ ____ ____ ____ ____ ____ ___    ____ _  _ ___  ____ ____ ___ ____ ____
|\/| | |\ | |___ |    |__/ |__| |___  |  __ |___  \/  |__] |  | |__/  |  |___ |__/
|  | | | \| |___ |___ |  \ |  | |     |     |___ _/\_ |    |__| |  \  |  |___ |  \
`)

	logger.Info("Starting minecraft_exporter", "version", version.Info())
	logger.Info("Build context", "build", version.BuildContext())

	prometheus.MustRegister(v2.NewCollector("minecraft_exporter"))
	exp, err := exporter.New(cfg.RconAddress, cfg.RconPassword, cfg.WorldPath, cfg.NameSource, cfg.ModServerStats, cfg.DisabledMetrics, logger)
	if err != nil {
		logger.Error("Failed to create exporter", "err", err)
	}
	prometheus.MustRegister(exp)

	logger.Info("Disabling collection of exporter metrics (like go_*)", "value", cfg.DisableExporterMetrics)
	if cfg.DisableExporterMetrics {
		prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		prometheus.Unregister(collectors.NewGoCollector())
	}

	http.Handle(cfg.MetricsPath, promhttp.Handler())
	tpl := template.NewIndexTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if err := tpl.Execute(w, cfg); err != nil {
			logger.Error("Error executing template", "err", err)
		}
	})

	flagConfig := cfg.WebFlagConfig()
	go func() {
		logger.Info("Listening on address", "address", (*flagConfig.WebListenAddresses)[0])
		srv := &http.Server{
			Addr:              (*flagConfig.WebListenAddresses)[0],
			ReadHeaderTimeout: 60 * time.Second,
		}
		if err := web.ListenAndServe(srv, flagConfig, logger); err != nil {
			logger.Error("Error running HTTP server", "err", err)
			os.Exit(1)
		}
	}()
	done := make(chan struct{})
	go func() {
		logger.Info("Listening signals...")
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		exp.StopRCON()
		close(done)
	}()

	<-done
	logger.Info("Shutting down...")
}
