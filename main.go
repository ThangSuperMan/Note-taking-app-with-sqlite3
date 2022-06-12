package main

import (
	"NoteTakingApp/routers"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database, _ := sql.Open("sqlite3", "./mydatabase.db")

	statement, _ := database.Prepare(`
		CREATE TABLE IF NOT EXISTS "note" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"title"	INTEGER NOT NULL,
	"content"	INTEGER NOT NULL,
	"updateAt"	INTEGER
);
	`)

	statement.Exec()

	// statement.Exec()

	insertStatement, _ := database.Prepare("INSERT INTO note(title, content) VALUES (?, ?)")

	insertStatement.Exec("note 2", "content superman haha")

	rows, err := database.Query("SELECT id, title, content FROM note")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		var id string
		var title string
		var content string

		err = rows.Scan(&id, &title, &content)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, title, content)
	}

	fmt.Println("hello ohyeah")
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./assets"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	r.HandleFunc("/", routers.HomePage)
	r.HandleFunc("/save_note", routers.SaveNote).Methods("post")
	http.ListenAndServe(":3000", r)
}
