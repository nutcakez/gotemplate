package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var tmplt *template.Template
var bossTimes []Boss
var apikey = "D850DD23-4F76-C444-B5B0-EA5909CF614CF8538C0A-84BF-41B9-9307-AD7A5E60D9C8"

type something struct {
	Stuff1, Stuff2 string
}

type BossTimer struct {
	Boss []Boss
}

func hello(c echo.Context) error {
	tmplt, _ := template.ParseFiles("page.html")
	boss := FilterForTime(GetBossTimes())
	bossInput := BossTimer{
		Boss: boss,
	}
	err := tmplt.Execute(c.Response().Writer, bossInput)

	return err
}

func bossTimers(c echo.Context) error {

	tmplt, _ := template.ParseFiles("page.html")
	sthing := something{
		Stuff1: "a",
		Stuff2: "b",
	}
	err := tmplt.Execute(c.Response().Writer, sthing)

	return err
}

func test(c echo.Context) error {
	tmplt, _ := template.ParseFiles("response.html")

	err := tmplt.Execute(c.Response().Writer, nil)

	return err
}

func main() {
	e := echo.New()

	GetAvailableWorldBosses()

	bossTimes = GetBossTimes()
	filtered := FilterForTime(bossTimes)
	filtered = append(filtered, Boss{})
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/test", test)
	e.GET("/times", bossTimers)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
