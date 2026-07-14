package helper

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//successres
func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Status:  200,
		Message: "Success",
		Data:    data,
	}
}

//err response
func ErrorResponse(status int, message string) Response {
	return Response{
		Success: false,
		Status:  status,
		Message: message,
	}
}