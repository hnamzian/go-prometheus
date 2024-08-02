package device

import (
	"net/http"
	"prometest/metrics"
)

type Module struct{}

func (m *Module) Start(mux *http.ServeMux, metrics *metrics.Metrics) {
	ds := NewDevices()

	dm := NewMetrics(metrics)
	dm.SetDevices(len(ds))

	dh := NewHandlers(dm, ds)
	dh.RegisterHandlers(mux)
}
