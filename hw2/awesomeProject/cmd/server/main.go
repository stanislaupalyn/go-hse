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

	e.DELETE("/account/delete", accountsHandler.DeleteAccount)

	e.PATCH("/account/change-name", accountsHandler.ChangeAccountName)

	e.PATCH("/account/change-amount", accountsHandler.ChangeAccountAmount)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

/*
curl --header "Content-Type: application/json" --request POST --data '{"name":"alice","amount": 20}' 127.0.0.1:1323/account/create
curl --header "Content-Type: application/json" --request GET --data '{"name":"alice"}' 127.0.0.1:1323/account

curl --header "Content-Type: application/json" --request PATCH --data '{"name":"alice","new-name":"bob"}' 127.0.0.1:1323/account/change-name

curl --header "Content-Type: application/json" --request GET --data '{"name":"alice"}' 127.0.0.1:1323/account
curl --header "Content-Type: application/json" --request GET --data '{"name":"bob"}' 127.0.0.1:1323/account

curl --header "Content-Type: application/json" --request PATCH --data '{"name":"bob","new-amount":228}' 127.0.0.1:1323/account/change-amount
curl --header "Content-Type: application/json" --request GET --data '{"name":"bob"}' 127.0.0.1:1323/account

curl --header "Content-Type: application/json" --request DELETE --data '{"name":"bob"}' 127.0.0.1:1323/account/delete
curl --header "Content-Type: application/json" --request GET --data '{"name":"bob"}' 127.0.0.1:1323/account
*/
