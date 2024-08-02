package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metrics struct {
	reg *prometheus.Registry
}

func NewMetrics() *Metrics {
	return &Metrics{
		reg: prometheus.NewRegistry(),
	}
}

func (m *Metrics) RegisterHandler(mux *http.ServeMux) {
	promHandler := promhttp.HandlerFor(m.reg, promhttp.HandlerOpts{})
	mux.Handle("/metrics", promHandler)
}

type GaugeOpts struct {
	Namespace string
	Name      string
	Help      string
}

func (m *Metrics) NewGauge(opts *GaugeOpts) prometheus.Gauge {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	})
	m.reg.MustRegister(g)
	return g
}

type GaugeVecOpts struct {
	Namespace string
	Name      string
	Help      string
	Labels    []string
}

func (m *Metrics) NewGaugeVec(opts *GaugeVecOpts) *prometheus.GaugeVec {
	gv := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	}, opts.Labels)
	m.reg.MustRegister(gv)
	return gv
}

type CounterOpts struct {
	Namespace string
	Name      string
	Help      string
}

func (m *Metrics) NewCounter(opts *CounterOpts) prometheus.Counter {
	c := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	})
	m.reg.MustRegister(c)
	return c
}

type CounterVecOpts struct {
	Namespace string
	Name      string
	Help      string
	Labels    []string
}

func (m *Metrics) NewCounterVec(opts *CounterVecOpts) *prometheus.CounterVec {
	cv := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
	}, opts.Labels)
	m.reg.MustRegister(cv)
	return cv
}

type HistogramOpts struct {
	Namespace string
	Name      string
	Help      string
	Buckets   []float64
}

func (m *Metrics) NewHistogram(opts *HistogramOpts) prometheus.Histogram {
	h := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
		Buckets:   opts.Buckets,
	})
	m.reg.MustRegister(h)
	return h
}

type SummaryOpts struct {
	Namespace string
	Name      string
	Help      string
	Objectives map[float64]float64
}

func (m *Metrics) NewSummary(opts *SummaryOpts) prometheus.Summary {
	s := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  opts.Namespace,
		Name:       opts.Name,
		Help:       opts.Help,
		Objectives: opts.Objectives,
	})
	m.reg.MustRegister(s)
	return s
}

type HistogramVecOpts struct {
	Namespace string
	Name      string
	Help      string
	Buckets   []float64
	Labels    []string
}

func (m *Metrics) NewHistogramVec(opts *HistogramVecOpts) *prometheus.HistogramVec {
	hv := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: opts.Namespace,
		Name:      opts.Name,
		Help:      opts.Help,
		Buckets:   opts.Buckets,
	}, opts.Labels)
	m.reg.MustRegister(hv)
	return hv
}

