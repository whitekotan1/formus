package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server running")
		os.Exit(1)
	}
}
