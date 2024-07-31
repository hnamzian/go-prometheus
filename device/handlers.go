package device

import (
	"encoding/json"
	"net/http"
	"prometest/utils"
	"strings"
	"time"
)

type Handlers struct {
	metrics *Metrics
	devices Devices
}

func NewHandlers(m *Metrics, ds Devices) *Handlers {
	return &Handlers{
		metrics: m,
		devices: ds,
	}
}

func (h *Handlers) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/devices", h.registerDevicesHandler())
	mux.HandleFunc("/devices/", h.registerManageDevicesHandler())
}

func (h *Handlers) registerDevicesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getDevices(w, r)
		case http.MethodPost:
			h.createDevice(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *Handlers) registerManageDevicesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			h.upgradeDevice(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (h *Handlers) getDevices(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	utils.Sleep(200)
	duration := time.Since(now).Seconds()
	h.metrics.duration.WithLabelValues("200", "GET").Observe(duration)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.devices)
}

func (h *Handlers) createDevice(w http.ResponseWriter, r *http.Request) {
	var d Device
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.devices = append(h.devices, d)

	h.metrics.devices.Set(float64(len(h.devices)))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

func (h *Handlers) upgradeDevice(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.URL.Path, "/devices")

	var d Device
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, dev := range h.devices {
		if dev.ID == id {
			h.devices[i].Firmware = d.Firmware

			h.metrics.upgrades.WithLabelValues("router").Inc()

			break
		}
	}
}
