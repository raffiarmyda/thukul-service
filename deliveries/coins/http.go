package coins

import (
	"aprian1337/thukul-service/business/coins"
	"aprian1337/thukul-service/deliveries"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	CoinsUsecase coins.Usecase
}

func NewCoinsController(uc coins.Usecase) *Controller {
	return &Controller{
		CoinsUsecase: uc,
	}
}

func (ctrl *Controller) GetBySymbol(c echo.Context) error {
	symbol := c.QueryParam("symbol")
	ctx := c.Request().Context()
	data, err := ctrl.CoinsUsecase.GetBySymbol(ctx, symbol)
	if err != nil {
		return deliveries.NewErrorResponse(c, http.StatusForbidden, err)
	}
	return deliveries.NewSuccessResponse(c, data)
}
