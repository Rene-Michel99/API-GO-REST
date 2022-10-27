package main

import (
    "github.com/Rene-Michel99/API-GO-REST/database"
    "github.com/gofiber/fiber/v2"
)

func main(){
	database.ConnectDB()

    app := fiber.New()

    setupRoutes(app)

    app.Listen(":3000")
}