package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/", "public/index1.html")
	e.Logger.Fatal(e.Start(":1323"))
}