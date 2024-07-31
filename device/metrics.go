package device

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	devices  prometheus.Gauge
	upgrades prometheus.CounterVec
	duration prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "devices",
			Help:      "Number of devices",
		}),
		upgrades: *prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "app",
			Name:      "upgrades",
			Help:      "Number of upgrades",
		}, []string{"type"}),
		duration: *prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "app",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
	}

	reg.MustRegister(m.devices, m.upgrades, m.duration)

	return m
}

func (m *Metrics) SetDevices(num int) {
	m.devices.Set(float64(num))
}