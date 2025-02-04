package entities

type Record struct {
	ID          int64    `json:"id"`
	Temperature float64  `json:"temperature"`
	Humidity 	float64 `json:"humidity"`
	Gas_level 	float64 `json:"gas_level"`
}