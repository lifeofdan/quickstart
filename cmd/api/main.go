package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		fmt.Fprintln(w, "Hello world!")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Sprintf("cannot start http server: %s", err))
	}
}
