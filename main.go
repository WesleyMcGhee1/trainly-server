package main

import (
	"trainly/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()

  e.GET("/test", controllers.Test)

  e.Logger.Fatal(e.Start(":3003"))
}
