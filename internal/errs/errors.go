package errs

import (
	"airplane/internal/components/errortool"
)

type Error = errortool.Error

var TraceWrap = errortool.TraceWrap

func ParseError(err error) (Error, bool) {
	return errortool.Parse(err)
}

// sys
var (
	sysErrGroup    = errortool.Codes.CustomGroup("SYS")
	ErrParseFailed = sysErrGroup.Error("parse failed")
	ErrUnknown     = sysErrGroup.Error("unknown error")
)

// database
var (
	rdbErrGroup      = errortool.Codes.CustomGroup("RDB")
	ErrDBQueryFailed = rdbErrGroup.Error("db query failed")
)

// redis
var (
	redisErrGroup           = errortool.Codes.CustomGroup("RDS")
	ErrRedisNotFound        = redisErrGroup.Error("redis cache not found")
	ErrRedisTxFailed        = redisErrGroup.Error("redis tx failed")
	ErrRedisOperationFailed = redisErrGroup.Error("redis operation failed")
)

// mq
var (
	mqErrGroup  = errortool.Codes.CustomGroup("MQ")
	ErrMQFailed = mqErrGroup.Error("mq failed")
)

// common
var (
	// common
	commonErrGroup      = errortool.Codes.CustomGroup("COM")
	ErrInvalidPayload   = commonErrGroup.Error("invalid payload")
	ErrInvalidParameter = commonErrGroup.Error("invalid parameter")
	ErrRecordNotFound   = commonErrGroup.Error("record not found")
	ErrDuplicateRecord  = commonErrGroup.Error("duplicate record")
	ErrStatusNotMatch   = commonErrGroup.Error("status not match")
)

// flight
var (
	flightErrGroup        = errortool.Codes.CustomGroup("FLT")
	ErrFlightSoldOut      = flightErrGroup.Error("flight sold out")
	ErrFlightNotAvailable = flightErrGroup.Error("flight not available")
	ErrSeatNotAvailable   = flightErrGroup.Error("seat not available")
)
