package models

//Satellites struct
type Satellites struct {
	SatellitesUnit []SatelliteUnit
}

//SatelliteUnit struct
type SatelliteUnit struct {
	Name     string   `json:"name"`
	Distance float64  `json:"distance"`
	Message  []string `json:"message"`
}

//ShipCoordinates struct
type ShipCoordinates struct {
	X string `json:"x"`
	Y string `json:"y"`
}

//DBconsulting
type DBconsult struct {
	Name     string
	Distance float64
	Message  string
}
