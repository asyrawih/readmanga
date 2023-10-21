package model

type Response struct {
	ErrorCode int    `json:"error_code,omitempty"`
	Message   string `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
}

func NewResponse() *Response {
	return new(Response)
}

func (r *Response) SetErrorCode(code int) *Response {
	r.ErrorCode = code
	return r
}

func (r *Response) SetMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}
