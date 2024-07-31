package device

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type Module struct {}

func (m *Module) Start(mux *http.ServeMux, reg prometheus.Registerer) {
	ds := NewDevices()

	dm := NewMetrics(reg)
	dm.SetDevices(len(ds))

	dh := NewHandlers(dm, ds)
	dh.RegisterHandlers(mux)
}