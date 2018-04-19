/*
	服务器上的资源可以attachment下载
 */
package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.File("index1.html")
	})
	e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})

	e.GET("/attachment", func(c echo.Context) error {
		return c.Attachment("attachment.txt", "attachment.txt")
	})


	e.Logger.Fatal(e.Start(":1323"))
}