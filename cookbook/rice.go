// 处理静态文件
package main

import (
	"net/http"
	"github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// the file server for rice. "xxx" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())
	// serves the index.html from rice
	e.GET("/", echo.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))

	e.Logger.Fatal(e.Start(":1323"))
}