package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/lifeofdan/quickstart/cmd/web"
)

func main() {
	http.Handle("/assets/", http.FileServer(http.FS(web.Assets)))
	http.Handle("/", templ.Handler(web.Home()))

	// Check htmx is working
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		fmt.Fprintf(w, "<h1 id=\"parent-div\">Hello from backend</h1>")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Sprintf("cannot start http server: %s", err))
	}
}
