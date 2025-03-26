package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StandardData struct {
	Response `json:"-"`

	errorHandle func(err error) any

	Code       string
	Data       any
	Pagination *Pagination
}

// Format
// ContentType default is application/json
func (s *StandardData) Format(ctx *gin.Context, meta *Meta) (*Response, error) {
	var dataFunc = func(data any) any {
		if data != nil {
			return data
		}
		return EmptyData{}
	}

	data := func() any {
		standardResponse := StandardResponse{
			Meta: meta,
			Code: s.Code,
			Data: dataFunc(s.Data),
		}
		if s.Response.Error != nil {
			if s.errorHandle != nil {
				standardResponse.Error = s.errorHandle(s.Response.Error)
			} else {
				standardResponse.Error = s.Response.Error
			}
		}

		if s.Pagination != nil {
			return StandardListResponse{
				StandardResponse: standardResponse,
				Pagination:       *s.Pagination,
			}
		}

		return standardResponse
	}()

	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	s.Response.ContentType = ContentTypeJSON
	s.Response.Data = buf

	return &s.Response, nil
}

func (s *StandardData) SetStatus(status int) *StandardData {
	s.Response.Status = status
	return s
}

func (s *StandardData) SetCode(code string) *StandardData {
	s.Code = code
	return s
}

func (s *StandardData) SetData(data any) *StandardData {
	s.Data = data
	return s
}

func (s *StandardData) SetError(err error) *StandardData {
	s.Response.Error = err
	return s
}

func (s *StandardData) Set(statusCode int, code string, data any) *StandardData {
	s.SetStatus(statusCode).SetCode(code).SetData(data)
	return s
}

func (s *StandardData) SetPagination(pagination *Pagination) *StandardData {
	s.Pagination = pagination
	return s
}

func (s *StandardData) WithTraceNamed(loggerNamed string) *StandardData {
	s.TraceNamed = loggerNamed
	return s
}

func (s *StandardData) WithErrorHandle(handle func(err error) any) *StandardData {
	s.errorHandle = handle
	return s
}

func (s *StandardData) OK(data any) *StandardData {
	return s.Set(http.StatusOK, responseCode_Success, data)
}

func (s *StandardData) Created(data any) *StandardData {
	return s.Set(http.StatusCreated, responseCode_Success, data)
}

func (s *StandardData) NoContent() *StandardData {
	return s.Set(http.StatusNoContent, responseCode_Success, EmptyData{})
}

func (s *StandardData) BadRequest(data any, err error) *StandardData {
	var dataFunc = func(data any) any {
		if data != nil {
			return data
		}
		return EmptyData{}
	}

	return s.Set(http.StatusBadRequest, responseCode_Fail, dataFunc(data)).SetError(err)
}

func (s *StandardData) BadRequestWithCode(data any, err error, code string) *StandardData {
	return s.Set(http.StatusBadRequest, responseCode_Fail, data).SetError(err).SetCode(code)
}

func (s *StandardData) Unauthorized() *StandardData {
	return s.Set(http.StatusUnauthorized, responseCode_Fail, nil)
}
