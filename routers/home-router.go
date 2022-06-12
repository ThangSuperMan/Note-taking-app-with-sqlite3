package routers

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Homepage")
	t, _ := template.ParseFiles("./views/index.html")

	t.Execute(w, nil)
}
