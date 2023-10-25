package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	group := e.Group("/api")

	group.Use(LoggerMiddleware)

	group.GET("/users", func(c echo.Context) error {
		return c.String(200, "Get all users")
	})
	group.POST("/users", func(c echo.Context) error {
		return c.String(200, "Create a user")
	})

	// add middleware to only one route
	group.GET("/users/:id", func(c echo.Context) error {
		return c.String(200, "Get user with id")
	})

	// curl -X GET http://localhost:1323/api/users/1

	// Nested group
	group2 := group.Group("/idk")
	group2.Use(AuthMiddleware)

	group2.GET("/users", func(c echo.Context) error {
		return c.String(200, "Get all users")
	})

	e.Logger.Fatal(e.Start(":1323"))

}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("hiiiiiii")
		log.Printf("method=%s, path=%s", c.Request().Method, c.Path())
		return next(c)
	}
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("authenticating request")
		return next(c)
	}
}

// to access nested group
