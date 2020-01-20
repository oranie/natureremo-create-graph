package pkg

type Event struct {
	Value     float64 `json:"val"`
	CreatedAt string  `json:"created_at"`
}

type NewestEvents struct {
	Humidity     Event `json:"hu"`
	Illumination Event `json:"il"`
	Movement     Event `json:"mo"`
	Temperature  Event `json:"te"`
}

type User struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	Superuser bool   `json:"superuser"`
}

type Device struct {
	Name              string       `json:"name"`
	Id                string       `json:"id"`
	CreatedAt         string       `json:"created_at"`
	UpdatedAt         string       `json:"updated_at"`
	MacAddress        string       `json:"mac_address"`
	SerialNumber      string       `json:"serial_number"`
	FirmwareVersion   string       `json:"firmware_version"`
	TemperatureOffset int          `json:"temperature_offset"`
	HumidityOffset    int          `json:"humidity_offset"`
	Users             []User       `json:"users"`
	NewestEvents      NewestEvents `json:"newest_events"`
}
