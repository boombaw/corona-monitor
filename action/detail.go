package action

import (
	"net/http"

	"github.com/boombaw/corona-monitor/app"
	"github.com/boombaw/corona-monitor/external/rapidapi"
	"github.com/labstack/echo/v4"
)

type InfoCountry struct {
	apiReq rapidapi.API
}

func (i *InfoCountry) Info(c echo.Context, url string) (rapidapi.RapidApiResponse, error) {
	result, err := i.apiReq.Request(c, url)
	if err != nil {
		return rapidapi.RapidApiResponse{}, err
	}
	return result, nil
}

//InfoByCountry info by country
func InfoByCountry(c echo.Context) (err error) {
	url := app.SetURL("latest_stat_by_country.php?country=" + c.Param("country"))

	i := InfoCountry{}

	result, err := i.Info(c, url)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, result)
}
