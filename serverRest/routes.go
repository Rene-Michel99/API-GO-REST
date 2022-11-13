package main

import (
	"net/rpc"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID         int
	Name       string
	Author     string
	Synopsis   string
	LaunchDate time.Time
	CopyQnt    int32
}

func Insert(c *fiber.Ctx) error {
	var reply Book
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	client, error := rpc.DialHTTP("tcp", "db:4040")
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	res := client.Call("API.AddItem", book, &reply)
	if res == nil {
		return c.Status(200).JSON(book)
	}

	return c.Status(500).JSON(book)
}

func Get(c *fiber.Ctx) error {
	var reply Book
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	client, error := rpc.DialHTTP("tcp", "db:4040")
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	res := client.Call("API.GetByName", book.Name, &reply)
	if res == nil {
		return c.Status(200).JSON(book)
	}

	return c.Status(500).JSON(book)
}

func Update(c *fiber.Ctx) error {
	var reply Book
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	client, error := rpc.DialHTTP("tcp", "db:4040")
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	res := client.Call("API.EditItem", book, &reply)
	if res == nil {
		return c.Status(200).JSON(book)
	}

	return c.Status(500).JSON(book)
}

func Delete(c *fiber.Ctx) error {
	var reply Book
	book := new(Book)

	error := c.BodyParser(book)
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	client, error := rpc.DialHTTP("tcp", "db:4040")
	if error != nil {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

	res := client.Call("API.DeleteItem", book, &reply)
	if res == nil {
		return c.Status(200).JSON(book)
	}

	return c.Status(500).JSON(book)
}

func setupRoutes(app *fiber.App) {
	app.Post("/create", Insert)
	app.Get("/get", Get)
	app.Put("/update", Update)
	app.Delete("/delete", Delete)
}
