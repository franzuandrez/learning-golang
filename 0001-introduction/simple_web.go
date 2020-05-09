package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	_, _ = fmt.Fprintf(w, "Whoa , Go is neat!")
}

func main() {
	http.HandleFunc("/", index_handler)
	_ = http.ListenAndServe(":9001", nil)
}
