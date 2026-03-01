package main

import (
	"fmt"
	"formus/handler"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/form", handler.FormHandler)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server running")
		os.Exit(1)
	}
}
