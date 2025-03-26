package flight

import (
	"airplane/internal/components/apis"
	"airplane/internal/domain/entities/bo"
	"airplane/internal/tools/timelogger"
	"github.com/gin-gonic/gin"
)

func newFlight(in dependence) *Flight {
	return &Flight{
		in:       in,
		Response: in.Response,
	}
}

type Flight struct {
	in       dependence
	Response apis.IResponse
}

// ListFlight
// @Summary  get flight list
// @Description get flight list
// @Tags 		Flight
// @Accept 		json
// @Produce		json
// @Param 		params query FlightListRequest false "Query Params"
// @Success 	200 {object} apis.StandardListResponse{data=[]FlightResponse}
// @Failure 	400 {object} apis.StandardResponse{error=apis.StandardError}
// @Failure 	500 {object} apis.StandardResponse{error=apis.StandardError}
// @Router 		/flight [get]
func (ctrl *Flight) ListFlight(ctx *gin.Context) {
	var err error
	defer func() {
		timelogger.LogTime(ctx)()
		if err != nil {
			ctrl.Response.Data(ctx, ctrl.in.StandardData().BadRequest(nil, err))
		}
	}()

	dto, err := apis.RequestParser[FlightListRequest]().Query().Bind(ctx)
	if err != nil {
		return
	}

	cond, err := apis.Convert[bo.ListFlightCond](ctx, dto)
	if err != nil {
		return
	}

	result, pagination, err := ctrl.in.Flight.ListFlight.ListFlight(ctx, cond)
	if err != nil {
		return
	}

	resp, err := apis.Convert[[]FlightResponse](ctx, result)
	if err != nil {
		return
	}

	ctrl.Response.Data(ctx, ctrl.in.StandardData().OK(resp).SetPagination(pagination))
}
