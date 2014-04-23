package main

import (
	"encoding/json"
	"fmt"
	"log"
	//for RESTful API
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	//for sqlite3 storage, from drivers of go-wiki: https://code.google.com/p/go-wiki/wiki/SQLDrivers
	"database/sql"
	//use the initialization inside without actual use it: http://golang.org/doc/effective_go.html#blank
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Description string `form: "Description" json:"Descrioption" binding:"required"`
	Due         string `form: "Due" json:"Due"`
	Completed   bool   `form: "Completed" json:"Completed"`
}

type Error struct {
	message string
}

func main() {
	m := martini.Classic()

	//get database
	db, err := sql.Open("sqlite3", "./db_tasks.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//using martini-contrib for receiving incoming JSON
	m.Post("/task/add/data.json", binding.Json(Task{}), func(task Task) {
		stmt, err := db.Prepare(`Insert Into t_tasks
					 (description, due, completed)
					 values (?, ?, ?)`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(task.Description, task.Due, task.Completed)
		if err != nil {
			log.Fatal(err)
		}
	})

	m.Get("/task/list", func() []byte {
		//define SQL
		rows, err := db.Query(`Select description, due, completed
				       From t_tasks`)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
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
		ret_json, err := json.Marshal(res)
		return ret_json
	})

	m.Get("/task/:id", func(params martini.Params) []byte {
		//define SQL
		stmt, err := db.Prepare(`Select description, due, completed
					 From t_tasks
					 Where id = ?`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		//execute with parameter
		t := Task{}
		err = stmt.QueryRow(params["id"]).Scan(&t.Description, &t.Due, &t.Completed)
		if err != nil {
			log.Fatal(err)
		}
		//now should be equals to-> m := Task{"text", "2008-01-01 10:00:00", 0} for record 1
		ret_json, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		return ret_json
	})

	m.Put("/task/:id/data.json", binding.Json(Task{}), func(params martini.Params, task Task) {
		stmt, err := db.Prepare(`Update t_tasks
					 Set description=?, due=?, completed=?
					 Where id=?`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(task.Description, task.Due, task.Completed, params["id"])
		if err != nil {
			log.Fatal(err)
		}
	})

	m.Delete("/task/delete/:id", func(params martini.Params) {
		stmt, err := db.Prepare(`Delete From t_tasks
					 Where id = ?`)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(params["id"])
		if err != nil {
			log.Fatal(err)
		}
	})

	m.Run()
}
