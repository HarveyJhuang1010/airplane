package errortool

import (
	"errors"
	"fmt"
	"path"
	"runtime"
)

type errorCode string

type Error interface {
	Error() string
	GetCode() string
	GetMessage() string
	GetData() any
	Trace(data ...any) Error
	TraceWrap(err error, data ...any) Error
	Is(err error) bool
}

type customerError struct {
	code    errorCode
	message string
	data    any
	stack   string

	cause error
}

func (e *customerError) Error() string {
	if e.cause != nil {
		return e.Format() + "\ncaused by: " + e.cause.Error()
	}

	return e.Format()
}

func (e *customerError) Format() string {
	if e.stack == "" {
		return fmt.Sprintf("%s: %s",
			e.code, e.message,
		)
	}
	return fmt.Sprintf("%s: %s >> %s %v",
		e.code, e.message, e.stack, e.data,
	)
}

func (e *customerError) GetCode() string {
	return string(e.code)
}

func (e *customerError) GetMessage() string {
	return e.message
}

func (e *customerError) GetData() any {
	return e.data
}

func (e *customerError) Trace(data ...any) Error {
	return e.traceWrap(e.cause, data...)
}

func (e *customerError) TraceWrap(err error, data ...any) Error {
	return e.traceWrap(err, data...)
}

func (e *customerError) traceWrap(err error, data ...any) Error {
	return &customerError{
		code:    e.code,
		message: e.message,
		data:    data,
		cause:   err,
		stack:   caller(3),
	}
}

func (e *customerError) Is(err error) bool {
	target := &customerError{}
	if errors.As(err, &target) {
		return target.code == e.code
	}
	return e.Error() == err.Error()
}

func Parse(err error) (Error, bool) {
	if err == nil {
		return nil, false
	}

	var e *customerError
	if errors.As(err, &e) {
		return e, true
	}

	return nil, false
}

func TraceWrap(err error, data ...any) Error {
	if err == nil {
		return nil
	}

	var e *customerError
	if errors.As(err, &e) {
		return e.traceWrap(err, data...)
	}

	return &customerError{
		code:    errorCode(err.Error()),
		message: "",
		data:    data,
		cause:   err,
		stack:   caller(2),
	}
}

func caller(skip int) string {
	pc, file, line, _ := runtime.Caller(skip)
	fn := runtime.FuncForPC(pc)
	return fmt.Sprintf("%s:%d [%s]", file, line, path.Base(fn.Name()))
}
