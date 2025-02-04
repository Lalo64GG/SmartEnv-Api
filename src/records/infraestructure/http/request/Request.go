package request

type CreateRecordRequest struct {
	Temperature  float64   `json:"temperature" validate:"required"`
	Humedity 	 float64   `json:"humedity" validate:"required"`
	Gas_level    float64   `json:"gas" validate:"required"`
}