package action

import (
	"net/http"

	"github.com/boombaw/corona-monitor/app"
	"github.com/boombaw/corona-monitor/external/rapidapi"
	"github.com/labstack/echo/v4"
)

type Affected struct {
	apiReq rapidapi.API
}

func (a *Affected) Info(c echo.Context, url string) (rapidapi.RapidApiResponse, error) {
	result, err := a.apiReq.Request(c, url)
	if err != nil {
		return rapidapi.RapidApiResponse{}, err
	}
	return result, nil
}

// AffectedCountry for info counting affected country
func AffectedCountry(c echo.Context) (err error) {
	url := app.SetURL("affected.php")

	a := Affected{}

	result, err := a.Info(c, url)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
