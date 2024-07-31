package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"prometest/app"
	dev "prometest/device"
)

var version string

func init() {
	version = "1.0.0"
}

func main() {
	mux := http.NewServeMux()

	reg := prometheus.NewRegistry()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promHandler)

	am := app.NewMetrics(reg)
	am.SetVersion(version)

	dm := dev.Module{}
	dm.Start(mux, reg)

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}
