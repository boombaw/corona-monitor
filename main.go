package main

import (
	"log"

	"github.com/boombaw/corona-monitor/action"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	e := echo.New()
	// middleware
	e.Use(middleware.Logger())
	// Routes
	// Get Information All Country
	e.GET("/corona", action.AllCountry)
	// Get Information Corona Virus By Country
	e.GET("/corona/:country", action.InfoByCountry)
	// Get Information Affected Country Corona Virus
	e.GET("/corona/affected", action.AffectedCountry)
	// Get Information World Total Status
	e.GET("/corona/world", action.WorldStat)
	// Get Information History Corona Virus By Country
	e.GET("/corona/history/:country", action.HistoryByCountry)

	e.Logger.Fatal(e.Start(":8080"))
}
