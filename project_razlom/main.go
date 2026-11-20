package main

import (
	"html"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	e.GET("/anketa", func(c echo.Context) error {
		return c.File("template/anketa-short.html")
	})

	e.POST("/res", func(c echo.Context) error {
		name := c.FormValue("name")
		age := c.FormValue("age")
		city := c.FormValue("city")
		gender := c.FormValue("gender")
		height := c.FormValue("height")
		weight := c.FormValue("weight")
		about := c.FormValue("about")

		if gender == "f" {
			gender = "Женский"
		} else {
			gender = "Мужской"
		}

		file, err := os.ReadFile("template/result.html")
		if err != nil {
			return err
		}

		page := string(file)

		page = strings.ReplaceAll(page, "@name@", html.EscapeString(name))
		page = strings.ReplaceAll(page, "@age@", html.EscapeString(age))
		page = strings.ReplaceAll(page, "@city@", html.EscapeString(city))
		page = strings.ReplaceAll(page, "@gender@", gender)
		page = strings.ReplaceAll(page, "@height@", html.EscapeString(height))
		page = strings.ReplaceAll(page, "@weight@", html.EscapeString(weight))
		page = strings.ReplaceAll(page, "@about@", about)

		return c.HTML(http.StatusOK, page)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
