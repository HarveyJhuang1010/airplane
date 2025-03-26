package common

import (
	"airplane/internal/errs"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func NewErrorRepository() *ErrorMappingRepository {
	return &ErrorMappingRepository{
		mapping: make(map[string]*errorMapping),
	}
}

type ErrorMappingRepository struct {
	mx      sync.RWMutex
	mapping map[string]*errorMapping
}

// Mapping initialization function should return a map of error codes and their corresponding errors
func (e *ErrorMappingRepository) Mapping(initialization func() map[errs.Error][]error) *errorMapping {
	key := e.getFuncName()

	e.mx.RLock()
	e.mx.RUnlock()
	if val, ok := e.mapping[key]; ok {
		return val
	}

	e.mx.Lock()
	defer e.mx.Unlock()
	return func() *errorMapping {
		if val, ok := e.mapping[key]; ok {
			return val
		}

		obj := &errorMapping{
			m: make(map[string]errs.Error),
		}
		for code, errs := range initialization() {
			obj.Add(code, errs...)
		}

		e.mapping[key] = obj

		return obj
	}()

}

func (e *ErrorMappingRepository) getFuncName() string {
	pc, _, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name() + ":" + fmt.Sprint(line)
	return funcName[strings.LastIndex(funcName, "/")+1:]
}

type errorMapping struct {
	mx sync.RWMutex
	m  map[string]errs.Error
}

func (e *errorMapping) Add(code errs.Error, errors ...error) *errorMapping {
	e.mx.Lock()
	defer e.mx.Unlock()
	for _, err := range errors {
		customerErr, ok := errs.ParseError(err)
		if ok {
			e.m[customerErr.GetCode()] = code
		} else {
			e.m[err.Error()] = code
		}
	}
	return e
}

func (e *errorMapping) Get(err error) errs.Error {
	e.mx.RLock()
	defer e.mx.RUnlock()

	if customerErr, ok := errs.ParseError(err); ok {
		if err, ok := e.m[customerErr.GetCode()]; ok {
			return err
		} else {
			return errs.ErrUnknown
		}

	} else {
		if err, ok := e.m[err.Error()]; ok {
			return err
		} else {
			return errs.ErrUnknown
		}
	}

	return errs.ErrUnknown
}
