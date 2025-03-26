package bo

import "airplane/internal/domain/entities/po"

type FetchDBCond[T any] struct {
	Pager *po.Pager
	Cond  T
}
