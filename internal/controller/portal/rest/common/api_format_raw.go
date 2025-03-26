package common

import (
	"airplane/internal/components/apis"
)

var (
	RawData    = rawData{}
	RawDataOtp = rawDataImpl{}
)

type rawData struct {
}

// New
// param opts use func RawDataOtp
func (rawData) New(otps ...RawDataOtpFunc) *apis.RawData {
	result := apis.NewRawData()
	for _, opt := range otps {
		opt(result)
	}

	return result
}

type RawDataOtpFunc func(*apis.RawData)

type rawDataImpl struct{}

func (r *rawDataImpl) WithTraceNamed(loggerNamed string) RawDataOtpFunc {
	return func(r *apis.RawData) {
		r.Response.TraceNamed = loggerNamed
	}
}
