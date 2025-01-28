package request

type CreateRecordRequest struct {
	Temperature  float64  `json:"temperature" validate:"required"`
	Distance     float64  `json:"distance" validate:"required"`
}