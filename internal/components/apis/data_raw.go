package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRawData() *RawData {
	return &RawData{
		Response: Response{
			Headers: make(map[string]string),
		},
	}
}

type RawData struct {
	Response
}

func (s *RawData) Format(ctx *gin.Context, meta *Meta) (*Response, error) {
	if s.ContentType == "" {
		s.ContentType = ctx.ContentType()
	}

	return &s.Response, nil
}

func (s *RawData) SetStatus(statusCode int) *RawData {
	s.Status = statusCode
	return s
}

func (s *RawData) SetContentType(contentType string) *RawData {
	s.ContentType = contentType
	return s
}

func (s *RawData) SetHeaders(headers map[string]string) *RawData {
	for k, v := range headers {
		s.Headers[k] = v
	}
	return s
}

func (s *RawData) SetHeader(key, val string) *RawData {
	s.Headers[key] = val
	return s
}

func (s *RawData) SetData(data []byte) *RawData {
	s.Data = data
	return s
}

func (s *RawData) SetError(err error) *RawData {
	s.Error = err
	return s
}

func (s *RawData) Set(statusCode int, data []byte) *RawData {
	return s.SetStatus(statusCode).SetData(data)
}

func (s *RawData) WithTraceNamed(loggerNamed string) *RawData {
	s.TraceNamed = loggerNamed
	return s
}

func (s *RawData) OK(data []byte) *RawData {
	return s.Set(http.StatusOK, data)
}

func (s *RawData) File(data []byte, fileName, contentType string) *RawData {
	return s.Set(http.StatusOK, data).SetHeaders(map[string]string{
		"Content-Disposition": "attachment; filename=" + fileName,
	}).SetContentType(contentType)
}

func (s *RawData) BadRequest(data []byte) *RawData {
	return s.Set(http.StatusBadRequest, data)
}
