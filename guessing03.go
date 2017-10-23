// A guessing game web application

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Guessing Game<h1>")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}