package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/lifeofdan/quickstart/cmd/web"
	"github.com/lifeofdan/quickstart/cmd/web/pages"
	"github.com/vearutop/statigz"
	"github.com/vearutop/statigz/brotli"
)

func main() {
	http.Handle("/assets/", statigz.FileServer(web.Assets, brotli.AddEncoding))
	http.Handle("/", templ.Handler(pages.Home()))
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "User-agent: *")
	})

	// Check htmx is working
	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		fmt.Fprintf(w, "<h1 id=\"parent-div\" class=\"is-size-1\">Hello from backend</h1>")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Sprintf("cannot start http server: %s", err))
	}
}
