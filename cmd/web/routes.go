package web

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/lifeofdan/quickstart/cmd/web/pages"
	"github.com/vearutop/statigz"
	"github.com/vearutop/statigz/brotli"
)

func Routes() {
	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/assets/", statigz.FileServer(Assets, brotli.AddEncoding))
	http.HandleFunc("/clicked", ClickHomeBtnHandler)
	http.HandleFunc("/robots.txt", RobotsHandler)
}

func ClickHomeBtnHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "<h1 id=\"parent-div\" class=\"is-size-1\">Hello from backend</h1>")
}

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "User-agent: *")
}
