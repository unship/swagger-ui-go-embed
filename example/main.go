package main

import (
	"net/http"

	"github.com/biolee/swagger-ui-go-embed/ui"
)

func main() {
	http.Handle("/", ui.Handler)
	http.ListenAndServe(":8081", nil)
}
