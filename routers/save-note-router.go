package routers

import (
	"NoteTakingApp/controllers"
	"fmt"
	"net/http"
)

func SaveNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Save note method")
	controllers.GetInfoNote(w, r)
	// fmt.Fprintln(w, "<p>Successfully submit your note</p>")

}
