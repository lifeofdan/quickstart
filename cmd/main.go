package main

import (
	"fmt"
	"net/http"

	"github.com/lifeofdan/quickstart/cmd/api"
	"github.com/lifeofdan/quickstart/cmd/web"
)

func main() {
	api.Routes()
	web.Routes()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(fmt.Sprintf("cannot start http server: %s", err))
	}
}
