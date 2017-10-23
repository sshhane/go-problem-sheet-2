// A guessing game web application

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "r.URL:              ",  r.URL              )
	fmt.Fprintln(w, "Guessing Game")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}