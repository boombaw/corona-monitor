package rapidapi

import (
	"encoding/json"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/parnurzeal/gorequest"
)

// API struct
type API struct{}

// Request api
func (a *API) Request(ctx echo.Context, path string) (RapidApiResponse, error) {
	var result RapidApiResponse
	_, body, _ := gorequest.New().
		Get(path).
		Set(os.Getenv("auth_header_key1"), os.Getenv("auth_header_value1")).
		Set(os.Getenv("auth_header_key2"), os.Getenv("auth_header_value2")).
		End()

	var val map[string]interface{}
	err := json.Unmarshal([]byte(body), &val)
	if err != nil {
		return RapidApiResponse{}, err
	}

	result.Data = map[string]interface{}{
		"information": val,
	}

	return result, nil
}
