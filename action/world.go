package action

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/boombaw/corona-monitor/app"
	"github.com/boombaw/corona-monitor/external/rapidapi"
)

type World struct {
	apiReq rapidapi.API
}

func (w *World) Info(c echo.Context, url string) (rapidapi.RapidApiResponse, error) {
	result, err := w.apiReq.Request(c, url)
	if err != nil {
		return rapidapi.RapidApiResponse{}, err
	}
	return result, nil
}

//WorldStat for monitoring cases total status
func WorldStat(c echo.Context) (err error) {
	url := app.SetURL("worldstat.php")
	
	i := InfoCountry{}
	result, err := i.Info(c, url)
	
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, result)
}
