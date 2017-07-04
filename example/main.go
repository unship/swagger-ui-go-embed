package main

import (
	"net/http"

	"fmt"

	"github.com/biolee/swagger-ui-go-embed/ui"
)

func main(){
	fmt.Println(ui.Asset("index.html"))

	http.Handle("/",ui.Handler)
	http.ListenAndServe(":8081",nil)
}
