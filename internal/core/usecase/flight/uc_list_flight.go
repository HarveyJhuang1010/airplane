package flight

import (
	"airplane/internal/domain/entities/po"
	"airplane/internal/errs"
	"context"
	"github.com/jinzhu/copier"
	"sync"

	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
)

func newListFlight(in dependence) *ListFlight {
	return &ListFlight{
		in: in,
	}
}

type ListFlight struct {
	in dependence
}

func (uc *ListFlight) ListFlight(ctx context.Context, cond *bo.ListFlightCond) ([]*bo.FlightDetail, *bo.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	poList, pagePo, err := uc.getFlights(ctx, cond)
	if err != nil {
		return nil, nil, err
	}

	resultPagination := &bo.Pagination{}
	if err = copier.Copy(resultPagination, pagePo); err != nil {
		return nil, nil, errs.ErrInvalidPayload.Trace(err)
	}

	resultItems := make([]*bo.FlightDetail, 0, len(poList))
	if err = copier.Copy(&resultItems, poList); err != nil {
		return nil, nil, errs.ErrInvalidPayload.Trace(err)
	}

	return resultItems, resultPagination, nil
}

func (uc *ListFlight) getFlights(ctx context.Context, cond *bo.ListFlightCond) ([]*po.Flight, *po.Pagination, error) {
	defer timelogger.LogTime(ctx)()

	var (
		poCond           *po.FlightListCond
		wg               sync.WaitGroup
		errList, errPage error
		poList           []*po.Flight
		pagePo           *po.Pagination
	)
	if cond != nil {
		poCond := &po.FlightListCond{
			PreloadCabinClasses:  true,
			DepartureAirport:     cond.DepartureAirport,
			ArrivalAirport:       cond.ArrivalAirport,
			DepartureTimeStartAt: cond.DepartureTimeStartAt,
			DepartureTimeEndAt:   cond.DepartureTimeEndAt,
			CanSell:              true,
		}
		if cond.Pager != nil {
			poCond.Pager = &po.Pager{
				Index: cond.Pager.Index,
				Size:  cond.Pager.Size,
			}
		}
	}

	wg.Add(2)

	go func() {
		defer wg.Done()
		poList, errList = uc.in.DBRepository.ReadOnly().FlightDAO().List(ctx, poCond)
	}()

	go func() {
		defer wg.Done()
		pagePo, errPage = uc.in.DBRepository.ReadOnly().FlightDAO().ListPager(ctx, poCond)
	}()

	wg.Wait()

	if errList != nil {
		return nil, nil, errList
	}

	if errPage != nil {
		return nil, nil, errPage
	}

	return poList, pagePo, nil
}
