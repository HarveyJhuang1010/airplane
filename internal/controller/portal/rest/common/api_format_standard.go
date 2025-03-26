package common

import (
	"airplane/internal/components/apis"
	"airplane/internal/components/errortool"
)

var (
	StandardData    = standardData{}
	StandardDataOtp = standardDataImpl{}
)

type standardData struct {
}

// New
// param opts use func StandardDataOtp
func (standardData) New(otps ...StandardDataOtpFunc) *apis.StandardData {
	result := &apis.StandardData{}
	otps = append(otps, StandardDataOtp.WithErrorHandle())
	for _, opt := range otps {
		opt(result)
	}

	return result
}

type StandardDataOtpFunc func(*apis.StandardData)

type standardDataImpl struct{}

func (r *standardDataImpl) WithTraceNamed(loggerNamed string) StandardDataOtpFunc {
	return func(r *apis.StandardData) {
		r.Response.TraceNamed = loggerNamed
	}
}

func (r *standardDataImpl) WithErrorHandle() StandardDataOtpFunc {
	return func(r *apis.StandardData) {
		var handle = func(err error) any {
			if err == nil {
				return nil
			}

			result := &apis.StandardError{}
			v, ok := errortool.Parse(err)
			if !ok {
				result.Message = err.Error()
				return result
			}

			result.Code = v.GetCode()
			result.Message = v.GetMessage()

			return result
		}

		r.WithErrorHandle(handle)
	}
}
