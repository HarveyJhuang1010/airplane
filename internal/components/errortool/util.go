package errortool

import (
	"log"
	"sync"
	"sync/atomic"
)

// newCodeRepository: registration error code, and check error code is unique.
func newCodeRepository() iCodeRepository {
	return &codeRepository{}
}

type iCodeRepository interface {
	Add(code errorCode, err *customerError)
	Get(code errorCode) (*customerError, bool)
	Keys() []errorCode
}

type codeRepository struct {
	m sync.Map
}

func (c *codeRepository) Add(code errorCode, err *customerError) {
	if _, ok := c.m.LoadOrStore(code, err); ok {
		log.Panicf("error code duplicate definition, code: %s", code)
	}
}

func (c *codeRepository) Get(code errorCode) (*customerError, bool) {
	val, ok := c.m.Load(code)
	return val.(*customerError), ok
}

func (c *codeRepository) Keys() []errorCode {
	result := make([]errorCode, 0)

	c.m.Range(func(key, value interface{}) bool {
		result = append(result, key.(errorCode))
		return true
	})

	return result
}

func newSequence(begin int, max int) *sequence {
	s := &sequence{
		max: max,
	}

	s.now.Store(int64(begin))

	return s
}

type sequence struct {
	now atomic.Int64
	max int
}

func (s *sequence) Next() int {
	now := int(s.now.Add(1))
	if now > s.max {
		panic("max sequence")
	}
	return now
}
