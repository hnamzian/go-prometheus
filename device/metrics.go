package device

import (
	"prometest/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	devices  prometheus.Gauge
	upgrades *prometheus.CounterVec
	duration *prometheus.HistogramVec
}

func NewMetrics(m *metrics.Metrics) *Metrics {
	dm := &Metrics{
		devices: m.NewGauge(&metrics.GaugeOpts{
			Namespace: "app",
			Name:      "devices",
			Help:      "Number of devices",
		}),
		upgrades: m.NewCounterVec(&metrics.CounterVecOpts{
			Namespace: "app",
			Name:      "upgrades",
			Help:      "Number of upgrades",
			Labels:    []string{"status"},
		}),
		duration: m.NewHistogramVec(&metrics.HistogramVecOpts{
			Namespace: "app",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
			Labels:    []string{"status", "method"},
		}),
	}

	return dm
}

func (m *Metrics) SetDevices(num int) {
	m.devices.Set(float64(num))
}
