package main

import (
	"github.com/Rene-Michel99/API-GO-REST/database"
	"github.com/Rene-Michel99/API-GO-REST/models"
	"github.com/gofiber/fiber/v2"
)

func Insert(c *fiber.Ctx) error {
    book := new(models.Book)

    error := c.BodyParser(book)
    if error != nil {
        return c.Status(500).JSON(fiber.Map{"message": error.Error()})
    }

    result := database.Insert(book)

    if result.RowsAffected > 0 {
        return c.Status(200).JSON(book)
    }
    return c.Status(500).JSON(result.Error)
}

func Get(c *fiber.Ctx) error {
	book := new(models.Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	result := database.Get(book)

	if result.RowsAffected > 0 {
        return c.Status(200).JSON(book)
	}
	return c.Status(500).JSON(result.Error)
}

func Update(c *fiber.Ctx) error {
	book := new(models.Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	result := database.Update(book)

	if result.RowsAffected > 0 {
		return c.Status(200).JSON(result.RowsAffected)
	}
	return c.Status(500).JSON(result.Error)
}

func Delete(c *fiber.Ctx) error {
	book := new(models.Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	result := database.Delete(book)

	if result.RowsAffected > 0 {
		return c.Status(200).JSON(result.RowsAffected)
	}
	return c.Status(500).JSON(result.Error)
}

func setupRoutes(app *fiber.App) {
    app.Post("/create", Insert)
	app.Get("/get", Get)
	app.Put("/update", Update)
	app.Delete("/delete", Delete)
}
