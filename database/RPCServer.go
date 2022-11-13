package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
    "errors"
    "github.com/Rene-Michel99/API-GO-REST/database/models"
)

// Objeto de resposta para a aplicação REST
type Response struct{
    Status int64
    Book models.Book
}

type API int

// CRUD
func (a *API) GetBook(book models.Book, reply *Response) error {
    if (book == models.Book{}) {
        return errors.New("None object has sent")
    }
    result := Get(&book)

    reply.Status = result.RowsAffected
    reply.Book = book

    return nil
}

func (a *API) AddBook(book models.Book, reply *Response) error {
    if (book == models.Book{}) {
        return errors.New("None object has sent")
    }
    result := Insert(&book)

    reply.Status = result.RowsAffected
    reply.Book = book

    return nil
}

func (a *API) UpdateBook(book models.Book, reply *Response) error {
    if (book == models.Book{}) {
        return errors.New("None object has sent")
    }
    result := Update(&book)

    reply.Status = result.RowsAffected
    reply.Book = book

    return nil
}

func (a *API) DeleteBook(book models.Book, reply *Response) error {
    if (book == models.Book{}) {
        return errors.New("None object has sent")
    }
    result := Delete(&book)

    reply.Status = result.RowsAffected
    reply.Book = book

    return nil
}

// Servidor RPC
func main() {
	ConnectDB()   // Conexão com o banco
	rpcServerPort := ":4040"

	var api = new(API)
	error := rpc.Register(api) // Endereça a API RPC

    if (error != nil) {
		log.Fatal("Error registering API", error)
	}

	rpc.HandleHTTP()  // utiliza o HTTP
	listener, error := net.Listen("tcp", rpcServerPort)

	if (error != nil) {
		log.Fatal("Error listening on PORT: ", rpcServerPort, error)
	}

	log.Printf("RPC Server on port %s", rpcServerPort)
	error = http.Serve(listener, nil)  // Disponibiliza acesso ao server RPC via HTTP
    if (error != nil) {
		log.Fatal("Error initializinng server ", error)
	}
}
