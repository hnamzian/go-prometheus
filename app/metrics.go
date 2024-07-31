package app

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	info prometheus.GaugeVec
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		info: *prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "app_info",
			Help:      "Application information",
		}, []string{"version"}),
	}

	reg.MustRegister(m.info)

	return m
}

func (m *Metrics) SetVersion(ver string) {
	m.info.WithLabelValues(ver).Set(1)
}
