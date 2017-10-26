// A guessing game web application
// https://stackoverflow.com/questions/30609485/with-golang-templates-how-can-i-set-a-variable-in-each-template

package main

import (
	// "fmt"
	"html/template"
	"net/http"
)
type WebData struct {
    Title string
	Head2 string
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

    // Title := "Guessing Game"
	// fmt.Println(Title)


	http.ServeFile(w, r, "guessing.html")
	// srv, _ := http.ServeFile(w, r, "guessing.html")

	// srv.Execute(w, Title)
	
}

func layoutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("guess.tmpl")
    wd := &WebData {
        Title: "Guessing Game",
		Head2: "Pick a number between 1 & 20:",
    }
    tmpl.Execute(w, wd)
}

// main
func main() {
	http.HandleFunc("/guess", guessHandler)
	http.HandleFunc("/layout", layoutHandler)
	http.ListenAndServe(":8080", nil)
}