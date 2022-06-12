package controllers

import (
	"NoteTakingApp/models"
	"fmt"
	"net/http"
	"time"
)

type Note struct {
	title     string
	content   string
	updateAt  time.Time
	id_author string
}

func GetInfoNote(w http.ResponseWriter, r *http.Request) {
	var noteTitle string = r.FormValue("note-title")
	var noteCotent string = r.FormValue("note-content")

	fmt.Print("GetInfoNote controler")

	var note Note = Note{
		title:    noteTitle,
		content:  noteCotent,
		updateAt: time.Now(),
	}

	if noteTitle != "" && noteCotent != "" {
		fmt.Println("fine")
	}

	models.SaveNote(note)

	fmt.Println(note)
}
