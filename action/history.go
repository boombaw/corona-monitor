package action

import (
	"net/http"

	"github.com/boombaw/corona-monitor/app"
	"github.com/boombaw/corona-monitor/external/rapidapi"
	"github.com/labstack/echo/v4"
)

type History struct {
	apiReq rapidapi.API
}

func (h *History) Info(c echo.Context, url string) (rapidapi.RapidApiResponse, error) {
	result, err := h.apiReq.Request(c, url)
	if err != nil {
		return rapidapi.RapidApiResponse{}, err
	}
	return result, nil
}

//HistoryByCountry history info by country
func HistoryByCountry(c echo.Context) (err error) {
	url := app.SetURL("cases_by_particular_country.php?country=" + c.Param("country"))

	i := History{}

	result, err := i.Info(c, url)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, result)
}
