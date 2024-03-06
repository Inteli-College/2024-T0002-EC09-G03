package entity

type DataSensor struct {
	Measurament float64 `json:"measurament"`
	Unit        string  `json:"unit"`
	Material    string  `json:"material"`
}
