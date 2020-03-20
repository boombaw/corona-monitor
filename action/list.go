package action

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/boombaw/corona-monitor/app"
	"github.com/boombaw/corona-monitor/external/rapidapi"
)

type Countries struct {
	apiReq rapidapi.API
}

func (cr *Countries) Info(c echo.Context, url string) (rapidapi.RapidApiResponse, error) {
	result, err := cr.apiReq.Request(c, url)
	if err != nil {
		return rapidapi.RapidApiResponse{}, err
	}
	return result, nil
}

//AllCountry for monitoring cases all country
func AllCountry(c echo.Context) (err error) {
	url := app.SetURL("cases_by_country.php")
	
	i := InfoCountry{}
	result, err := i.Info(c, url)
	
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, result)
}
