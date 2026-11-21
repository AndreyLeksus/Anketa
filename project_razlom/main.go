package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projectrazlom/models"
	"projectrazlom/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	ankets := []models.Anketa{}


	


	ankets, err := storage.LoadAnkets()
	if err != nil {
		fmt.Println(err)
		
	}

	e := echo.New()

	e.Static("/", "static")

	e.GET("/anketa", func(c echo.Context) error {
		return c.File("template/anketa-short.html")
	})

	tmpl := template.Must(template.ParseFiles("template/result.html"))

	e.POST("/res", func(c echo.Context) error {
		a := models.Anketa{
			ID:     len(ankets) + 1,
			Name:   c.FormValue("name"),
			Age:    c.FormValue("age"),
			City:   c.FormValue("city"),
			Gender: c.FormValue("gender"),
			Height: c.FormValue("height"),
			Weight: c.FormValue("weight"),
			About:  c.FormValue("about"),
		}

		ankets = append(ankets, a)
		fmt.Println(ankets)
		storage.SavedAnket(ankets)

		return tmpl.Execute(c.Response(), a)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Сервер работает!")
	})

	e.GET("/exit", func(c echo.Context) error {
		go e.Close()
		return c.String(200, "сервер закрыт")
	})

	err = e.Start(":8080")

	if err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}
