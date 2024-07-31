package device

type Device struct {
	ID       string
	Mac      string
	Firmware string
}

type Devices []Device

func NewDevices() Devices {
	return Devices{
		{"1", "00:11:22:33:44:55", "1.0.0"},
		{"2", "00:11:22:33:44:56", "1.0.0"},
		{"3", "00:11:22:33:44:57", "1.0.0"},
		{"4", "00:11:22:33:44:58", "1.0.0"},
		{"5", "00:11:22:33:44:59", "1.0.0"},
	}
}

func (ds Devices) GetDevice(id string) *Device {
	for _, d := range ds {
		if d.ID == id {
			return &d
		}
	}

	return nil
}

func (ds Devices) AddDevice(d Device) {
	ds = append(ds, d)
}

func (ds Devices) UpdateDeviceFirmware(id string, d Device) {
	for i, dev := range ds {
		if dev.ID == id {
			ds[i].Firmware = d.Firmware
			break
		}
	}
}
