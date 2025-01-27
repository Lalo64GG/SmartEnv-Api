package responses

type Response struct {
	Success  bool        `json:"succes"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}
