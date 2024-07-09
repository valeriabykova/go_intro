package main

import (
	"awesomeProject/accounts"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account", accountsHandler.GetAccount)
	e.POST("/account/create", accountsHandler.CreateAccount)
	e.POST("/account/rename", accountsHandler.RenameAccount)
	e.POST("/account/update", accountsHandler.SetBalance)
	e.POST("/account/delete", accountsHandler.DeleteAccount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
