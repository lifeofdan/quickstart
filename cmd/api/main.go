package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/lifeofdan/quickstart/cmd/web"
)

func main() {
	http.Handle("/", templ.Handler(web.Home()))

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Sprintf("cannot start http server: %s", err))
	}
}
