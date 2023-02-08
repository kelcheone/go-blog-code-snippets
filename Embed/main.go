package main

import (
    _ "embed"
//    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
)

//go:embed static/index.html
var indexHTML string

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.HTML(http.StatusOK, indexHTML)
    })

    e.Logger.Fatal(e.Start(":8080"))
}
