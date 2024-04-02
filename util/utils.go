package util

func ErrorResponse(err error) Response {
	return Response{
		Err: err.Error(),
	}
}

type Response struct {
	Results any    `json:"results,omitempty"`
	Count   int    `json:"count,omitempty"`
	Err     string `json:"error,omitempty"`
}
