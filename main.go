package main

import (
	//"fmt"
	"log"
	//for RESTful API
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render" //for rendering returning JSON
	//for sqlite3 storage, from drivers of go-wiki: https://code.google.com/p/go-wiki/wiki/SQLDrivers
	"database/sql"
	//use the initialization inside without actual use it: http://golang.org/doc/effective_go.html#blank
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Description string `json:"description" binding:"required"`
	Due         string `json:"due"`
	Completed   bool   `json:"completed"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	//get database
	db, err := sql.Open("sqlite3", "./db_tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//using martini-contrib/binding for receiving incoming JSON
	m.Post("/task/add/data.json", binding.Json(Task{}), func(task Task, r render.Render) {
		stmt, err := db.Prepare(`Insert Into t_tasks
					 (description, due, completed)
					 Values (?, ?, ?)`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(task.Description, task.Due, task.Completed)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Get("/task/list", func(r render.Render) {
		rows, err := db.Query(`Select description, due, completed
				       From t_tasks`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer rows.Close()

		//create a Struct array
		res := []Task{}
		for rows.Next() {
			//var tempID string
			t := Task{}
			rows.Scan(&t.Description, &t.Due, &t.Completed)
			//append things after the array
			res = append(res, t)
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": res})
	})

	m.Get("/task/:id", func(params martini.Params, r render.Render) {
		//define SQL
		stmt, err := db.Prepare(`Select description, due, completed
					 From t_tasks
					 Where id = ?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()
		//execute with parameter
		t := Task{}
		err = stmt.QueryRow(params["id"]).Scan(&t.Description, &t.Due, &t.Completed)

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success", "results": t})
	})

	m.Put("/task/:id/data.json", binding.Json(Task{}), func(params martini.Params, task Task, r render.Render) {
		stmt, err := db.Prepare(`Update t_tasks
					 Set description=?, due=?, completed=?
					 Where id=?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(task.Description, task.Due, task.Completed, params["id"])

		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Delete("/task/delete/:id", func(params martini.Params, r render.Render) {
		stmt, err := db.Prepare(`Delete From t_tasks
					 Where id = ?`)
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(params["id"])
		if err != nil {
			r.JSON(500, map[string]interface{}{"status": "error", "message": err.Error()})
			return
		}

		r.JSON(200, map[string]interface{}{"status": "success"})
	})

	m.Run()
}
