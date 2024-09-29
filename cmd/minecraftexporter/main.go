package minecraftexporter

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/minecraft-exporter/pkg/config"
	"github.com/minecraft-exporter/pkg/exporter"
	"github.com/minecraft-exporter/pkg/template"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	v2 "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/promslog"
	"github.com/prometheus/common/promslog/flag"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
)

func Run() {
	config := config.NewConfg()
	promlogConfig := &promslog.Config{
		Level: &promslog.AllowedLevel{},
	}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("minecraft_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promslog.New(promlogConfig)

	config.LoadFile()

	fmt.Print(`_  _ _ _  _ ____ ____ ____ ____ ____ ___    ____ _  _ ___  ____ ____ ___ ____ ____ 
|\/| | |\ | |___ |    |__/ |__| |___  |  __ |___  \/  |__] |  | |__/  |  |___ |__/ 
|  | | | \| |___ |___ |  \ |  | |     |     |___ _/\_ |    |__| |  \  |  |___ |  \ 
`)

	logger.Info("Starting minecraft_exporter", "version", version.Info())
	logger.Info("Build context", "build", version.BuildContext())

	prometheus.MustRegister(v2.NewCollector("minecraft_exporter"))
	exporter, err := exporter.New(*config.RconAddress, *config.RconPassword, *config.WorldPath, *config.NameSource, *config.ModServerStats, config.DisabledMetrics, logger)
	if err != nil {
		logger.Error("Failed to create exporter", "err", err)
	}
	prometheus.MustRegister(exporter)

	logger.Info("Disabling collection of exporter metrics (like go_*)", "value", config.DisableExporterMetrics)
	if *config.DisableExporterMetrics {
		prometheus.Unregister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
		prometheus.Unregister(collectors.NewGoCollector())
	}

	http.Handle(*config.MetricsPath, promhttp.Handler())
	template := template.NewIndexTemplate()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		err := template.Execute(w, config)
		if err != nil {
			logger.Error("Error executing template", "err", err)
		}
	})

	go func() {
		logger.Info("Listening on address", "address", (*config.FlagConfig.WebListenAddresses)[0])
		srv := &http.Server{
			Addr:              (*config.FlagConfig.WebListenAddresses)[0],
			ReadHeaderTimeout: 60 * time.Second,
		}
		if err := web.ListenAndServe(srv, config.FlagConfig, logger); err != nil {
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
		exporter.StopRCON()
		close(done)
	}()

	<-done
	logger.Info("Shutting down...")
}
