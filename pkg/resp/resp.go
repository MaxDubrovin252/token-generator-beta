package resp

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "Ok"
	StatusError = "Error"
)

func OK() *Response {
	return &Response{
		Status: StatusOK,
	}
}

func Error(message string) *Response {
	return &Response{
		Status: StatusError,
		Error:  message,
	}
}
