package main

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
	"fmt"
)

// create cookie
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// read cookie
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

//read all cookie
func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all cookie")
}

func main() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
	e.POST("writecookie", writeCookie)
	e.GET("read_cookie", readCookie)
	e.GET("read_cookies", readAllCookies)
}
