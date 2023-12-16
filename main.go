package main

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var tmplt *template.Template
var bossTimes []Boss

type something struct {
	Stuff1, Stuff2 string
}

type BossTimer struct {
	boss []Boss
}

func hello(c echo.Context) error {
	tmplt, _ := template.ParseFiles("page.html")
	sthing := something{
		Stuff1: "a",
		Stuff2: "b",
	}
	err := tmplt.Execute(c.Response().Writer, sthing)

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
