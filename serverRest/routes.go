package main

import (
	"github.com/gofiber/fiber/v2"
)

type Book struct{
    Name string
    Author string
}

func Insert(c *fiber.Ctx) error {
    book := new(Book)

    error := c.BodyParser(book)
    if error != nil {
        return c.Status(500).JSON(fiber.Map{"message": error.Error()})
    }


    return c.Status(500).JSON(book)
}

func Get(c *fiber.Ctx) error {
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}


    return c.Status(500).JSON(book)
}

func Update(c *fiber.Ctx) error {
	book := new(Book)

	error := c.BodyParser(book)
    if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}


    return c.Status(500).JSON(book)
}

func Delete(c *fiber.Ctx) error {
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	return c.Status(500).JSON(book)
}

func setupRoutes(app *fiber.App) {
    app.Post("/create", Insert)
	app.Get("/get", Get)
	app.Put("/update", Update)
	app.Delete("/delete", Delete)
}
