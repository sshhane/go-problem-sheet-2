// A guessing game web application

package main

import (
	"fmt"
	"net/http"
	"bytes"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "r.URL:              ",  r.URL              )
	

	fmt.Fprintln(w, "Header:")
	for key, value := range r.Header {
		fmt.Fprintln(w, "\t" + key + ":", value)
	}

	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Fprintln(w, "r.Body:             ",  body.String())
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}