package main

import (
	"fmt"
	"formus/config"
	"formus/handler"
	"net/http"
	"os"
)

func main() {

	cfg := config.LoadConfig()

	passwordHnldr := &handler.PasswordHandler{
		Password: cfg.Password,
	}
	http.HandleFunc("/", handler.LandingHandler)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/login", passwordHnldr.LoginHandler)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server running")
		os.Exit(1)
	}
}
