package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// CREATE TABLE "Note" (
// 	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
// 	"title"	TEXT,
// 	"content"	TEXT,
// 	"updateAt"	TEXT,
// 	"id_author"	TEXT
// );

func SaveNote(note any) {

	fmt.Println("Save note method")
	fmt.Println(note)

	database, _ := sql.Open("sqlite3", "../mydatabase.db")

	statement, _ := database.Prepare(`
CREATE TABLE "Note" (
	"id"	INTEGER PRIMARY KEY AUTOINCREMENT,
	"title"	TEXT,
	"content"	TEXT,
	"updateAt"	TEXT,
	"id_author"	TEXT
);
	`)

	statement.Exec()

}
