package entities

type Record struct {
	ID          int64    `json:"id"`
	Temperature float64  `json:"temperature"`
	Distance 	float64  `json:"distance"`
}