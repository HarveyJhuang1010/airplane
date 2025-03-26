package common

import (
	"airplane/internal/components/apis"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	V1DataOpt = v1DataOptImpl{}
	V1Data    = v1Data{}
)

type v1Data struct{}

// New
// param opts use func GetApiV1DataOpt
func (v1Data) New(opts ...V1DataOptFunc) *ApiV1Data {
	result := &ApiV1Data{}
	for _, opt := range opts {
		opt(result)
	}

	return result
}

type V1DataOptFunc func(*ApiV1Data)

type v1DataOptImpl struct{}

func (r *v1DataOptImpl) WithTraceNamed(loggerNamed string) V1DataOptFunc {
	return func(r *ApiV1Data) {
		r.Response.TraceNamed = loggerNamed
	}
}

type ApiV1Data struct {
	apis.Response

	data any
}

func (r *ApiV1Data) Format(ctx *gin.Context, meta *apis.Meta) (*apis.Response, error) {
	var (
		resp = r.data
		buf  []byte
	)

	if r.data == nil {
		buf = []byte{}
	} else {
		b, err := json.Marshal(resp)
		if err != nil {
			return nil, err
		}
		buf = b
	}

	r.Response.Data = buf
	r.Response.ContentType = apis.ContentTypeJSON
	return &r.Response, nil
}

func (r *ApiV1Data) Set(status int, data any) *ApiV1Data {
	r.Response.Status = status
	r.data = data
	return r
}

func (r *ApiV1Data) WithTraceNamed(loggerNamed string) *ApiV1Data {
	r.Response.TraceNamed = loggerNamed
	return r
}

func (r *ApiV1Data) OK(data any) *ApiV1Data {
	return r.Set(http.StatusOK, data)
}

func (r *ApiV1Data) Created(data any) *ApiV1Data {
	return r.Set(http.StatusCreated, data)
}

func (r *ApiV1Data) BadRequest(errorCode Code, err error) *ApiV1Data {
	r.Response.Error = err
	return r.Set(http.StatusBadRequest, ErrorResponse{
		Code: int(errorCode),
		Data: err,
		Error: func() string {
			if err != nil {
				return err.Error()
			} else {
				return ""
			}
		}(),
	})
}
