package main

import (
	"fmt"
	"formus/config"
	"formus/database"
	"formus/handler"
	"formus/middleware"
	"net/http"
	"os"
)

func main() {
	database.InitDB()

	cfg := config.LoadConfig()

	passwordHnldr := &handler.PasswordHandler{
		Password: cfg.Password,
	}
	http.HandleFunc("/", handler.LandingHandler)
	http.HandleFunc("/form", handler.FormHandler)
	http.HandleFunc("/login", passwordHnldr.LoginHandler)
	http.HandleFunc("/admin", middleware.Auth(handler.AdminPanelHandler))
	http.HandleFunc("/form/create", handler.CreateFormHandler)

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println("error server running")
		os.Exit(1)
	}
}
