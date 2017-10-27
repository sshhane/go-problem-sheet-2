// A guessing game web application
// https://stackoverflow.com/questions/30609485/with-golang-templates-how-can-i-set-a-variable-in-each-template
// https://github.com/data-representation/go-cookies/blob/master/go-cookie.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"math/rand"
	"strconv"
	"time"
)
type WebData struct {
    Title string
	Head2 string
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	// http.ServeFile(w, r, "guessing.html")


	// cookies
	target:=0
	var cookie, err = r.Cookie("target")
	if err == nil {
		target, _ = strconv.Atoi(cookie.Value)
	}

	target = rand.Intn(20)
	// test
	// fmt.Printf("target: %v", target)

	// create cookie
	cookie = &http.Cookie {
		Name: "target",
		Value: strconv.Itoa(target),
		Expires: time.Now().Add(72 * time.Hour),
	}

	// set cookie
	http.SetCookie(w, cookie)
	//////////////////////////

    wd := &WebData {
        Title: "Guessing Game",
		Head2: "Pick a number between 1 & 20:",
    }
	
	tmpl, _ := template.ParseFiles("guess.tmpl")

    tmpl.Execute(w, wd)

}

func layoutHandler(w http.ResponseWriter, r *http.Request) {
}

// main
func main() {
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}