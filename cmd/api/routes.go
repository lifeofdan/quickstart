package api

import (
	"fmt"
	"net/http"
)

func Routes() {
	// All api routes are prefixed with "/api/v1" by convention
	http.HandleFunc("/api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from quickstart api!")
	})
}
