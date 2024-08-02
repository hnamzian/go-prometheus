package main

import (
	"net/http"

	"prometest/app"
	dev "prometest/device"
	"prometest/metrics"
)

var version string

func init() {
	version = "1.0.0"
}

func main() {
	mux := http.NewServeMux()

	m := metrics.NewMetrics()
	m.RegisterHandler(mux)

	am := app.NewMetrics(m)
	am.SetVersion(version)

	dm := dev.Module{}
	dm.Start(mux, m)

	s := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	s.ListenAndServe()
}
