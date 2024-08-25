package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Game struct {
	Pressed  map[string]bool
	Attempts int
	Guessed  [4]string
}

func newGame() *Game {
	g := Game{}
	g.Attempts = 4
	g.Guessed = [4]string{"_", "_", "_", "_"}
	g.Pressed = make(map[string]bool)
	g.Pressed["Q"] = false
	g.Pressed["W"] = false
	g.Pressed["E"] = false
	g.Pressed["R"] = false
	g.Pressed["T"] = false
	g.Pressed["Y"] = false
	g.Pressed["U"] = false
	g.Pressed["I"] = false
	g.Pressed["O"] = false
	g.Pressed["P"] = false
	g.Pressed["A"] = false
	g.Pressed["S"] = false
	g.Pressed["D"] = false
	g.Pressed["F"] = false
	g.Pressed["G"] = false
	g.Pressed["H"] = false
	g.Pressed["J"] = false
	g.Pressed["K"] = false
	g.Pressed["L"] = false
	g.Pressed["Z"] = false
	g.Pressed["X"] = false
	g.Pressed["C"] = false
	g.Pressed["V"] = false
	g.Pressed["B"] = false
	g.Pressed["N"] = false
	g.Pressed["M"] = false
	return &g
}

func main() {
	e := echo.New()

	game := newGame()
	answer := [4]string{"S", "H", "I", "P"}

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		game := newGame()
		fmt.Println(game.Guessed)
		return c.Render(200, "index", game)
	})

	e.POST("/try/:letter", func(c echo.Context) error {
		letter := c.Param("letter")
		game.Pressed[letter] = true
		if letter == "S" {
			game.Guessed[0] = "S"
		} else if letter == "H" {
			game.Guessed[1] = "H"
		} else if letter == "I" {
			game.Guessed[2] = "I"
		} else if letter == "P" {
			game.Guessed[3] = "P"
		} else {
			game.Attempts--
		}
		fmt.Println(game.Guessed)
		if game.Attempts == 0 {
			return c.Render(200, "lost", game)
		} else if game.Guessed == answer {
			game := newGame()
			return c.Render(200, "win", game)
		} else {
			return c.Render(200, "index", game)
		}
	})

	e.Logger.Fatal(e.Start(":42069"))
}
