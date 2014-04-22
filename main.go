package main

import (
    //for RESTful API
    "github.com/go-martini/martini"
    //for sqlite3 storage, from drivers of go-wiki: https://code.google.com/p/go-wiki/wiki/SQLDrivers
    //"github.com/mattn/go-sqlite3"
    import "code,google.com/p/go-sqlite/go1/sqlite3"

)

func main() {
    m := martini.Classic()

    m.Get("/", func() string {
        return "Hello world"
    })

    m.Get("/hello/:name", func() string {
        return "Hello " + params["name"]
    })


    //Create
    //POST http://localhost:3000/task/add/data.json?
    m.Post("/task/add", func() {

    })

    //Read one
    //GET http://localhost:3000/task/:id
    m.Get("/task/:id", func() {

    })
    
    //Read all
    //GET http://localhost:3000/task/list
    m.Get("/task/list", func() {

    })

    //Update one
    //POST http://localhost:3000/task/:id/data.json?
    m.Post("/task/:id", func() {

    })

    //Delete one
    //Delete http://localhost:3000/task/delete/:id
    m.Delete("/task/delete/:id", func() {

    })

    m.Run()
}

