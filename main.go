package main

import (
	//"fmt"
	//"bytes"
	"log"
	"encoding/json"

    //for RESTful API
    "github.com/go-martini/martini"
	"github.com/codegangsta/martini-contrib/binding"
	//"github.com/astaxie/beego"

    //for sqlite3 storage, from drivers of go-wiki: https://code.google.com/p/go-wiki/wiki/SQLDrivers
	"database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type Task struct {
	Description	string;
	Due			string;
	Completed	bool;	
}

func main() {
	m := martini.Classic()

	m.Post("/task/add/json.data", func() []byte {
		return []byte(`test byte data`)	
	})

	m.Get("/task/:id", func(params martini.Params) []byte {
		//get database
		db, err := sql.Open("sqlite3", "./db_tasks.db")
		if err != nil {
			//Fatal is equivaleng to Printf() followed by a call to os.Exit(1)
			log.Fatal(err)
		}
		defer db.Close()
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
		//now should be equals to-> m := Task{"text", "2008-01-01 10:00:00", 0}
		ret_json, err := json.Marshal(t)
		if err != nil {
			log.Fatal(err)
		}
		//return []byte(`test byte data`)
		return ret_json
	})

	m.Get("/task/list", func() []byte {
		/*
		//get database
		db, err := sql.Open("sqlite3", "./db_tasks.db")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		//define SQL
		rows, err := db.Query("Select *
							   From t_tasks")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		//create a Map

		for rows.Next() {
			var description string;
			var due			string;
			var completed	boolean;
			rows.Scan(&description, &due, &completed)
			//append things after the map
		}
		*/
		return []byte(`test byte data`)
	})

	m.Post("/task/add", binding.Bind(Task{}), func(task Task) []byte {

		return []byte(`test byte data`)

	})

	m.Put("/task/:id/data.json", binding.Bind(Task{}), func(task Task) []byte {
		return []byte(`test byte data`)

	})

	m.Delete("/task/delete/:id", func(task Task) []byte {
		return []byte(`test byte data`)

	})

	m.Run()
}
