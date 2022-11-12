package main

import (
    "log"
    "net"
    "net/rpc"
    "net/http"
)

type API int

type Book struct{
    Name string
    Author string
}

var listTest []Book

func (a *API) GetDB(empty string, reply *[]Book) error {
    *reply = listTest
    return nil
}

func (a *API) GetByName(name string, reply *Book) error {
    var getItem Book

    for _, val := range listTest {
        if val.Name == name {
            getItem = val
        }
    }

    *reply = getItem

    return nil
}

func (a *API) AddItem(item Book, reply *Book) error {
    log.Printf("Adding item ", item)
    listTest = append(listTest, item)
    *reply = item
    return nil
}

func (a *API) EditItem(item Book, reply *Book) error {
    var changed Book

    for idx, val := range listTest {
        if val.Name == item.Name {
            listTest[idx] = Book{item.Name, item.Author}
            changed = listTest[idx]
        }
    }

    *reply = changed
    return nil
}

func (a *API) DeleteItem(item Book, reply *Book) error {
    var del Book

    for idx, val := range listTest {
        if val.Name == item.Name && val.Author == item.Author {
            listTest = append(listTest[:idx], listTest[idx+1:]...)
            del = item
            break
        }
    }

    *reply = del
    return nil
}

func main() {
    ConnectDB()
    rpcServerPort := ":4040"
    var api = new(API)
    error := rpc.Register(api)

    if (error != nil){
        log.Fatal("Error registering API", error)
    }

    rpc.HandleHTTP()
    listener, error := net.Listen("tcp", rpcServerPort)

    if (error != nil){
        log.Fatal("Error listening on PORT: ", rpcServerPort, error)
    }

    log.Printf("RPC Server on port %s", rpcServerPort)
    error = http.Serve(listener, nil)
    if(error != nil) {
        log.Fatal("Error initializinng server ", error)
    }
}