package request

type CreateRecordRequest struct {
	Temperature  float64   `json:"temperature" validate:"required"`
	Humidity 	 float64   `json:"humidity" validate:"required"`
	Gas_level    float64   `json:"gas_level" validate:"required"`
}