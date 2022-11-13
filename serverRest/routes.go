package main

import (
	"net/rpc"
	"time"
    "gorm.io/gorm"
	"github.com/gofiber/fiber/v2"
)


// Objeto Book utilizado pelo banco
type Book struct{
    gorm.Model
    ID int                  `json:"id" gorm:primaryKey;autoIncrement`
    Name string             `json:"name" gorm:text;not null;default:null`
    Author string           `json:"author" gorm:text;not null;default:null`
    Synopsis string         `json:"synopsis" gorm:text;not null;default:null`
    LaunchDate time.Time    `json:"launchDate" gorm:date;not null;default:null`
    CopyQnt int32           `json:"copyQnt" gorm:int;not null;default:null`
}


// Objeto de resposta para o REST
type Response struct{
    Status int64
    Book Book
}

// CRUD
func Insert(c *fiber.Ctx) error {
    var reply Response
	book := new(Book)

	error := c.BodyParser(book)
	if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    // Conecta ao servidor RPC
    client, error := rpc.DialHTTP("tcp", "database-manager:4040")
    if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    // Chamada da função AddBook do servidor RPC (RPCServer.go)
    res := client.Call("API.AddBook", book, &reply)
    if (res == nil) {
		return c.Status(200).JSON(reply)
	}

	return c.Status(500).JSON(book)
}

func Get(c *fiber.Ctx) error {
    var reply Response
	book := new(Book)

	error := c.BodyParser(book)
    if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    client, error := rpc.DialHTTP("tcp", "database-manager:4040")
	if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    res := client.Call("API.GetBook", book, &reply)
    if (res == nil) {
        return c.Status(200).JSON(reply)
	}

	return c.Status(500).JSON(book)
}

func Update(c *fiber.Ctx) error {
    var reply Response
	book := new(Book)

	error := c.BodyParser(book)
	if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    client, error := rpc.DialHTTP("tcp", "database-manager:4040")
    if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    res := client.Call("API.UpdateBook", book, &reply)
	if (res == nil) {
		return c.Status(200).JSON(reply)
	}

	return c.Status(500).JSON(book)
}

func Delete(c *fiber.Ctx) error {
    var reply Response
	book := new(Book)

	error := c.BodyParser(book)
    if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    client, error := rpc.DialHTTP("tcp", "database-manager:4040")
	if (error != nil) {
		return c.Status(500).JSON(fiber.Map{"message": error.Error()})
	}

    res := client.Call("API.DeleteBook", book, &reply)
    if (res == nil) {
		return c.Status(200).JSON(reply)
	}

	return c.Status(500).JSON(book)
}

func setupRoutes(app *fiber.App) {
	app.Post("/create", Insert)
	app.Get("/get", Get)
	app.Put("/update", Update)
	app.Delete("/delete", Delete)
}
