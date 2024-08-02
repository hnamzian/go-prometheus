package app

import (
	"prometest/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	info *prometheus.GaugeVec
}

func NewMetrics(m *metrics.Metrics) *Metrics {
	ma := &Metrics{
		info: m.NewGaugeVec(&metrics.GaugeVecOpts{
			Namespace: "app",
			Name:      "app_info",
			Help:      "Application information",
			Labels:    []string{"version"},
		}),
	}
	return ma
}

func (m *Metrics) SetVersion(ver string) {
	m.info.WithLabelValues(ver).Set(1)
}
