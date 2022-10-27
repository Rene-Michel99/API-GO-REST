package main

import (
	"github.com/gofiber/fiber/v2"
    "github.com/Rene-Michel99/API-GO-REST/database"
	"github.com/Rene-Michel99/API-GO-REST/models"
)


func Update(c *fiber.Ctx) error {
	book := new(models.Book)

	error := c.BodyParser(book)
	if (error != nil) {
        return c.Status(500).JSON(fiber.Map{"message": error.Error()})
    }

    database.Update(book)

    return c.Status(200).JSON(book)
}

func Delete(c *fiber.Ctx) error {
    book := new(models.Book)

    error := c.BodyParser(book)
    if (error != nil) {
        return c.Status(500).JSON(fiber.Map{"message": error.Error()})
    }

    database.Delete(book)

    return c.Status(200).JSON(book)
}

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}


func setupRoutes(app *fiber.App) {
	app.Get("/", Home)
    app.Get("/update", Update)
    app.Get("/delete", Delete)
}