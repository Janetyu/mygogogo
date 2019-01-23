package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(
		writer http.ResponseWriter,
		request *http.Request) {
		fmt.Fprintf(writer,
			"<h1>Hello World %s!</h1>",
			request.FormValue("name"))
	})

	http.ListenAndServe(":8888",nil)
}
