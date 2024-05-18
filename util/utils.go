package util

import "go.uber.org/zap"

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

func ZapInfo(msg string, tag string, service string) {
	zap.L().Info(
		msg,
		zap.String("tag", tag),
		zap.String("service", service))
}

func ZapError(err error, tag string, service string) {
	zap.L().Error(
		err.Error(),
		zap.String("tag", tag),
		zap.String("service", service))
}
